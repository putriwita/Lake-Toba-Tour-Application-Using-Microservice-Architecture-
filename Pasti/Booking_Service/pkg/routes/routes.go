package routes

import (
	"Booking_Service/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterHotelRoutes = func(router *mux.Router) {
	router.HandleFunc("/booking", controllers.CreateBooking).Methods("POST")
	router.HandleFunc("/booking", controllers.GetBooking).Methods("GET")
	router.HandleFunc("/booking/{BookingId}", controllers.GetBookingById).Methods("GET")
	router.HandleFunc("/booking/{BookingId}", controllers.UpdateBooking).Methods("PUT")
	router.HandleFunc("/booking/{BookingId}", controllers.DeleteBooking).Methods("DELETE")
}
