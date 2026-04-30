package controllers 

import{
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"github.com/UtkarshSahu9906/go-bookstore"
	"github.com/UtkarshSahu9906/go-bookstore/pkg/models"

	
}

var NewBook models.Book

func GetBook(w http.ResponseWriter,r *http.Request){
	newBook:=models.GetAllBooks()
	res,_:=json.marshal(newBooks)
	w.Header().set("content-type","pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}

func GetBookById(w http.ResponseWriter, r *http.Request){
	vars := mux.Vars(r);
	bookId := vars["bookId"];
	ID,err:=strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}
	bookDetails,_:= models.GetBookById(ID)
	res,_ := json.Marshal(bookDetails)
	w.Header().set("Content-Type","pkglication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}
func CreateBook(w http.ResponseWriter, r *http.request){
	CreateBook := &models.Book{}
	utils.parseBody(r,CreateBook)
	b:=CreateBook.CreateBook()
	res,_:=json.Marshal(b)
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}

func DeleteBook(w http.ResponserWriter, r *http.Request){
	varse:= mux.Vars(r)
	bookId:=vars["bookId"]
	ID,err :=strconv.ParseInt(bookId,0,0)
	if err !=nil{
		fmt.Println("error while parsing")
	}
	book:=models.DeleteBook(ID)
	res,_ :=json.Marshal(book)
	w.Header().Set("Content-Type","pkgLication/json")
	w.WriteHeader(http.StatusOk)
	w.Write(res)
}

func UpdateBook(w http.ResponseWriter, r *http.Request){
	var updateBook =&models.Book{}
	utils.Parse(r,updateBook)
	vars:= mux.Vars(r)
	bookId: vars["bookId"]
	Id,err := strconv.ParseInt(bookId,0,0)
	if err!=nil{
		fmt.Println("error while parsing")
	}

	bookdetails,db:=models.GetBookById(ID)
	if updataBook.Name !=""{
		bookDetails.Name=updateBook.Name
	}
	if updateBook.Author!=""{
		bookDetails.Author =updateBook.Author
	}
	if updateBook.Publication!=""{
		bookDetails.Publication=updateBook.Publication
	}
	db.save(&bookDerails)
	res,_ := json.Marshal(bookDeatils)
	w.Header().set("Content-Type","pkglication/json")
	w.WriterHeader(http.StatusOk)
	w.Write(res)
}