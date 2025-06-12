package heartbeat;

import(
	Fmt      "fmt"
	Context  "context"
	FrontAPI "github.com/PxnPub/pxnMetrics/api/front"
	WebAPI   "github.com/PxnPub/pxnMetrics/api/web"
);



type API_Front struct {
	FrontAPI.UnimplementedWebFrontAPIServer
	Heart *HeartBeat
}



func NewFrontAPI(heart *HeartBeat) *API_Front {
	return &API_Front{
		Heart: heart,
	};
}

func (api *API_Front) FetchStatusJSON(ctx Context.Context,
		_ *FrontAPI.Empty) (*FrontAPI.StatusJSON, error) {
	json_shards := make([]WebAPI.StatusShard, api.Heart.NumShards);
	for i, shard := range api.Heart.Shards {
		json_shards[i] = WebAPI.StatusShard{
			Name:      Fmt.Sprintf("Shard-%d", i),
			LastSeen:  uint32(shard.LastSeen.Seconds() ),
			LastBatch: uint32(shard.LastBatch.Seconds()),
//BatchWaiting: uint32 (v.BatchWaiting       ),
//QueueWaiting: uint32 (v.QueueWaiting       ),
//ReqPerSec:    float32(v.ReqPerSec          ),
//ReqPerMin:    float32(v.ReqPerMin          ),
//ReqPerHour:   float32(v.ReqPerHour         ),
//ReqPerDay:    float32(v.ReqPerDay          ),
//ReqTotal:     uint64 (v.ReqTotal           ),
		};
		if shard.IsOnline { json_shards[i].Status = "Online";
		} else {            json_shards[i].Status = "Offline" }
	}
	json, err := JSON.Marshal(
		WebAPI.StatusAPI{
			Shards: json_shards,
		},
	);
	if err != nil { return nil, err; }
	return &FrontAPI.StatusJSON{ Data: json }, nil;
}
