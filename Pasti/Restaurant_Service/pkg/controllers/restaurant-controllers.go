package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"PA2_TourService/pkg/models"
	"PA2_TourService/pkg/utils"

	"github.com/gorilla/mux"
)

var NewRestaurant models.Restaurant

func GetRestaurant(w http.ResponseWriter, r *http.Request) {
	newRestaurant := models.GetAllRestaurant()
	res, _ := json.Marshal(newRestaurant)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func GetRestaurantById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantId := vars["restaurantId"]
	ID, err := strconv.ParseInt(restaurantId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	restaurantDetails, _ := models.GetRestaurantbyId(ID)
	res, _ := json.Marshal(restaurantDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func CreateRestaurant(w http.ResponseWriter, r *http.Request) {
	CreateRestaurant := &models.Restaurant{}
	utils.ParseBody(r, CreateRestaurant)
	b := CreateRestaurant.CreateRestaurant()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func DeleteRestaurant(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	restaurantId := vars["restaurantId"]
	ID, err := strconv.ParseInt(restaurantId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	restaurant := models.DeleteRestaurant(ID)
	res, _ := json.Marshal(restaurant)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
func UpdateRestaurant(w http.ResponseWriter, r *http.Request) {
	var updateRestaurant = &models.Restaurant{}
	utils.ParseBody(r, updateRestaurant)
	vars := mux.Vars(r)
	restaurantId := vars["restaurantId"]
	ID, err := strconv.ParseInt(restaurantId, 0, 0)
	if err != nil {
		fmt.Println("error while parsing")
	}
	restaurantDetails, db := models.GetRestaurantbyId(ID)
	if updateRestaurant.Nama_restaurant != "" {
		restaurantDetails.Nama_restaurant = updateRestaurant.Nama_restaurant
	}
	if updateRestaurant.Lokasi != "" {
		restaurantDetails.Lokasi = updateRestaurant.Lokasi
	}
	if updateRestaurant.Description != "" {
		restaurantDetails.Description = updateRestaurant.Description
	}
	if updateRestaurant.Harga != 0 {
		restaurantDetails.Harga = updateRestaurant.Harga
	}
	db.Save(&restaurantDetails)
	res, _ := json.Marshal(restaurantDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}
