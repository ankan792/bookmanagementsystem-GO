package main

import (
	"log"
	"net/http"

	"github.com/ankan792/bookmanagementsystem-GO/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisteredRoutes(r)
	http.Handle("/", r)
	log.Fatalln(http.ListenAndServe(":8080", r))
}
