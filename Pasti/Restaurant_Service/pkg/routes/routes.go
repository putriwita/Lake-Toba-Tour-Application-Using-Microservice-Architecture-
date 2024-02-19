package routes

import (
	"PA2_TourService/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterRestaurantRoutes = func(router *mux.Router) {
	router.HandleFunc("/restaurant",controllers.CreateRestaurant).Methods("POST")
	router.HandleFunc("/restaurant", controllers.GetRestaurant).Methods("GET")
	router.HandleFunc("/restaurant/{restaurantId}",controllers.GetRestaurantById).Methods("GET")
	router.HandleFunc("/restaurant/{restaurantId}",controllers.UpdateRestaurant).Methods("PUT")
	router.HandleFunc("/restaurant/{restaurantId}",controllers.DeleteRestaurant).Methods("DELETE")
}
