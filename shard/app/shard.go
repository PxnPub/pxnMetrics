package app;

import(
	Log     "log"
	Fmt     "fmt"
	Time    "time"
	Ctx     "context"
	StrConv "strconv"
	Atomic  "sync/atomic"
	TrapC   "github.com/PxnPub/pxnGoUtils/trapc"
	GNet    "github.com/panjf2000/gnet/v2"
	APIv1   "github.com/PxnPub/pxnMetrics/shard/apiv1"
);



type Shard struct {
	TrapC          *TrapC.TrapC
	NumSlivers     uint8
	ShardIndex     uint8
	ShardTotal     uint8
	Bind           string
	Engine         GNet.Engine
	ProcessorV1    *APIv1.Processor
	NumPackets     uint64
	ErrPackets     uint64
	LastNumPackets uint64
	LastErrPackets uint64
	BytesSent      uint64
}



func NewShard(trapc *TrapC.TrapC, num_slivers uint8, shard_index uint8, shard_total uint8) *Shard {
	proc1 := APIv1.NewProcessor();
	return &Shard{
		TrapC:       trapc,
		NumSlivers:  num_slivers,
		ShardIndex:  shard_index,
		ShardTotal:  shard_total,
//TODO
		Bind: "udp://127.0.0.1:9001",
		ProcessorV1: proc1,
	};
}



func (shard *Shard) Start() {
	Log.Printf("[Shard] Starting %d slivers.. %s", shard.NumSlivers, shard.Bind);
	// start listening
	go func() {
		shard.TrapC.WaitGroup.Add(1);
		defer shard.TrapC.WaitGroup.Done();
		// shutdown hook
		shard.TrapC.AddStopHook(func() {
			shard.Close();
		});
		// start listening
		if err := GNet.Run(
			shard, shard.Bind,
			GNet.WithMulticore(true),
			GNet.WithReusePort(true),
//TODO: is this needed?
			GNet.WithTicker(true),
//TODO: number of threads
			GNet.WithNumEventLoop(20),
			GNet.WithLoadBalancing(GNet.RoundRobin),
		); err != nil {
			Log.Printf("Listener error: %v\n", err);
		}
	}();
}

func (shard *Shard) Close() {
	if shard.Engine != (GNet.Engine{}) {
		shard.Engine.Stop(Ctx.Background());
	}
}



func (shard *Shard) OnBoot(engine GNet.Engine) GNet.Action {
	shard.Engine = engine;
print("BOOT\n");
	return GNet.None;
}

func (shard *Shard) OnShutdown(engine GNet.Engine) {
print("SHUTDOWN\n");
//TODO: submit last chip here?
}



func (shard *Shard) OnTraffic(conn GNet.Conn) GNet.Action {
	// get packet data
	Atomic.AddUint64(&shard.NumPackets, 1);
	data, err := conn.Next(-1);
	if err != nil {
		Atomic.AddUint64(&shard.ErrPackets, 1);
		Log.Printf("Packet error: %v\n", err);
		return GNet.Close;
	}
	// validate json
	submit, reply, err := shard.ProcessorV1.Validate(data);
	if err != nil {
		Atomic.AddUint64(&shard.ErrPackets, 1);
		Log.Printf("Validate error: %v\n", err);
		return GNet.Close;
	}
	// send reply
	sent, err := conn.Write(reply);
	Atomic.AddUint64(&shard.BytesSent, uint64(sent));
	if err != nil {
		Atomic.AddUint64(&shard.ErrPackets, 1);
		Log.Printf("Process error: %v\n", err);
		return GNet.Close;
	}
	// process json
	if err := shard.ProcessorV1.Process(submit); err != nil {
		Atomic.AddUint64(&shard.ErrPackets, 1);
		Log.Printf("Process error: %v\n", err);
		return GNet.Close;
	}
	return GNet.Close;
}

func (shard *Shard) OnTick() (Time.Duration, GNet.Action) {
	cnt_num := shard.NumPackets;
	cnt_err := shard.ErrPackets;
	Fmt.Printf(
		" %s per sec (%s errors)\n",
		Format(int64(cnt_num - shard.LastNumPackets)),
		Format(int64(cnt_err - shard.LastErrPackets)),
	);
	shard.LastNumPackets = cnt_num;
	shard.LastErrPackets = cnt_err;
	return Time.Second, GNet.None;
}



// only needed with tcp, not used with udp
func (shard *Shard) OnOpen(conn GNet.Conn) ([]byte, GNet.Action) { return nil, GNet.None; }
func (shard *Shard) OnClose(conn GNet.Conn, err error) GNet.Action { return GNet.None; }



//TODO: move this
func Format(n int64) string {
    in := StrConv.FormatInt(n, 10)
    numOfDigits := len(in)
    if n < 0 {
        numOfDigits-- // First character is the - sign (not a digit)
    }
    numOfCommas := (numOfDigits - 1) / 3

    out := make([]byte, len(in)+numOfCommas)
    if n < 0 {
        in, out[0] = in[1:], '-'
    }

    for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
        out[j] = in[i]
        if i == 0 {
            return string(out)
        }
        if k++; k == 3 {
            j, k = j-1, 0
            out[j] = ','
        }
    }
}
