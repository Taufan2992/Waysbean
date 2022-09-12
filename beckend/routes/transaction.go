package routes

import (
	"waysbean/handlers"
	"waysbean/pkg/middleware"
	"waysbean/pkg/mysql"
	"waysbean/repositories"

	"github.com/gorilla/mux"
)

// Create UserRoutes function here ...
func TransactionRoutes(r *mux.Router) {
	transactionRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(transactionRepository)

	r.HandleFunc("/transactions", h.FindTransactions).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.GetTransaction).Methods("GET")
	r.HandleFunc("/transactionId", middleware.Auth(h.FindTransactionId)).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(h.CreateTransaction)).Methods("POST")
	// r.HandleFunc("/transaction/{id}", middleware.Auth(h.UpdateTransaction)).Methods("PATCH")
	r.HandleFunc("/transaction/{id}", middleware.Auth(h.DeleteTransaction)).Methods("DELETE")
	r.HandleFunc("/midtrans/{id}", middleware.Auth(h.Midtrans)).Methods("GET")
}
