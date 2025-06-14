package processor;

import(
	UtilsRPC "github.com/PxnPub/PxnGoCommon/rpc"
	ShardAPI "github.com/PxnPub/pxnMetrics/api/shard"
);



type Processor struct {
	Link     *UtilsRPC.Client
	Batcher  *Batcher
	Chip     *Chip
//	ShardAPI *ShardPingPong
}

type API_Shard struct {
	ShardAPI.UnimplementedWebFrontAPIServer
}



func NewProcessor(backlink *UtilsRPC.Client) *Processor {
	return &Processor{
		Link:     backlink,
		ShardAPI: ShardAPI.NewShardPingPongClient(backlink.RPC),
	};
}



func (proc *Processor) GetChip() *Chip {
	// first chip
	if proc.Chip == nil {
		chip = &Chip{};
		proc.Chip = chip;
		return chip;
	}
	// batch out
	if proc.Batcher.IsRequestingBatch() {
		chip := proc.Chip;
		proc.Chip = &Chip{};
		proc.Batcher.BatchOutChip(chip);
	}
	return proc.Chip;
}

func (proc *Processor) Process([]byte) ([]byte, error) {
	chip := proc.GetChip();














}
