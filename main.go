package main

import (
	"fmt"
	"gorm-imp/database"
	"gorm-imp/pkg/mysql"
	"gorm-imp/routes"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	// initiate DB
	mysql.DatabaseInit()

	// run migration
	database.RunMigration()

	r := mux.NewRouter()

	routes.RouteInit(r.PathPrefix("/api/v1").Subrouter())

	fmt.Println("server running on localhost:5000")
	http.ListenAndServe("localhost:5000", r)
}
