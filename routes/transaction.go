package routes

import (
	"gorm-imp/handlers"
	"gorm-imp/pkg/middleware"
	"gorm-imp/pkg/mysql"
	"gorm-imp/repositories"

	"github.com/gorilla/mux"
)

func TransactionRoutes(r *mux.Router) {
	TransRepository := repositories.RepositoryTransaction(mysql.DB)
	h := handlers.HandlerTransaction(TransRepository)

	// NOTES: undefined routes will cause error method not allowed in client like postman

	r.HandleFunc("/orders", h.FindTrans).Methods("GET")
	r.HandleFunc("/transaction/{id}", h.FindTran).Methods("GET")
	r.HandleFunc("/transaction", middleware.Auth(middleware.UploadFile(h.CreateTrans))).Methods("POST")
	r.HandleFunc("/transaction/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTrans))).Methods("PATCH")
	// r.HandleFunc("/trip/{id}", middleware.Auth(h.DeleteTrip)).Methods("DELETE")
}
