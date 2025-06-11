package uplink;

import(
//	Fmt      "fmt"
	Context  "context"
	FrontAPI "github.com/PxnPub/pxnMetrics/api/front"
);



type API_Front struct {
	FrontAPI.UnimplementedWebFrontAPIServer
}



func (api *API_Front) FetchStatusJSON(ctx Context.Context,
		_ *FrontAPI.Empty) (*FrontAPI.StatusJSON, error) {
	out := FrontAPI.StatusJSON{
		Data: "{}",
	};
	return &out, nil;
}
//a := "Test Works!";
//reply = &a;
//Fmt.Printf("Request:: %s\n", request);
//Fmt.Printf("REPLY:: %s\n", *reply);
