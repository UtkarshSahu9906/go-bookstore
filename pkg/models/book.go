package models

import{
	"github.com/jinzhu/gorm"
	"github.com/utkarshsahu9906/go-bookstore/config"
}

var db *gorm.DB

type Book struct{
	gorm.Model
	Name string `gorm:"" json:"name"`
	Author string `json:"author"`
}
func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}