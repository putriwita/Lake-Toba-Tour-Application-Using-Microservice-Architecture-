package routes

import (
	"PA2_Hotel/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterHotelRoutes = func(router *mux.Router) {
	router.HandleFunc("/hotel",controllers.CreateHotel).Methods("POST")
	router.HandleFunc("/hotel", controllers.GetHotel).Methods("GET")
	router.HandleFunc("/hotel/{HotelId}",controllers.GetHotelById).Methods("GET")
	router.HandleFunc("/hotel/{HotelId}",controllers.UpdateHotel).Methods("PUT")
	router.HandleFunc("/hotel/{HotelId}",controllers.DeleteHotel).Methods("DELETE")
}
