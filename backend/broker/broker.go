package broker;

import(
	Log      "log"
	Fmt      "fmt"
	Time     "time"
	Net      "net"
	RPC      "net/rpc"
	Sync     "sync"
	UtilsNet "github.com/PxnPub/pxnGoUtils/net"
	TrapC    "github.com/PxnPub/pxnGoUtils/trapc"
	ShardAPI "github.com/PxnPub/pxnMetrics/api/shardapi"
);



const ShardBasePort = 9901;



type ShardBroker struct {
	NumShards     int
	IntervalBatch Time.Duration
	IntervalShard Time.Duration
	Shards        []ShardServer
}

type ShardServer struct {
	Index    int
	Bind     string
	StopChan chan bool
	WaitGrp  *Sync.WaitGroup
	Socket   Net.Listener
	Rpc      *RPC.Server
}



func New(trapc *TrapC.TrapC, num_shards int, interval_str string) *ShardBroker {
	interval_batch, err := Time.ParseDuration(interval_str);
	if err != nil { panic(err); }
	interval_shard := interval_batch / Time.Duration(num_shards);
	// shard rpc listeners
	shards := make([]ShardServer, num_shards);
	for index:=0; index<num_shards; index++ {
		shards[index] = *NewShard(trapc, index);
	}
	return &ShardBroker{
		NumShards:     num_shards,
		IntervalBatch: interval_batch,
		IntervalShard: interval_shard,
		Shards:        shards,
	};
}

func NewShard(trapc *TrapC.TrapC, index int) *ShardServer {
	bind := Fmt.Sprintf("tcp://127.0.0.1:%d", ShardBasePort+index);
	listen, err := UtilsNet.NewSock(bind);
	if err != nil { panic(err); }
	shard := ShardServer{
		Index:    index,
		Bind:     bind,
		StopChan: trapc.NewStopChan(),
		WaitGrp:  trapc.WaitGrp,
		Socket:   *listen,
		Rpc:      RPC.NewServer(),
	};
	shard.Rpc.Register(&shard);
	go shard.Loop();
	return &shard;
}



func (shard *ShardServer) Loop() {
	shard.WaitGrp.Add(1);
	defer func() {
		shard.Socket.Close();
		shard.WaitGrp.Done();
	}();
	Log.Printf("[ API %d ] Listening: %s\n", shard.Index, shard.Bind);
	listentimeout := Time.Duration(200) * Time.Millisecond;
	for {
		select {
		case stopping := <-shard.StopChan:
			if stopping {
				Log.Printf(" [ API %d ] Stopping listener..", shard.Index);
				return;
			}
		default:
		}
		shard.Socket.(*Net.TCPListener).
			SetDeadline(Time.Now().Add(listentimeout));
		conn, err := shard.Socket.Accept();
		if err == nil {
print("SHARD RPC START..");
			shard.Rpc.ServeConn(conn);
print("SHARD RPC END!");
		} else {
			neterr, ok := err.(Net.Error);
			if !ok || !neterr.Timeout() {
				Log.Printf("Connection failed: %v", err);
			}
		}
	}


}



func (shard *ShardServer) Query(request ShardAPI.Query, reply *ShardAPI.QueryReply) error {

print("QUERY");




	return nil;
}
