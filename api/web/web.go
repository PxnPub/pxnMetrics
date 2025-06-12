package web;

import(
	Time "time"
);



type StatusShard struct {
	Name         string
	LastSeen     Time.Time
	LastBatch    Time.Time
	BatchWaiting uint32
	QueueWaiting uint32
	ReqPerSec    float32
	ReqPerMin    float32
	ReqPerHour   float32
	ReqPerDay    float32
	ReqTotal     uint64
}
