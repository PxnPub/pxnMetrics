package heartbeat;

import(
	UtilsNet "github.com/PxnPub/PxnGoCommon/utils/net"
);



type TaskType uint8;
const(
	TaskType_SyncIP TaskType = iota
	TaskType_Batch
);



type Task struct {
	TaskType TaskType
	Task_IPDB  *TaskData_IPDB
	Task_Batch *TaskData_Batch
}

type TaskData_IPDB struct {
	Updates map[UtilsNet.TupleIP]int32
	IP        UtilsNet.TupleIP
	AddTokens uint16
}

type TaskData_Batch struct {
//TODO
}



func (heart *HeartBeat) QueueSyncIP() {
//TODO
list := TaskData_IPDB{};
	task := Task{
		TaskType:  TaskType_SyncIP,
		Task_IPDB: &list,
	};
	heart.TaskQueue <- task;
}

func (heart *HeartBeat) QueueBatch() {
//TODO
batch := TaskData_Batch{};
	task := Task{
		TaskType:   TaskType_Batch,
		Task_Batch: &batch,
	};
	heart.TaskQueue <- task;
}
