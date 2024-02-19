package models

import (
	"Booking_Service/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Booking struct {
	gorm.Model
	Nama	string `json:"nama"`
	Email  	string `json:"email"`
	Telepon	string `json:"telepon"`
	Status	string `json:"status"`
	Tanggal	string `json:"tanggal"`
}


func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Booking{})
}
func (b *Booking) CreateBooking() *Booking {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllBooking() []Booking {
	var Booking []Booking
	db.Find(&Booking)
	return Booking
}
func GetBookingbyId(id int64) (*Booking, *gorm.DB) {
	var GetBooking Booking
	db := db.Where("id=?", id).Find(&GetBooking)
	return &GetBooking, db
}
func DeleteBooking(id int64) Booking {
	var booking Booking
	db.Where("id=?", id).Delete(booking)
	return booking
}
