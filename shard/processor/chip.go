package processor;

import(
	Time "time"
);



type Chip struct {
	FinalBatch   bool
	Timestamp    Time.Time
	// totals
	TotalServers uint64
	TotalPlayers uint64
	// per server
	Servers      map[uint64]ChipServer
}

type ChipServer struct {
	NumPlayers int16
}
