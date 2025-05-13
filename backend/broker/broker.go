package broker;

import(
//	Log      "log"
	Fmt      "fmt"
	Time     "time"
	Net      "net"
	RPC      "net/rpc"
	UtilsNet "github.com/PxnPub/pxnGoUtils/net"
	ShardAPI "github.com/PxnPub/pxnMetrics/MetricsBackendAPI/api/shardapi"
);



type ShardBroker struct {
	NumShards     int
	IntervalBatch Time.Duration
	IntervalShard Time.Duration
	Shards        []ShardServer
}

type ShardServer struct {
	Index    int
	Api      *ShardAPI.ShardAPI
	Listener Net.Listener
	Rpc      RPC.Server
}



func New(num_shards int, interval_str string) *ShardBroker {
	interval_batch, err := Time.ParseDuration(interval_str);
	if err != nil { panic(err); }
	interval_shard := interval_batch / Time.Duration(num_shards);
	// shard rpc listeners
	shards := make([]ShardServer, num_shards);
	for index:=0; index<num_shards; index++ {
		shards[index] = *NewShard(index);
	}
	return &ShardBroker{
		NumShards:     num_shards,
		IntervalBatch: interval_batch,
		IntervalShard: interval_shard,
		Shards:        shards,
	};
}

func NewShard(index int) *ShardServer {
	bind := Fmt.Sprintf("tcp://127.0.0.1:%d", 9900+index);
	listen, err := UtilsNet.NewSock(bind);
	if err != nil { panic(err); }
	rpc := RPC.NewServer();
	api := ShardAPI.ShardAPI{};
	rpc.Register(api);
	shard := ShardServer{
		Index:    index,
		Api:      *api,
		Listener: *listen,
		Rpc:      *rpc,

	};
	go shard.Loop();
	return &shard;
}



func (shard *ShardServer) Loop() {
	defer shard.Listener.Close();








}
