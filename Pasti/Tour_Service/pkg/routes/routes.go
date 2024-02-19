package routes

import (
	"TourService/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterTourServiceRoutes = func(router *mux.Router) {
	router.HandleFunc("/tours", controllers.CreateTours).Methods("POST")
	router.HandleFunc("/tours", controllers.GetTours).Methods("GET")
	router.HandleFunc("/tours/{toursId}", controllers.GetToursById).Methods("GET")
	router.HandleFunc("/tours/{toursId}", controllers.UpdateTours).Methods("PUT")
	router.HandleFunc("/tours/{toursId}", controllers.DeleteTours).Methods("DELETE")
}
