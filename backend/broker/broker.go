package broker;

import(
	Log   "log"
	Time  "time"
	TrapC "github.com/PxnPub/pxnGoUtils/trapc"
	Shard "github.com/PxnPub/pxnMetrics/backend/broker/shard"
);



const ShardPortBase = 9901;



type Broker struct {
	TrapC         *TrapC.TrapC
	NumShards     uint8
	IntervalBatch Time.Duration
	Shards        []*Shard.Shard
}



func New(trapc *TrapC.TrapC, num_shards uint8, batch_interval string, bind string) *Broker {
	interval, err := Time.ParseDuration(batch_interval);
	if err != nil { panic(err); }
	// shards
	Log.Printf("[Broker] Starting %d shards..", num_shards);
	shards := make([]*Shard.Shard, num_shards);
	for index:=uint8(0); index<num_shards; index++ {
		shard, err := Shard.New(trapc, index, bind, ShardPortBase);
		if err != nil { panic(err); }
		shards[index] = shard;
	}
	return &Broker{
		TrapC:         trapc,
		NumShards:     num_shards,
		IntervalBatch: interval,
		Shards:        shards,
	};
}

func (broker *Broker) Start() {
	sleep, err := Time.ParseDuration("1s");
	if err != nil { panic(err); }
	size := int16(broker.NumShards);
	for index:=int16(-3); index<size; index++ {
		if broker.TrapC.IsStopping() { return; }
		if index == -1 { print("\n"); }
		Time.Sleep(sleep);
		if index >= 0 {
			broker.Shards[index].Start();
		}
	}
}



func (broker *Broker) LogState(msg string) {
//TODO: "[_1_|#2#|_3_|#4#|#5#] msg"
}
