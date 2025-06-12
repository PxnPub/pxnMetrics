package app;
// minecraftmetrics.com

import(
	Log       "log"
	Flag      "flag"
	Service   "github.com/PxnPub/PxnGoCommon/service"
	Flagz     "github.com/PxnPub/PxnGoCommon/utils/flagz"
	WebServer "github.com/PxnPub/PxnGoCommon/utils/net/web"
	BackLink  "github.com/PxnPub/pxnMetrics/frontend/backlink"
	Pages     "github.com/PxnPub/pxnMetrics/frontend/pages"
);



type AppFrontend struct {
	Version  string
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
	var bind   string;
	var broker string;
	Flagz.String(&bind,   "bind",   WebServer.DefaultBindWeb);
	Flagz.String(&broker, "broker", DefaultBrokerAddress    );
	Flag.Parse();
	// rpc to broker
	backlink := BackLink.New(service, broker);

//	app.BackLink = UtilsRPC.NewBackLink(broker);





//	weblink := WebLink.New(service, broker);
	// web server
	webserv := WebServer.NewWebServer(bind);
	webserv.WaitGroup = service.WaitGroup;
	service.AddCloseable(webserv);
	Pages.New(webserv.Router, backlink);
	// start things
	if err := backlink.Start(); err != nil { Log.Panic(err); }
	if err := webserv.Start();  err != nil { Log.Panic(err); }
	service.WaitUntilEnd();
}
