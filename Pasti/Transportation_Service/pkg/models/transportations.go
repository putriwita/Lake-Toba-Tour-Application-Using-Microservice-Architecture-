package models

import (
	"Transportation_Service/pkg/config"

	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Transportations struct {
	gorm.Model
	Name     string `json:"name"`
	Image    string `json:"image"`
	Type     string `json:"type"`
	Price    uint   `json:"price"`
	Route    string `json:"route"`
	Duration uint   `json:"duration"`
	Location string `json:"location"`
}

func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Transportations{})
}
func (b *Transportations) CreateTransportations() *Transportations {
	db.NewRecord(b)
	db.Create(&b)
	return b
}
func GetAllTransportations() []Transportations {
	var Transportations []Transportations
	db.Find(&Transportations)
	return Transportations
}
func GetTransportationsbyId(id int64) (*Transportations, *gorm.DB) {
	var GetTransportations Transportations
	db := db.Where("id=?", id).Find(&GetTransportations)
	return &GetTransportations, db
}
func DeleteTransportations(id int64) Transportations {
	var transportations Transportations
	db.Where("id=?", id).Delete(transportations)
	return transportations
}
