package routes
import (
	"github.com/gofiber/fiber/v2"
	"github.com/utkarshsahu9906/go-bookstore/controllers"
)



var ResisterBookStoreRoutes =func(router *mux.Router){
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetBooks).Methods("GET")
	router.HandleFunc("/book/{id}",controllers.getBookById).Methods("GET")
	router.HandleFunc("/book/{id}",controllers.UpdateBook).Methods("PUT")
	router.HandleFunc("/book/{id}",controllers.DeleteBook).Methods("DELETE")
}









