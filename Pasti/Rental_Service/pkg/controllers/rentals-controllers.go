package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"Rental_Service/pkg/models"
	"Rental_Service/pkg/utils"

	"github.com/gorilla/mux"
)

var NewRentals models.Rentals 

func GetRentals(w http.ResponseWriter, r *http.Request) {
	newRentals := models.GetAllRentals()
	res, _ := json.Marshal(newRentals)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetRentalsById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rentalsId := vars["rentalsId"]
	ID, err := strconv.ParseInt(rentalsId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	rentalsDetails, _ := models.GetRentalsbyId(ID)
	res, _ := json.Marshal(rentalsDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateRentals(w http.ResponseWriter, r *http.Request) {
	CreateRentals := &models.Rentals{}
	utils.ParseBody(r, CreateRentals)
	b := CreateRentals.CreateRentals()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteRentals(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	rentalsId := vars["rentalsId"]
	ID, err := strconv.ParseInt(rentalsId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	rentals := models.DeleteRentals(ID)
	res, _ := json.Marshal(rentals)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateRentals(w http.ResponseWriter, r *http.Request) {
	var updateRentals = &models.Rentals{}
	utils.ParseBody(r, updateRentals)
	vars := mux.Vars(r)
	rentalsId := vars["rentalsId"]
	ID, err := strconv.ParseInt(rentalsId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	rentalsDetails, db := models.GetRentalsbyId(ID)
	if updateRentals.Pickinglocation != "" {
		rentalsDetails.Pickinglocation = updateRentals.Pickinglocation
	}
	if updateRentals.Pickingoff != "" {
		rentalsDetails.Pickingoff = updateRentals.Pickingoff
	}
	if updateRentals.Pickingupdate != "" {
		rentalsDetails.Pickingupdate = updateRentals.Pickingupdate
	}
	if updateRentals.Pickinguphour != ""{
		rentalsDetails.Pickinguphour = updateRentals.Pickinguphour
	}
	if updateRentals.Pickingoffdate != "" {
		rentalsDetails.Pickingoffdate = updateRentals.Pickingoffdate
	}
	if updateRentals.Pickingoffhour !="" {
		rentalsDetails.Pickingoffhour = updateRentals.Pickingoffhour
	}
	
	db.Save(&rentalsDetails)
	res, _ := json.Marshal(rentalsDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
