package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"TourService/pkg/models"
	"TourService/pkg/utils"

	"github.com/gorilla/mux"
)

var NewTours models.Tours 

func GetTours(w http.ResponseWriter, r *http.Request) {
	newTours := models.GetAllTours()
	res, _ := json.Marshal(newTours)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetToursById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	toursId := vars["toursId"]
	ID, err := strconv.ParseInt(toursId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	toursDetails, _ := models.GetToursbyId(ID)
	res, _ := json.Marshal(toursDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateTours(w http.ResponseWriter, r *http.Request) {
	CreateTours := &models.Tours{}
	utils.ParseBody(r, CreateTours)
	b := CreateTours.CreateTours()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteTours(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	toursId := vars["toursId"]
	ID, err := strconv.ParseInt(toursId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	tours := models.DeleteTours(ID)
	res, _ := json.Marshal(tours)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateTours(w http.ResponseWriter, r *http.Request) {
	var updateTours = &models.Tours{}
	utils.ParseBody(r, updateTours)
	vars := mux.Vars(r)
	toursId := vars["toursId"]
	ID, err := strconv.ParseInt(toursId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	toursDetails, db := models.GetToursbyId(ID)
	if updateTours.Image != "" {
		toursDetails.Image = updateTours.Image
	}
	if updateTours.Nama != "" {
		toursDetails.Nama = updateTours.Nama
	}
	if updateTours.Detail != "" {
		toursDetails.Detail = updateTours.Detail
	}
	if updateTours.Harga != 0 {
		toursDetails.Harga = updateTours.Harga
	}
	if updateTours.Jumlahorang != 0 {
		toursDetails.Jumlahorang = updateTours.Jumlahorang
	}
	if updateTours.Diskon !=0{
		toursDetails.Diskon = updateTours.Diskon
	}
	if updateTours.Hargatotal !=0 {
		toursDetails.Hargatotal = updateTours.Hargatotal
	}
	if updateTours.Jumlah_hari != 0 {
		toursDetails.Jumlah_hari = updateTours.Jumlah_hari
	}
	db.Save(&toursDetails)
	res, _ := json.Marshal(toursDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
