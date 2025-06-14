package pages;

import(
	Fmt      "fmt"
	HTTP     "net/http"
	Template "html/template"
	Context  "context"
//	JSON     "encoding/json"
	Runtime  "runtime"
	GRPC     "google.golang.org/grpc"
	GZIP     "google.golang.org/grpc/encoding/gzip"
	HTML     "github.com/PxnPub/PxnGoCommon/utils/html"
	FrontAPI "github.com/PxnPub/pxnMetrics/api/front"
);



func (pages *Pages) PageWeb_Status(out HTTP.ResponseWriter, in *HTTP.Request) {
	HTML.SetContentType(out, "html");
	build := pages.GetBuilder().
		AddBotJS("/static/status.js");
//TODO
build.IsDev = true;
	tpl, err := Template.ParseFiles(
		"html/main.tpl",
		"html/pages/status.tpl",
	);
	if err != nil { panic(err); }
	vars := struct {
		Page  string
		Title string
	}{
		Page:  "status",
		Title: "title",
	};
	out.Write([]byte(build.RenderTop()));
	tpl.ExecuteTemplate(out, "main.tpl", vars);
	tpl.ExecuteTemplate(out, "status.tpl", vars);
	out.Write([]byte(build.RenderBottom()));
}



func (pages *Pages) PageAPI_Status(out HTTP.ResponseWriter, in *HTTP.Request) {
	HTML.SetContentType(out, "json");
	result, err := pages.FrontAPI.FetchStatusJSON(
		Context.Background(),
		&FrontAPI.Empty{},
//TODO: optional? only when not unix socket?
		GRPC.UseCompressor(GZIP.Name),
	);
//TODO: make this into a function?
	if err != nil {
		trace := make([]byte, 1024);
		n := Runtime.Stack(trace, true);
		HTTP.Error(out,
			Fmt.Sprintf("%s\n\n%s", err.Error(), trace[:n]),
			HTTP.StatusInternalServerError,
		);
		return;
	}
	out.WriteHeader(HTTP.StatusOK);
	out.Write(result.Data);
//	JSON.NewEncoder(out).Encode(result.Data);
}
