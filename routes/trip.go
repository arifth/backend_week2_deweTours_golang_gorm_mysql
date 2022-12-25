package routes

import (
	"gorm-imp/handlers"
	"gorm-imp/pkg/middleware"
	"gorm-imp/pkg/mysql"
	"gorm-imp/repositories"

	"github.com/gorilla/mux"
)

func TripRoutes(r *mux.Router) {
	TripRepository := repositories.RepositoryTrip(mysql.DB)
	h := handlers.HandlerTrip(TripRepository)

	// NOTES: undefined routes will cause error method not allowed in client like postman
	r.HandleFunc("/trips", h.FindTrips).Methods("GET")
	r.HandleFunc("/trip/{id}", h.FindTrip).Methods("GET")
	r.HandleFunc("/trip", middleware.Auth(middleware.UploadFile(h.CreateTrip))).Methods("POST")
	r.HandleFunc("/trip/{id}", middleware.Auth(middleware.UploadFile(h.UpdateTrip))).Methods("PATCH")
	// BUGS: error merah
	r.HandleFunc("/trip/{id}", middleware.Auth(h.CreateTrip)).Methods("DELETE")
}
