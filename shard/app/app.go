package app;

import(
	TrapC  "github.com/PxnPub/pxnGoUtils/trapc"
);



func Main(trapc *TrapC.TrapC) {

//TODO: flag?
var num_slivers uint8 = 2;
var shard_index uint8 = 1;
var shard_total uint8 = 2;

	//TCP:99xx client to broker
//	brokee := Broker.New(trapc, BatchInterval, shardindex);
//	brokee.Start();
	//UDP:9001
	shard := NewShard(trapc, num_slivers, shard_index, shard_total);
	shard.Start();
	print("Ready!\n");
}
