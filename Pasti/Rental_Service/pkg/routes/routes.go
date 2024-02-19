package routes

import (
	"Rental_Service/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRental_ServiceRoutes = func(router *mux.Router) {
	router.HandleFunc("/rentals", controllers.CreateRentals).Methods("POST")
	router.HandleFunc("/rentals", controllers.GetRentals).Methods("GET")
	router.HandleFunc("/rentals/{rentalsId}", controllers.GetRentalsById).Methods("GET")
	router.HandleFunc("/rentals/{rentalsId}", controllers.UpdateRentals).Methods("PUT")
	router.HandleFunc("/rentals/{rentalsId}", controllers.DeleteRentals).Methods("DELETE")
}
