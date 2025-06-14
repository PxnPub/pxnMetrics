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
	Version    string
	Service    *Service.Service
	BackLink   *UtilsRPC.Client
	WebServer  *WebServer.WebServer
	Pages      *Pages.Pages
	BindAddr   string
	BrokerAddr string
}



func New(version string) Service.App {
	return &AppFrontend{
		Version: version,
	};
}

func (app *AppFrontend) Main() {
	app.Service = Service.New();
	app.Service.Start();
	// flags
	var flag_bind   string;
	var flag_broker string;
	Flagz.String(&flag_bind,   "bind",   WebServer.DefaultBindWeb);
	Flagz.String(&flag_broker, "broker", DefaultBrokerAddress    );
	Flag.Parse();
	app.BindAddr   = flag_bind;
	app.BrokerAddr = flag_broker;
	// rpc to broker
	app.BackLink = UtilsRPC.NewClient(app.Service, app.BrokerAddr);

//	client := FrontAPI.NewWebFrontAPIClient(app.BackLink.RPC);

	// web server
	app.WebServer = WebServer.NewWebServer(app.BindAddr);
	app.WebServer.WaitGroup = app.Service.WaitGroup;
	app.Service.AddCloseable(app.WebServer);
	app.Pages = Pages.New(app.WebServer.Router);
	// start things
	if err := app.BackLink.Start();  err != nil { Log.Panic(err); }
	app.Pages.Init(app.BackLink);
	if err := app.WebServer.Start(); err != nil { Log.Panic(err); }
	app.Service.WaitUntilEnd();
}
