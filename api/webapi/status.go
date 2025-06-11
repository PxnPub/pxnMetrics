package webapi;



type JSON_Status struct {
	Shards []JSON_StatusShard `json:"Shards"`
}

type JSON_StatusShard struct {
	Name         string  `json:"Name"`
	Status       string  `json:"Status"`
	LastBatch    uint32  `json:"LastBatch"`
	BatchWaiting uint32  `json:"BatchWaiting"`
	QueueWaiting uint32  `json:"QueueWaiting"`
	ReqPerSec    float32 `json:"ReqPerSec"`
	ReqPerMin    float32 `json:"ReqPerMin"`
	ReqPerHour   float32 `json:"ReqPerHour"`
	ReqPerDay    float32 `json:"ReqPerDay"`
	ReqTotal     uint64  `json:"ReqTotal"`
}
