package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/middleware"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/gorilla/mux"
)

func CartRoutes(r *mux.Router) {
	cartRepository := repositories.RepositoryCart(mysql.DB)
	h := handlers.HandlerCart(cartRepository)

	r.HandleFunc("/cart", middleware.Auth(h.CreateCart)).Methods("POST")
	r.HandleFunc("/cartuser", middleware.Auth(h.FindCartId)).Methods("GET")
	r.HandleFunc("/cart/{id}", h.UpdateCartQty).Methods("PATCH")
	r.HandleFunc("/carttrans/{id}", h.UpdateCartTrans).Methods("PATCH")
	r.HandleFunc("/cart/{id}", middleware.Auth(h.DeleteCart)).Methods("DELETE")
	r.HandleFunc("/cartproduct", middleware.Auth(h.CartProduct)).Methods("GET")
	r.HandleFunc("/cartproduct/{id}", middleware.Auth(h.CartProductId)).Methods("GET")
}
