package main;

import(
	Service "github.com/PxnPub/pxnGoUtils/service"
	App     "github.com/PxnPub/pxnMetrics/shard/app"
);



func main() {
	trapc := Service.Pre();
	App.Main(trapc);
	Service.Post(trapc);
}
