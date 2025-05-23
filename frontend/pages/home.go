package pages;

import(
//	Fmt     "fmt"
	HTTP "net/http"
	TPL  "html/template"
	HTML "github.com/PxnPub/pxnGoUtils/html"
//	Gorilla "github.com/gorilla/mux"
);



func HandleHome(out HTTP.ResponseWriter, in *HTTP.Request) {
	title := "Home";
	build := HTML.NewBuilder();
	build.AddBootstrap();
	build.AddCSS("/static/metrics.css");
	build.SetTitle(title);
	tpl, err := TPL.ParseFiles(
		"html/main.tpl",
		"html/home.tpl",
	);
	if err != nil { panic(err); }
	vars := struct {
		Page  string
		Title string
	}{
		Page:  "home",
		Title: title,
	};
	out.Header().Set("Content-Type", "text/html");
	out.Write([]byte(build.RenderTop()));
	tpl.Execute(out, vars);
	out.Write([]byte(build.RenderBottom()));
}
//	vars := Gorilla.Vars(in);
//	query := in.URL.Query();
//	contents := Fmt.Sprintf("Hello %s %s", query.Get("name"), vars["name"]);
