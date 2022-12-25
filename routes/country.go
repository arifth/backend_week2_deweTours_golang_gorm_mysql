package routes

import (
	"gorm-imp/handlers"
	"gorm-imp/pkg/middleware"
	"gorm-imp/pkg/mysql"
	"gorm-imp/repositories"

	"github.com/gorilla/mux"
)

func CountryRoutes(r *mux.Router) {
	CountryRepository := repositories.RepositoryCountry(mysql.DB)
	h := handlers.HandlerCountry(CountryRepository)

	// NOTES: undefined routes will cause error method not allowed in client like postman

	r.HandleFunc("/countries", h.FindCountries).Methods("GET")
	r.HandleFunc("/country/{id}", h.FindCountry).Methods("GET")
	r.HandleFunc("/country", middleware.Auth(h.CreateCountry)).Methods("POST")
	r.HandleFunc("/country/{id}", middleware.Auth(h.UpdateCountry)).Methods("PATCH")
	r.HandleFunc("/country/{id}", middleware.Auth(h.DeleteCountry)).Methods("DELETE")
}
