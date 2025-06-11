package app;
// pxnMetrics Backend/Broker App

import(
	Log       "log"
	Flag      "flag"
	Time      "time"
	Service   "github.com/PxnPub/PxnGoCommon/service"
	Flagz     "github.com/PxnPub/PxnGoCommon/utils/flagz"
	HeartBeat "github.com/PxnPub/pxnMetrics/backend/heartbeat"
	Uplink    "github.com/PxnPub/pxnMetrics/backend/uplink"
);



type AppBackend struct {
	Version string
}



func New(version string) Service.App {
	return &AppBackend{
		Version: version,
	};
}

func (app *AppBackend) Main() {
	service := Service.New();
	service.Start();
	var flag_num_shards     int;
	var bind                string;
	var flag_batch_interval string;
	Flagz.Int   (&flag_num_shards,     "num-shards",     DefaultNumShards    );
	Flagz.String(&bind,                "bind",           DefaultBind         );
	Flagz.String(&flag_batch_interval, "batch-interval", DefaultBatchInterval);
	Flag.Parse();
	// num shards
	if flag_num_shards < 0   { flag_num_shards = 0; }
	if flag_num_shards > 255 { Log.Panic("Invalid number of shards: %d", flag_num_shards); }
	num_shards := uint8(flag_num_shards);
	// batch interval
	batch_interval, err := Time.ParseDuration(flag_batch_interval);
	if err != nil { Log.Panic(err); }
	if batch_interval <= 0 || batch_interval > Time.Hour {
		Log.Panic("Invalid batch-interval: %s", batch_interval);
	}
	// databases
//TODO
	// heartbeat
	heartbeat := HeartBeat.New(service, num_shards);
	// rpc server
	if num_shards == 0 {
		Log.Print("[Broker] Shard brokering is disabled");
	}
	uplink := Uplink.New(service, num_shards, bind);
	// start things
	if err := heartbeat.Start(); err != nil { Log.Panic(err); }
	if err := uplink.Start();    err != nil { Log.Panic(err); }
	service.WaitUntilEnd();
}
