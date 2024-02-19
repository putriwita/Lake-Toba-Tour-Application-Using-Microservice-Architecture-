package models

import (
	"Rental_Service/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Rentals struct {
	gorm.Model
	Pickinglocation string `json:"pickinglocation"`
	Pickingoff      string `json:"pickingoff"`
	Pickingupdate   string `json:"pickingupdate"`
	Pickinguphour   string   `json:"pickinguphour"`
	Pickingoffdate  string   `json:"pickingoffdate"`
	Pickingoffhour  string   `json:"pickingoffhour"`
} 

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Rentals{})
}
func (b *Rentals) CreateRentals() *Rentals {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllRentals() []Rentals {
	var Rentals []Rentals
	db.Find(&Rentals)
	return Rentals
}
func GetRentalsbyId(id int64) (*Rentals, *gorm.DB) {
	var getRentals Rentals
	db := db.Where("id=?", id).Find(&getRentals)
	return &getRentals, db
}
func DeleteRentals(id int64) Rentals {
	var rentals Rentals
	db.Where("id=?", id).Delete(rentals)
	return rentals
}
