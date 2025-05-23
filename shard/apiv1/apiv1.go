package apiv1;

import(
Fmt "fmt"
	Atomic "sync/atomic"
	JSON   "encoding/json"
	API    "github.com/PxnPub/pxnMetrics/api/submitapi"
);



type Processor struct {
	Chip *API.Chip
}



func NewProcessor() *Processor {
	return &Processor{
		Chip: NewChip(),
	};
}

func NewChip() *API.Chip {
	return &API.Chip{
	};
}



func (proc *Processor) Validate(data []byte) (*API.Submit, []byte, error) {
	var api API.Submit;
	if err := JSON.Unmarshal(data, &api); err != nil {
		return nil, nil, err;
	}
//TODO: check values
	// check timestamp
	// check num players
	reply := API.SubmitReply{
		Status: 11,
	};
	json, err := JSON.Marshal(reply);
	if err != nil {
		return nil, nil, err;
	}
	return &api, json, nil;
}



func (proc *Processor) Process(submit *API.Submit) error {
	// total servers
	Atomic.AddUint64(&proc.Chip.TotalServers, 1);
	// total players
	Atomic.AddUint64(&proc.Chip.TotalPlayers, uint64(submit.NumPlayers));
if proc.Chip.TotalServers % 1000 == 0 {
Fmt.Printf("Total Servers: %d  Total Players: %d\n",
proc.Chip.TotalServers, proc.Chip.TotalPlayers);
}
	return nil;
}
