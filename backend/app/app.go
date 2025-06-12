package app;
// pxnMetrics Backend/Broker App

import(
	Log       "log"
	Flag      "flag"
	Time      "time"
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
	UpLink    *UtilsRPC.Server
	Bind      string
}



func New(version string) Service.App {
	return &AppBackend{
		Version: version,
	};
}

func (app *AppBackend) Main() {
	app.Service := Service.New();
	app.Service.Start();
	var flag_num_shards     int;
	var flag_bind           string;
	var flag_batch_interval string;
	Flagz.Int   (&flag_num_shards,     "num-shards",     DefaultNumShards    );
	Flagz.String(&flag_bind,           "bind",           DefaultBind         );
	Flagz.String(&flag_batch_interval, "batch-interval", DefaultBatchInterval);
	Flag.Parse();
	// num shards
	if flag_num_shards < 0   { flag_num_shards = 0; }
	if flag_num_shards > 255 { Log.Panic("Invalid number of shards: %d", flag_num_shards); }
	num_shards := uint8(flag_num_shards);
	if num_shards == 0 {
		Log.Printf("%sShard brokering is disabled", LogPrefix);
	}
	// batch interval
	batch_interval, err := Time.ParseDuration(flag_batch_interval);
	if err != nil { Log.Panic(err); }
	if batch_interval <= 0 || batch_interval > Time.Hour {
		Log.Panic("Invalid batch-interval: %s", batch_interval);
	}
	app.Bind = flag_bind;
	// databases
//TODO
	// heartbeat
	app.HeartBeat = HeartBeat.New(app.Service, num_shards);
	// rpc server
	app.UpLink = UtilsRPC.NewServer(app.Service, app.Bind);
	FrontAPI.RegisterWebFrontAPIServer(
		app.UpLink.RPC,
		&API_Front{},
	);
	// start things
	if err := app.HeartBeat.Start(); err != nil { Log.Panic(err); }
	if err := app.UpLink.Start();    err != nil { Log.Panic(err); }
	app.Service.WaitUntilEnd();
}
