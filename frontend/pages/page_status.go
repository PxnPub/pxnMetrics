package pages;

import(
//	Context  "context"
	HTTP     "net/http"
	Template "html/template"
	Context  "context"
	HTML     "github.com/PxnPub/PxnGoCommon/utils/html"
	// api
//	FrontAPI "github.com/PxnPub/pxnMetrics/api/front"
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
	result, err := pages.Link.Client.FetchStatusJSON(
		Context.Background(),
		nil,
	);
	if err != nil { panic(err); }
	out.Write([]byte(result.Data));
}
//	api := FrontAPI.NewWebFrontAPIClient(pages.BackLink.Client);
//api.FetchStatusJSON(Context.Background(), &FrontAPI.Empty{});
//	reply, err := api.FetchStatusJSON(Context.Background(), &FrontAPI.Empty{});
//	if err != nil { panic(err); }
//	out.Write([]byte(reply.Data));

//	out.Write(pages.BackLink.Client.Call);
//	out.Write(pages.BackLink.Client.FetchStatusJSON());
