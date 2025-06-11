package app;
// minecraftmetrics.com

import(
	Log       "log"
	Flag      "flag"
	Service   "github.com/PxnPub/PxnGoCommon/service"
	Flagz     "github.com/PxnPub/PxnGoCommon/utils/flagz"
	WebServer "github.com/PxnPub/PxnGoCommon/utils/net/web"
	Pages     "github.com/PxnPub/pxnMetrics/frontend/pages"
	WebLink   "github.com/PxnPub/pxnMetrics/frontend/weblink"
);



type AppFrontend struct {
	Version string
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
	var flag_broker string;
	Flagz.String(&flag_bind,   "bind",   WebServer.DefaultBindWeb);
	Flagz.String(&flag_broker, "broker", DefaultBrokerAddress    );
	Flag.Parse();
	// rpc to broker
	weblink := WebLink.New(service, flag_broker);
	// web server
	webserv := WebServer.NewWebServer(flag_bind);
	webserv.WaitGroup = service.WaitGroup;
	service.AddCloseable(webserv);
	Pages.New(webserv.Router, weblink)
	// start things
	if err := weblink.Start(); err != nil { Log.Panic(err); }
	if err := webserv.Start(); err != nil { Log.Panic(err); }
	service.WaitUntilEnd();
}
