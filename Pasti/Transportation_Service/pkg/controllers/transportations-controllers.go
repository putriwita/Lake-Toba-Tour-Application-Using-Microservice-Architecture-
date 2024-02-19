package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"Transportation_Service/pkg/models"
	"Transportation_Service/pkg/utils"

	"github.com/gorilla/mux"
)

func GetTransportations(w http.ResponseWriter, r *http.Request) {
	transportations := models.GetAllTransportations()
	res, _ := json.Marshal(transportations)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetTransportationsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	TransportationsId := vars["TransportationsId"]
	if TransportationsId == "" {
		// Tangani kasus ketika TransportationsId kosong
		http.Error(w, "TransportationsId tidak valid", http.StatusBadRequest)
		return
	}

	ID, err := strconv.ParseInt(TransportationsId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing:", err)
		http.Error(w, "TransportationsId tidak valid", http.StatusBadRequest)
		return
	}

	transportationsDetails, _ := models.GetTransportationsbyId(ID)
	res, _ := json.Marshal(transportationsDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateTransportations(w http.ResponseWriter, r *http.Request) {
	transportations := &models.Transportations{}
	utils.ParseBody(r, transportations)
	createdTransportations := transportations.CreateTransportations()
	res, _ := json.Marshal(createdTransportations)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteTransportations(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	TransportationsId := vars["TransportationsId"]
	ID, err := strconv.ParseInt(TransportationsId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing:", err)
	}
	deletedTransportations := models.DeleteTransportations(ID)
	res, _ := json.Marshal(deletedTransportations)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateTransportations(w http.ResponseWriter, r *http.Request) {
	updateTransportations := &models.Transportations{}
	utils.ParseBody(r, updateTransportations)
	vars := mux.Vars(r)
	TransportationsId := vars["TransportationsId"]
	ID, err := strconv.ParseInt(TransportationsId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing:", err)
	}
	transportationsDetails, db := models.GetTransportationsbyId(ID)
	if updateTransportations.Name != "" {
		transportationsDetails.Name = updateTransportations.Name
	}
	if updateTransportations.Image != "" {
		transportationsDetails.Image = updateTransportations.Image
	}
	if updateTransportations.Type != "" {
		transportationsDetails.Type = updateTransportations.Type
	}
	if updateTransportations.Price != 0 {
		transportationsDetails.Price = updateTransportations.Price
	}
	if updateTransportations.Route != "" {
		transportationsDetails.Route = updateTransportations.Route
	}
	if updateTransportations.Duration != 0 {
		transportationsDetails.Duration = updateTransportations.Duration
	}
	if updateTransportations.Location != "" {
		transportationsDetails.Location = updateTransportations.Location
	}
	db.Save(&transportationsDetails)
	res, _ := json.Marshal(transportationsDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
