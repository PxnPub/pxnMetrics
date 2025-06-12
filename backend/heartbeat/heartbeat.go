package heartbeat;

import(
	Log      "log"
	Time     "time"
	Sync     "sync"
	Utils    "github.com/PxnPub/PxnGoCommon/utils"
	Service  "github.com/PxnPub/PxnGoCommon/service"
//	UpLink   "github.com/PxnPub/pxnMetrics/backend/uplink"
//	UtilsRPC "github.com/PxnPub/PxnGoCommon/utils/net/rpc"
);



type HeartBeat struct {
	Service   *Service.Service
//	UpLink    *UpLink.UpLink
	MuxState  Sync.Mutex
	TaskQueue chan Task
	NumShards uint8
	Shards    []ShardState
}

type ShardState struct {
	IsOnline  bool
	LastSeen  Time.Time
	LastBatch Time.Time
}



func New(service *Service.Service, num_shards uint8) *HeartBeat {
	return &HeartBeat{
		Service:   service,
		NumShards: num_shards,
		Shards:    make([]ShardState, num_shards),
	};
}

func (heart *HeartBeat) Start() error {
	heart.MuxState.Lock();
	defer heart.MuxState.Unlock();
	go heart.Serve();
	Utils.SleepC();
//	heart.UpLink.Start();
	Utils.SleepC();
	return nil;
}
//TODO: Close() also UpLink



func (heart *HeartBeat) Serve() {
	heart.Service.WaitGroup.Add(1);
	defer heart.Service.WaitGroup.Done();
	Log.Printf("Starting HeartBeat..");
	interval, _ := Time.ParseDuration("20ms");
	timer := Time.NewTicker(interval);
	var stopping uint8 = 0;
	LOOP_SERVE:
	for {
		SELECT_TASK:
		select {
		case task := <-heart.TaskQueue:
			heart.Handle(&task);
			LOOP_DRAIN:
			for {
				SELECT_DRAIN:
				select {
				case <-timer.C: break SELECT_DRAIN;
				default:        break LOOP_DRAIN;
				}
			}
		case <-timer.C:
			if heart.Service.IsStopping() {
				stopping++;
				if stopping > 10 {
					break LOOP_SERVE;
				}
			}
			break SELECT_TASK;
		}
	}
}

func (heart *HeartBeat) Handle(task *Task) {


print("TASK\n");




}
