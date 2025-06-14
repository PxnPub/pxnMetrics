package app;
// pxnMetrics Backend/Broker App

import(
	Log       "log"
	Flag      "flag"
	Math      "math"
	Time      "time"
	GRPC      "google.golang.org/grpc"
	Service   "github.com/PxnPub/PxnGoCommon/service"
	Flagz     "github.com/PxnPub/PxnGoCommon/utils/flagz"
	HeartBeat "github.com/PxnPub/pxnMetrics/backend/heartbeat"
	UtilsRPC  "github.com/PxnPub/PxnGoCommon/rpc"
	FrontAPI  "github.com/PxnPub/pxnMetrics/api/front"
);



const LogPrefix = "[Broker] ";



type AppBackend struct {
	Version   string
	Service   *Service.Service
	HeartBeat *HeartBeat.HeartBeat
	UserMan   *HeartBeat.UserManager
	UpLink    *UtilsRPC.Server
	Bind      string
	NumShards uint8
	Checksum  uint16
	FrontAPI  *HeartBeat.API_Front
}



func New(version string) Service.App {
	return &AppBackend{
		Version: version,
	};
}

func (app *AppBackend) Main() {
	app.Service = Service.New();
	app.Service.Start();
	var flag_num_shards     int;
	var flag_checksum_init  int;
	var flag_batch_interval string;
	var flag_bind           string;
	Flagz.Int   (&flag_num_shards,     "num-shards",     DefaultNumShards    );
	Flagz.Int   (&flag_checksum_init,  "checksum",       DefaultChecksumInit );
	Flagz.String(&flag_batch_interval, "batch-interval", DefaultBatchInterval);
	Flagz.String(&flag_bind,           "bind",           DefaultBind         );
	Flag.Parse();
	// num shards
	if flag_num_shards < 0 { flag_num_shards = 0; }
	if flag_num_shards > Math.MaxUint8 {
		Log.Panic("Invalid number of shards: %d", flag_num_shards); }
	app.NumShards = uint8(flag_num_shards);
	if app.NumShards == 0 { Log.Printf("%sShard brokering is disabled", LogPrefix); }
	// checksum base init
	if flag_checksum_init < 0 { flag_checksum_init = 0; }
	if flag_checksum_init > Math.MaxUint16 {
		Log.Panic("Invalid checksum base value: %d", flag_checksum_init); }
	app.Checksum = uint16(flag_checksum_init);
	// batch interval
	batch_interval, err := Time.ParseDuration(flag_batch_interval);
	if err != nil { Log.Panic(err); }
	if batch_interval <= 0 || batch_interval > Time.Hour {
		Log.Panic("Invalid batch-interval: %s", batch_interval);
	}
	app.Bind = flag_bind;
	// user manager
	app.UserMan = HeartBeat.NewUserManager().
		AllowIP("127.0.0.1", "lop").
		AddPermWeb("lop").
		AddPermShard("lop", 1);
	// databases
//TODO
	// heartbeat
	app.HeartBeat = HeartBeat.New(app.Service, app.NumShards);
	// rpc server
	app.UpLink = UtilsRPC.NewServer(app.Service, app.Bind);
	app.UpLink.RPC = GRPC.NewServer(
		GRPC.ChainUnaryInterceptor(
			UtilsRPC.NewAuthByIP(app.UserMan.AllowIPs),
			app.UserMan.NewInterceptor(),
		),
	);
	app.FrontAPI = HeartBeat.NewFrontAPI(
		app.HeartBeat,
		app.UserMan,
		app.NumShards,
		app.Checksum,
	);
	FrontAPI.RegisterWebFrontAPIServer(app.UpLink.RPC, app.FrontAPI);
	// start things
	if err := app.HeartBeat.Start(); err != nil { Log.Panic(err); }
	if err := app.UpLink.Start();    err != nil { Log.Panic(err); }
	app.Service.WaitUntilEnd();
}
