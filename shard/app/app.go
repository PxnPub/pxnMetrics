package app;
// pxnMetrics Shard App

import(
	Log      "log"
	Flag     "flag"
	Service  "github.com/PxnPub/PxnGoCommon/service"
	Flagz    "github.com/PxnPub/PxnGoCommon/utils/flagz"
	UtilsRPC "github.com/PxnPub/PxnGoCommon/rpc"
	Worker   "github.com/PxnPub/pxnMetrics/shard/worker"
);



type AppShard struct {
	Version    string
	Service    *Service.Service
	BackLink   *UtilsRPC.Client
	BindPublic string
	BrokerAddr string
	Worker     *Worker.Worker
}



func New(version string) Service.App {
	return &AppShard{
		Version: version,
	};
}

func (app *AppShard) Main() {
	app.Service = Service.New();
	app.Service.Start();
	// flags
	var flag_bind   string;
	var flag_broker string;
	Flagz.String(&flag_bind,   "bind",   DefaultBindPublic   );
	Flagz.String(&flag_broker, "broker", DefaultBrokerAddress);
	Flag.Parse();
	app.BindPublic = flag_bind;
	app.BrokerAddr = flag_broker;
	// rpc to broker
	app.BackLink = UtilsRPC.NewClient(app.Service, app.BrokerAddr);
	// public listener
	app.Worker = Worker.New(app.Service, app.BindPublic);
//TODO
	app.Worker.ChecksumSeed = 9001;
//	app.Proc = Proc.NewProcessor(app.BackLink);
	// start things
	if err := app.BackLink.Start(); err != nil { Log.Panic(err); }
	if err := app.Proc.Start();     err != nil { Log.Panic(err); }
	app.Service.WaitUntilEnd();
}
