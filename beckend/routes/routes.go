package routes

import (
	"github.com/gorilla/mux"
)

func RouteInit(r *mux.Router) {
	// isi routes disini
	AuthRoutes(r)
	ProductRoutes(r)
	CartRoutes(r)
	TransactionRoutes(r)
	UserRoutes(r)
}
