package submitapi;
// mcserver to shard

import(
	_ "encoding/json"
);



type Submit struct {
	Timestamp  int64  `json:"Timestamp"`
	ServerUID  uint64 `json:"ServerUID"`
	Platform   string `json:"Platform"`
	NumPlayers int16  `json:"NumPlayers"`
}

type SubmitReply struct {
	Status uint8 `json:"Status"`
}



type Chip struct {
	TotalServers uint64
	TotalPlayers uint64
	Servers      map[uint64]ChipServer
}

type ChipServer struct {
	NumPlayers int16
}
