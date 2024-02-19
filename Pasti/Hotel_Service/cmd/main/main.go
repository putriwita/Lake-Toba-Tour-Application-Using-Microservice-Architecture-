package main

import (
	"PA2_Hotel/pkg/routes"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterHotelRoutes(r)
	http.Handle("/", r)
	fmt.Print("Starting Server localhost:9090")
	log.Fatal(http.ListenAndServe("localhost:9090", r))
}
