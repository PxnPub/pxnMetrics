package heartbeat;

import(
//	Log  "log"
//	Fmt  "fmt"
//	Time "time"
//	JSON "encoding/json"
	// api
//	FrontAPI "github.com/PxnPub/pxnMetrics/api/front"
);



func (heart *HeartBeat) FetchStatusJSON() string {


return "OK";

//	now := Time.Now();
//TODO
//	timeout_online, _ := Time.ParseDuration("5s");
//	// build json
//	shards := make([]API.JSON_StatusShard, heart.NumShards);
//	for index:=uint8(0); index<heart.NumShards; index++ {
//		// timeout online
//		if heart.Shards[index].IsOnline {
//			since := now.Sub(heart.Shards[index].LastSeen);
//			if since > timeout_online {
//				heart.Shards[index].IsOnline = false;
//			}
//		}
//		shards[index].Name = Fmt.Sprintf("Shard %d", index+1);
//		if heart.Shards[index].IsOnline { shards[index].Status = "Online";
//		} else {                          shards[index].Status = "Offline"; }
//	}
//	data, err := JSON.Marshal(API.JSON_Status{
//		Shards: shards,
//	});
//	if err != nil { Log.Printf("%v in FetchStatusJSON()", err); return nil; }
//	return data;
}
