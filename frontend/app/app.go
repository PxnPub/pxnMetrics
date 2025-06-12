package app;
// minecraftmetrics.com

import(
	Log       "log"
	Flag      "flag"
	Service   "github.com/PxnPub/PxnGoCommon/service"
	Flagz     "github.com/PxnPub/PxnGoCommon/utils/flagz"
	WebServer "github.com/PxnPub/PxnGoCommon/utils/net/web"
	UtilsRPC  "github.com/PxnPub/PxnGoCommon/rpc"
//	FrontAPI  "github.com/PxnPub/pxnMetrics/api/front"
	Pages     "github.com/PxnPub/pxnMetrics/frontend/pages"
);



type AppFrontend struct {
	Version  string
	Service  *Service.Service
	BackLink *UtilsRPC.Client
	WebServe *WebServer.WebServer
	Bind     string
	Remote   string
}



func New(version string) Service.App {
	return &AppFrontend{
		Version: version,
	};
}

func (app *AppFrontend) Main() {
	service := Service.New();
	service.Start();
	// flags
	var flag_bind   string;
	var flag_remote string;
	Flagz.String(&flag_bind,   "bind",   WebServer.DefaultBindWeb);
	Flagz.String(&flag_remote, "broker", DefaultBrokerAddress    );
	Flag.Parse();
	app.Bind   = flag_bind;
	app.Remote = flag_remote;
	// rpc to broker
	app.BackLink = UtilsRPC.NewClient(service, app.Remote);

//	client := FrontAPI.NewWebFrontAPIClient(app.BackLink.RPC);



	// web server
	app.WebServe = WebServer.NewWebServer(app.Bind);
	app.WebServe.WaitGroup = service.WaitGroup;
	service.AddCloseable(app.WebServe);
	Pages.New(app.WebServe.Router, app.BackLink);
	// start things
	if err := app.BackLink.Start(); err != nil { Log.Panic(err); }
	if err := app.WebServe.Start();  err != nil { Log.Panic(err); }
	service.WaitUntilEnd();
}
