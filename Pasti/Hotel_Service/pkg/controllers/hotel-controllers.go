package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"PA2_Hotel/pkg/models"
	"PA2_Hotel/pkg/utils"

	"github.com/gorilla/mux"
)

func GetHotel(w http.ResponseWriter, r *http.Request) {
	hotels := models.GetAllHotel()
	res, _ := json.Marshal(hotels)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetHotelById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	HotelId := vars["HotelId"]
	if HotelId == "" {
		// Tangani kasus ketika HotelId kosong
		http.Error(w, "HotelId tidak valid", http.StatusBadRequest)
		return
	}

	ID, err := strconv.ParseInt(HotelId, 10, 64) 	 
	if err != nil {
		fmt.Println("error while parsing:", err)
		http.Error(w, "HotelId tidak valid", http.StatusBadRequest)
		return
	}

	hotelDetails, _ := models.GetHotelbyId(ID)
	res, _ := json.Marshal(hotelDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateHotel(w http.ResponseWriter, r *http.Request) {
	hotel := &models.Hotel{}
	utils.ParseBody(r, hotel)
	createdHotel := hotel.CreateHotel()
	res, _ := json.Marshal(createdHotel)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteHotel(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	HotelId := vars["HotelId"]
	ID, err := strconv.ParseInt(HotelId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing:", err)
	}
	deletedHotel := models.DeleteHotel(ID)
	res, _ := json.Marshal(deletedHotel)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateHotel(w http.ResponseWriter, r *http.Request) {
	updateHotel := &models.Hotel{}
	utils.ParseBody(r, updateHotel)
	vars := mux.Vars(r)
	HotelId := vars["HotelId"]
	ID, err := strconv.ParseInt(HotelId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing:", err)
	}
	hotelDetails, db := models.GetHotelbyId(ID)
	if updateHotel.Nama_hotel != "" {
		hotelDetails.Nama_hotel = updateHotel.Nama_hotel
	}
	if updateHotel.Lokasi != "" {
		hotelDetails.Lokasi = updateHotel.Lokasi
	}
	if updateHotel.Description != "" {
		hotelDetails.Description = updateHotel.Description
	}
	if updateHotel.Harga != 0 {
		hotelDetails.Harga = updateHotel.Harga
	}
	db.Save(&hotelDetails)
	res, _ := json.Marshal(hotelDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
