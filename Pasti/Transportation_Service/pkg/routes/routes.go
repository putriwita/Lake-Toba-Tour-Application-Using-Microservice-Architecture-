package routes

import (
	"Transportation_Service/pkg/controllers"

	"github.com/gorilla/mux"
)

var RegisterTransportationsRoutes = func(router *mux.Router) {
	router.HandleFunc("/transportations", controllers.CreateTransportations).Methods("POST")
	router.HandleFunc("/transportations", controllers.GetTransportations).Methods("GET")
	router.HandleFunc("/transportations/{TransportationsId}", controllers.GetTransportationsById).Methods("GET")
	router.HandleFunc("/transportations/{TransportationsId}", controllers.UpdateTransportations).Methods("PUT")
	router.HandleFunc("/transportations/{TransportationsId}", controllers.DeleteTransportations).Methods("DELETE")
}
