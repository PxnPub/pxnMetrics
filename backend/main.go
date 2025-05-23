package main;

import(
	Service "github.com/PxnPub/pxnGoUtils/service"
	App     "github.com/PxnPub/pxnMetrics/backend/app"
);



func main() {
	trapc := Service.Pre();
	App.Main(trapc);
	Service.Post(trapc);
}
