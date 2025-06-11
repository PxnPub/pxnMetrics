package uplink;

import(
	Sync     "sync"
	Errors   "errors"
	Service  "github.com/PxnPub/PxnGoCommon/service"
	UtilsRPC "github.com/PxnPub/PxnGoCommon/utils/net/rpc"
	FrontAPI "github.com/PxnPub/pxnMetrics/api/front"
);



type UpLink struct{
	Service   *Service.Service
	MuxState  Sync.Mutex
	NumShards uint8
	Bind      string
	RPC       *UtilsRPC.UpLink
}



func New(service *Service.Service, num_shards uint8, bind string) *UpLink {
	return &UpLink{
		Service:   service,
		NumShards: num_shards,
		Bind:      bind,
		RPC:       UtilsRPC.NewUpLink(bind),
	};
}



func (uplink *UpLink) Start() error {
	uplink.MuxState.Lock();
	defer uplink.MuxState.Unlock();
	if uplink.Bind == "" { return Errors.New("Bind address is required"); }
	// register api
	FrontAPI.RegisterWebFrontAPIServer(uplink.RPC.Server, &API_Status{});
	return uplink.RPC.Start();
}

func (uplink *UpLink) Stop() {
	uplink.MuxState.Lock();
	defer uplink.MuxState.Unlock();
	uplink.RPC.Close();
}
