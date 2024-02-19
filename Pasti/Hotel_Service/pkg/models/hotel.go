package models

import (
	"PA2_Hotel/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Hotel struct {
	gorm.Model
	Nama_hotel  string `json:"nama_hotel"`
	Lokasi      string `json:"lokasi"`
	Description string `json:"description"`
	Harga       uint   `json:"harga"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Hotel{})
}
func (b *Hotel) CreateHotel() *Hotel {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllHotel() []Hotel {
	var Hotel []Hotel
	db.Find(&Hotel)
	return Hotel
}
func GetHotelbyId(id int64) (*Hotel, *gorm.DB) {
	var GetHotel Hotel
	db := db.Where("id=?", id).Find(&GetHotel)
	return &GetHotel, db
}
func DeleteHotel(id int64) Hotel {
	var hotel Hotel
	db.Where("id=?", id).Delete(hotel)
	return hotel
}
