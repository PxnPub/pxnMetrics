package pages;

import(
	Fmt      "fmt"
	HTTP     "net/http"
	Template "html/template"
	Context  "context"
	JSON     "encoding/json"
	Runtime  "runtime"
	HTML     "github.com/PxnPub/PxnGoCommon/utils/html"
	// api
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
	api := FrontAPI.NewWebFrontAPIClient(pages.Link.RPC);
	result, err := api.FetchStatusJSON(Context.Background(), &FrontAPI.Empty{});
	if err != nil {
		trace := make([]byte, 1024);
		n := Runtime.Stack(trace, true);
		HTTP.Error(out,
			Fmt.Sprintf("%s\n\n%s", err.Error(), trace[:n]),
			HTTP.StatusInternalServerError,
		);
		return;
	}
	out.WriteHeader(HTTP.StatusOK)
	JSON.NewEncoder(out).Encode(result.Data);
}
