package routes

import (
	"gorm-imp/handlers"
	"gorm-imp/pkg/mysql"
	"gorm-imp/repositories"

	"github.com/gorilla/mux"
)

func userRoutes(r *mux.Router) {
	UserRepository := repositories.RepositoryUser(mysql.DB)

	h := handlers.HandlerUser(UserRepository)

	r.HandleFunc("/users", h.FindUsers).Methods("GET")
	r.HandleFunc("/user/{id}", h.GetUser).Methods("GET")
	r.HandleFunc("/user", h.CreateUser).Methods("POST")
	r.HandleFunc("/user/{id}", h.UpdateUser).Methods("PATCH")
	r.HandleFunc("/user/{id}", h.DeleteUser).Methods("DELETE")
}
