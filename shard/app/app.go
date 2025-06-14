package app;
// pxnMetrics Shard App

import(
	Log      "log"
	Flag     "flag"
	Service  "github.com/PxnPub/PxnGoCommon/service"
	Flagz    "github.com/PxnPub/PxnGoCommon/utils/flagz"
	UtilsRPC "github.com/PxnPub/PxnGoCommon/rpc"
	Proc     "github.com/PxnPub/pxnMetrics/shard/processor"
);



type AppShard struct {
	Version    string
	Service    *Service.Service
	BackLink   *UtilsRPC.Client
	BindPublic string
	BrokerAddr string
	Proc       *Proc.Processor
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
	app.Worker = Worker.New();
	// start things
	if err := app.BackLink.Start(); err != nil { Log.Panic(err); }
	app.Worker.Init(app.BackLink);
	if err := app.Worker.Start();   err != nil { Log.Panic(err); }
	app.Service.WaitUntilEnd();
}
