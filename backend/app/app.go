package app;

import(
	TrapC       "github.com/PxnPub/pxnGoUtils/trapc"
	ShardBroker "github.com/PxnPub/pxnMetrics/MetricsBackend/broker"
);



const NumShards = 1;
const BatchInterval = "5s";



func Main(trapc *TrapC.TrapC) {
	// TCP:99xx | shard rpc servers
	//broker :=
	ShardBroker.New(trapc, NumShards, BatchInterval);
	// UDP:9001 | frontend rpc server
}
