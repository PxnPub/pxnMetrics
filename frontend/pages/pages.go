package pages;

import(
	Gorilla   "github.com/gorilla/mux"
	HTML      "github.com/PxnPub/PxnGoCommon/utils/html"
	WebServer "github.com/PxnPub/PxnGoCommon/utils/net/web"
	WebLink   "github.com/PxnPub/pxnMetrics/frontend/weblink"
);



type Pages struct {
	WebLink *WebLink.WebLink
}



func New(router *Gorilla.Router, weblink *WebLink.WebLink) *Pages {
	pages := Pages{
		WebLink: weblink,
	};
	WebServer.AddStaticRoute(router);
	router.HandleFunc("/",            pages.PageWeb_Global);
	router.HandleFunc("/wiki/",       pages.PageWeb_Wiki  );
	router.HandleFunc("/status/",     pages.PageWeb_Status);
	router.HandleFunc("/api/status/", pages.PageAPI_Status);
	router.HandleFunc("/about/",      pages.PageWeb_About );
	return &pages;
}



func (pages *Pages) GetBuilder() *HTML.Builder {
	return HTML.NewBuilder().
		WithBootstrap().
		WithBootstrapIcons().
		SetFavIcon("/static/line-chart.ico").
		AddCSS("/static/metrics.css").
		SetTitle("pxnMetrics");
}
