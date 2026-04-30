package main
import{
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm/dialects/mysql"
	"github.com/UtkarshSahu9906/go-bookstore/pkg/routes"
	
}

func main(){
	r:=mux.NewRouter()
	routes.ResisterBookStoreRoutes(r)
	http.Handle("/",r)
	log.Fatal(http.ListenAndServe("localhost:9010",r));
}