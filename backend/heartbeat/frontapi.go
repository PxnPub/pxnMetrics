package heartbeat;

import(
	Log      "log"
	Fmt      "fmt"
	Context  "context"
	JSON     "encoding/json"
	Errors   "errors"
	UtilsRPC "github.com/PxnPub/PxnGoCommon/rpc"
	FrontAPI "github.com/PxnPub/pxnMetrics/api/front"
	WebAPI   "github.com/PxnPub/pxnMetrics/api/web"
);



type API_Front struct {
	FrontAPI.UnimplementedWebFrontAPIServer
	Heart     *HeartBeat
	UserMan   *UserManager
	NumShards uint8
	Checksum  uint16
}



func NewFrontAPI(heart *HeartBeat, userman *UserManager,
		num_shards uint8, checksum_init uint16) *API_Front {
	return &API_Front{
		Heart:     heart,
		UserMan:   userman,
		NumShards: num_shards,
		Checksum:  checksum_init,
	};
}

func (api *API_Front) FetchStatusJSON(ctx Context.Context,
		_ *FrontAPI.Empty) (*FrontAPI.StatusJSON, error) {
//TODO: move this to a function
	user, ok := ctx.Value(KeyUserPerms).(*User);
	if !ok {
		Log.Printf("Invalid RPC User type");
		return nil, Errors.New("Invalid RPC User type");
	}
	if !user.AllowWebCalls {
		username := ctx.Value(UtilsRPC.KeyUsername).(string);
		Log.Printf("User lacks permissions: %s", username);
		return nil, Errors.New("User lacks permissions");
	}
	json_shards := make([]WebAPI.StatusShard, api.Heart.NumShards);
	for i, shard := range api.Heart.Shards {
		var last_seen int64;
		if shard.LastSeen.IsZero() { last_seen = 0;
		} else { last_seen = shard.LastSeen.Unix(); }
		var last_batch int64;
		if shard.LastBatch.IsZero() { last_batch = 0;
		} else { last_batch = shard.LastBatch.Unix(); }
		json_shards[i] = WebAPI.StatusShard{
			Name:         Fmt.Sprintf("Shard-%d", i+1),
			LastSeen:     uint32 (last_seen         ),
			LastBatch:    uint32 (last_batch        ),
			BatchWaiting: uint32 (shard.BatchWaiting),
			QueueWaiting: uint32 (shard.QueueWaiting),
			ReqPerSec:    float32(shard.ReqPerSec   ),
			ReqPerMin:    float32(shard.ReqPerMin   ),
			ReqPerHour:   float32(shard.ReqPerHour  ),
			ReqPerDay:    float32(shard.ReqPerDay   ),
			ReqTotal:     uint64 (shard.ReqTotal    ),
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
