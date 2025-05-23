package app;

import(
	TrapC       "github.com/PxnPub/pxnGoUtils/trapc"
	ShardBroker "github.com/PxnPub/pxnMetrics/backend/broker"
);



const NumShards uint8 = 2;
const BatchInterval = "5s";
const ShardBind = "127.0.0.1";



func Main(trapc *TrapC.TrapC) {
	// TCP:99xx shard broker
	broker := ShardBroker.New(trapc, NumShards, BatchInterval, ShardBind);
	broker.Start();
}
