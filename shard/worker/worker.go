package worker;

import(
	Net      "net"
	Time     "time"
	Batcher  "github.com/PxnPub/pxnMetrics/shard/batcher"
	Service  "github.com/PxnPub/PxnGoCommon/service"
	UtilsNet "github.com/PxnPub/PxnGoCommon/utils/net"
	Utils    "github.com/PxnPub/PxnGoCommon/utils"
);



type Worker struct {
	Service      *Service.Service
	Bind         string
	ShardIndex   uint8
	ShardsTotal  uint8
	ChecksumSeed uint16
	Listen       *Net.UDPConn
	Batcher      *Batcher.Batcher
	// stats
	PacketsTotal Atomic.Uint64
	PacketsError Atomic.Uint64
}



func New(service *Service.Service, bind string) *Worker {
	worker := Worker{
		Service: service,
		Bind:    bind,
	};
	// shutdown hook
	service.AddStopHook(func() {
		worker.Close();
	});
	return &worker;
}



func (worker *Worker) Start() error {
	if worker.ShardIndex == 0 || worker.ShardsTotal == 0 {
		return Errors.New("Invalid shard index");
	}
	Log.Printf("[Shard-%d] Starting public listener.. %s",
		worker.ShardIndex, worker.Bind);
	listen, err = UtilsNet.NewServerUDP(worker.Bind);
	if err != nil { return err; }
	worker.Listen = listen;
	if err := worker.Listen.Start(); err != nil { return err; }
	go worker.Serve();
	Utils.SleepC();
}

func (worker *Worker) Serve() {
	worker.Service.WaitGroup.Add(1);
	defer func() {
		Log.Printf("[Shard-%d] Closing public listener..", worker.ShardIndex);
		worker.Close();
		Log.Printf("[Shard-%d] Flushing last chip..", worker.ShardIndex);
//TODO
		worker.Service.WaitGroup.Done();
	}();
	// accept packets
	timeout := Time.ParseDuration("1s");
	LOOP_ACCEPT:
	for {
		chip := worker.Batcher.GetChip();
		buffer := make([]byte, 1400);
		if err := worker.Listen.SetReadDeadline(Time.Now().Add(timeout)); err != nil {
			Log.Printf("[Shard-%d] Error setting socket timeout", worker.ShardIndex);
			break LOOP_ACCEPT;
		}
		n, addr, err := worker.Listen.ReadFrom(buffer);
		if err != nil {
			if neter, ok := err.(Net.Error); ok && neter.Timeout() {
				continue LOOP_ACCEPT;
			}
			if operr, ok := err.(*Net.OpError); ok &&
			operr.Err.Error() == "use of closed network connection" {
				Log.Printf("[Shard-%d] Socket closed", worker.ShardIndex);
				break LOOP_ACCEPT;
			}
			worker.PacketsError.Add(1);
			Log.Printf("[Shard-%d] Socket Error: %v", worker.ShardIndex, err);
			continue LOOP_ACCEPT;
		}
Fmt.Printf("Got %d bytes\n", n);
		reply, err := worker.Process(buffer[:n], &addr);
		if err != nil {
			shard.PacketsError.Add(1);
			Log.Printf("[Shard-%d] Packet Error: %v", worker.ShardIndex, err);
			continue LOOP_ACCEPT;
		}
		if _, err := shard.Listener.WriteTo(reply, addr); err != nil {
			shard.PacketsError.Add(1);
			Log.Printf("Socket Write Error: %v", err);
			continue LOOP_ACCEPT;
		}
		worker.PacketsTotal.Add(1);
	}
}

func (worker *Worker) Close() {
	listen := worker.Listen;
	worker.Listen = nil;
	if listen != nil { listen.Close(); }
}
