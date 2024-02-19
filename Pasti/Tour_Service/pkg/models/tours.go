package models

import (
	"TourService/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Tours struct {
	gorm.Model
	Image       string `json:"image"`
	Nama        string `json:"nama"`
	Detail      string `json:"detail"`
	Harga       uint   `json:"harga"`
	Jumlahorang uint   `json:"jumlahorang"`
	Diskon      uint   `json:"diskon"`
	Hargatotal  uint   `json:"hargatotal"`
	Jumlah_hari uint   `json:"jumlah_hari"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Tours{})
}
func (b *Tours) CreateTours() *Tours {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllTours() []Tours {
	var Tours []Tours
	db.Find(&Tours)
	return Tours
}
func GetToursbyId(id int64) (*Tours, *gorm.DB) {
	var getTours Tours
	db := db.Where("id=?", id).Find(&getTours)
	return &getTours, db
}
func DeleteTours(id int64) Tours {
	var tours Tours
	db.Where("id=?", id).Delete(tours)
	return tours
}
