package processor;

import(
	Atomic "sync/atomic"
);



const BatchBufferSize = 1000;



//TODO: max time before batchout (in case broker is down)
type Batcher struct {
	ChipChan chan Chip
	Requesting uint8
}



func NewBatcher() *Batcher {
	return &Batcher{
		ChipChan: make(chan Chip, BatchBufferSize),
	};
}



func (batcher *Batcher) BatchOutChip(chip *Chip) {
	batcher.ChipChan <- chip;
	Atomic.StoreUint8(batcher.IsBatching, 0);
}

func (batcher *Batcher) GetBatchChip() *Chip {
	Atomic.StoreUint8(batcher.Requesting, 1);
	chip := <- batcher.ChipChan;
	return chip;
}

func (batcher *Batcher) IsRequestingBatch() bool {
	return (Atomic.Load(batcher.Requesting) != 0);
}
