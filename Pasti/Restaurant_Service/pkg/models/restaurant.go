package models

import (
	"PA2_TourService/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Restaurant struct {
	gorm.Model
	Nama_restaurant string `json:"nama_restaurant"`
	Lokasi          string `json:"lokasi"`
	Description     string `json:"description"`
	Harga           uint   `json:"harga"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Restaurant{})
}
func (b *Restaurant) CreateRestaurant() *Restaurant {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllRestaurant() []Restaurant {
	var Restaurant []Restaurant
	db.Find(&Restaurant)
	return Restaurant
}
func GetRestaurantbyId(id int64) (*Restaurant, *gorm.DB) {
	var getRestaurant Restaurant
	db := db.Where("id=?", id).Find(&getRestaurant)
	return &getRestaurant, db
}
func DeleteRestaurant(id int64) Restaurant {
	var restaurant Restaurant
	db.Where("id=?", id).Delete(restaurant)
	return restaurant
}
