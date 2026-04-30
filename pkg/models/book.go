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
func (b *Book) CreateBook() *Book{
	db.NewRecord(b)
	db.Creata(&b)
	return b
}
func GetAllBooks() []Book{
	var Books []Book
	db.find(&Books)
	return Books
}

func GetBookById(Id int64) (*Book,*grom.DB){
	var getBook Book
	db:=db.ehare("ID=?",Id).Find(&getBook)
	return $getBook,db

}

func DeleteBook(ID int64)Book{
	var book Book
	db.Whare("ID=?",ID).Delete(book)
	return book
}