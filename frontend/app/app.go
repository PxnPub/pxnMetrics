package app;

import(
	TrapC   "github.com/PxnPub/pxnGoUtils/trapc"
	Service "github.com/PxnPub/pxnGoUtils/service"
	Routes  "github.com/PxnPub/pxnMetrics/frontend/routes"
);



func Main(trapc *TrapC.TrapC) {
	service := Service.NewWebServer(trapc, "127.0.0.1:8888");
	// routes
	Routes.Routes(service.Mux);
	// start listening
	service.Start();
}
