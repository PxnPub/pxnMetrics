package routes;

import(
	Gorilla "github.com/gorilla/mux"
	Service "github.com/PxnPub/pxnGoUtils/service"
	Pages   "github.com/PxnPub/pxnMetrics/frontend/pages"
);



func Routes(mux *Gorilla.Router) {
	Service.AddRouteStatic(mux);
	mux.HandleFunc("/", Pages.HandleHome);
	mux.HandleFunc("/{name}/", Pages.HandleHome);
}
