package weblink;

import(
//	Log      "log"
	Fmt      "fmt"
	Sync     "sync"
	Errors   "errors"
	Service  "github.com/PxnPub/PxnGoCommon/service"
	UtilsRPC "github.com/PxnPub/PxnGoCommon/utils/net/rpc"
);



type WebLink struct {
	Service    *Service.Service
	MuxState   Sync.Mutex
	BrokerAddr string
	RPC        *UtilsRPC.BackLink
}



func New(service *Service.Service, broker string) *WebLink {
	return &WebLink{
		Service:    service,
		BrokerAddr: broker,
	};
}



func (link *WebLink) Start() error {
	link.MuxState.Lock();
	defer link.MuxState.Unlock();
	if link.BrokerAddr == "" { return Errors.New("Broker address is required"); }
	link.RPC = UtilsRPC.NewBackLink(link.BrokerAddr);
	return link.RPC.Start();
}



func (link *WebLink) Close() {
	link.RPC.Close();
}



func (link *WebLink) FetchStatusJSON() []byte {
	var reply string;
	err := link.RPC.Client.Call("API_Status.FetchStatusJSON", "request", &reply);
	if err != nil { panic(err); }
Fmt.Printf("REPLY: %s\n", reply);
	return []byte(reply);
}
//	data, err := link.RPC.Call("FetchStatusJSON");
//	if err != nil { Log.Printf("%s%s in FetchStatusJSON()",
//		UtilsRPC.LogPrefix, err); }
//	return []byte(data);
//	return data.([]byte);
