package main

import (
	"PA2_TourService/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterRestaurantRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9010")
	log.Fatal(http.ListenAndServe("localhost:9010", r))
}
