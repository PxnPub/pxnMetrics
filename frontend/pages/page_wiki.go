package pages;

import(
	HTTP      "net/http"
	Template  "html/template"
	HTML      "github.com/PxnPub/PxnGoCommon/utils/html"
);



func (pages *Pages) PageWeb_Wiki(out HTTP.ResponseWriter, in *HTTP.Request) {
	HTML.SetContentType(out, "html");
	build := pages.GetBuilder();
//TODO
build.IsDev = true;
	tpl, err := Template.ParseFiles(
		"html/main.tpl",
	);
	if err != nil { panic(err); }
	vars := struct {
		Page  string
		Title string
	}{
		Page:  "wiki",
		Title: "title",
	};
	out.Write([]byte(build.RenderTop()));
	tpl.ExecuteTemplate(out, "main.tpl", vars);
//	tpl.ExecuteTemplate(out, "wiki.tpl", vars);
	out.Write([]byte("Wiki goes here"));
	out.Write([]byte(build.RenderBottom()));
}
