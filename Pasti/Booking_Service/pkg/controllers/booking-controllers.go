package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"Booking_Service/pkg/models"
	"Booking_Service/pkg/utils"

	"github.com/gorilla/mux"
)

func GetBooking(w http.ResponseWriter, r *http.Request) {
	bookings := models.GetAllBooking()
	res, _ := json.Marshal(bookings)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookingById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookingId := vars["BookingId"]
	if BookingId == "" {
		// Tangani kasus ketika BookingId kosong
		http.Error(w, "BookingId tidak valid", http.StatusBadRequest)
		return
	}

	ID, err := strconv.ParseInt(BookingId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing:", err)
		http.Error(w, "BookingId tidak valid", http.StatusBadRequest)
		return
	}

	bookingDetails, _ := models.GetBookingbyId(ID)
	res, _ := json.Marshal(bookingDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func CreateBooking(w http.ResponseWriter, r *http.Request) {
	booking := &models.Booking{}
	utils.ParseBody(r, booking)
	createdBooking := booking.CreateBooking()
	res, _ := json.Marshal(createdBooking)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBooking(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	BookingId := vars["BookingId"]
	ID, err := strconv.ParseInt(BookingId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing:", err)
	}
	deletedBooking := models.DeleteBooking(ID)
	res, _ := json.Marshal(deletedBooking)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func UpdateBooking(w http.ResponseWriter, r *http.Request) {
	updateBooking := &models.Booking{}
	utils.ParseBody(r, updateBooking)
	vars := mux.Vars(r)
	BookingId := vars["BookingId"]
	ID, err := strconv.ParseInt(BookingId, 10, 64)
	if err != nil {
		fmt.Println("error while parsing:", err)
	}
	bookingDetails, db := models.GetBookingbyId(ID)
	if updateBooking.Nama != "" {
		bookingDetails.Nama = updateBooking.Nama
	}
	if updateBooking.Email != "" {
		bookingDetails.Email = updateBooking.Email
	}
	if updateBooking.Telepon != "" {
		bookingDetails.Telepon = updateBooking.Telepon
	}
	if updateBooking.Status != "" {
		bookingDetails.Status = updateBooking.Status
	}
	if updateBooking.Tanggal != "" {
		bookingDetails.Tanggal = updateBooking.Tanggal
	}

	db.Save(&bookingDetails)
	res, _ := json.Marshal(bookingDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
