package pages;

import(
	HTTP     "net/http"
	Template "html/template"
	HTML     "github.com/PxnPub/PxnGoCommon/utils/html"
);



func (pages *Pages) PageWeb_Global(out HTTP.ResponseWriter, in *HTTP.Request) {
	HTML.SetContentType(out, "html");
	build := pages.GetBuilder();
//TODO
build.IsDev = true;
	tpl, err := Template.ParseFiles(
		"html/main.tpl",
		"html/pages/global.tpl",
	);
	if err != nil { panic(err); }
	vars := struct {
		Page  string
		Title string
	}{
		Page:  "global",
		Title: "title",
	};
	out.Write([]byte(build.RenderTop()));
	tpl.ExecuteTemplate(out, "main.tpl", vars);
	tpl.ExecuteTemplate(out, "global.tpl", vars);
	out.Write([]byte(build.RenderBottom()));
}



func (pages *Pages) PageAPI_Global(out HTTP.ResponseWriter, in *HTTP.Request) {
	HTML.SetContentType(out, "json");
//	url, err := URL.ParseQuery(in.URL.RawQuery);
//	if err != nil { panic(err); }
	out.Write([]byte("{}"));
}
