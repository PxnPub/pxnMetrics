package heartbeat;

import(
	Fmt      "fmt"
	Context  "context"
	JSON     "encoding/json"
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
			Name:         Fmt.Sprintf("Shard-%d", i),
			LastSeen:     uint32 (shard.LastSeen.Unix() ),
			LastBatch:    uint32 (shard.LastBatch.Unix()),
			BatchWaiting: uint32 (shard.BatchWaiting    ),
			QueueWaiting: uint32 (shard.QueueWaiting    ),
			ReqPerSec:    float32(shard.ReqPerSec       ),
			ReqPerMin:    float32(shard.ReqPerMin       ),
			ReqPerHour:   float32(shard.ReqPerHour      ),
			ReqPerDay:    float32(shard.ReqPerDay       ),
			ReqTotal:     uint64 (shard.ReqTotal        ),
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
	return &FrontAPI.StatusJSON{ Data: string(json) }, nil;
}
