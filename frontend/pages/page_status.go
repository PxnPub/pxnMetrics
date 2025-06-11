package pages;

import(
	HTTP     "net/http"
	Template "html/template"
	UtilsWeb "github.com/PxnPub/pxnGoUtils/utils/web"
);



func (pages *Pages) PageWeb_Status(out HTTP.ResponseWriter, in *HTTP.Request) {
	UtilsWeb.SetContentType(out, "html");
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
	UtilsWeb.SetContentType(out, "json");
	// fetch from broker
	data := pages.WebLink.FetchStatusJSON();
	out.Write(data);
}
