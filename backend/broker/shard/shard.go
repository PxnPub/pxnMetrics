package shard;

import(
	Log      "log"
	Fmt      "fmt"
	Time     "time"
	Sync     "sync"
	Net      "net"
	RPC      "net/rpc"
	TrapC    "github.com/PxnPub/pxnGoUtils/trapc"
	UtilsNet "github.com/PxnPub/pxnGoUtils/net"
);



type Shard struct {
	Index     uint8
	Bind      string
	TrapC     *TrapC.TrapC
	StopChan  chan bool
	WaitGroup *Sync.WaitGroup
	Socket    Net.Listener
	Rpc       *RPC.Server
}



func New(trapc *TrapC.TrapC, index uint8, bind string, portbase int) (*Shard, error) {
	binding := Fmt.Sprintf("tcp://%s:%d", bind, portbase+int(index));
	listen, err := UtilsNet.NewSock(binding);
	if err != nil { return nil, err; }
	rpc := RPC.NewServer();
	shard := &Shard{
		Index:     index,
		Bind:      binding,
		TrapC:     trapc,
		StopChan:  trapc.NewStopChan(),
		WaitGroup: trapc.WaitGroup,
		Socket:    *listen,
		Rpc:       rpc,
	};
	rpc.Register(shard);
//TODO: remove this?
//	rpc.ServeConn(shard);
	return shard, nil;
}

func (shard *Shard) Close() error {
//print("CLOSE SHARD\n");
	shard.Socket.Close();
	return nil;
}



func (shard *Shard) Start() {
	go shard.Loop();
}

//TODO: what happens when the shard closes the socket?
func (shard *Shard) Loop() {
	shard.WaitGroup.Add(1);
	defer func() {
		shard.Close();
		shard.WaitGroup.Done();
	}();
	Log.Printf("[ Shard %d ] Listening: %s\n", shard.Index+1, shard.Bind);
	listentime := Time.Duration(200) * Time.Millisecond;
	LOOP:
	for {
		if shard.TrapC.IsStopping() {
			Log.Printf(" [ Shard %d ] Stopping listener..\n", shard.Index+1);
			break LOOP;
		}
		select {
		case stopping := <-shard.StopChan:
			if stopping {
				Log.Printf(" [ Shard %d ] Stopping listener..\n", shard.Index+1);
				break LOOP;
			}
		default:
		}
		shard.Socket.(*Net.TCPListener).
			SetDeadline(Time.Now().Add(listentime));
		conn, err := shard.Socket.Accept();
		if err == nil {
print("SHARD RPC START..\n");
			shard.Rpc.ServeConn(conn);
print("SHARD RPC END!\n");
		} else {
			neterr, ok := err.(Net.Error);
			if !ok || !neterr.Timeout() {
				Log.Printf("Connection failed: %v", err);
			}
		}
	}
}



type Batch struct {
	Data string
}

type BatchResult struct {
	Status int
}

func (shard *Shard) BatchOut(batch *Batch, result *BatchResult) error {




	result = &BatchResult{
		Status: 11,
	};
	return nil;
}
