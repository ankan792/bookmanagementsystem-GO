package routes

import (
	"github.com/ankan792/bookmanagementsystem-GO/controllers"
	"github.com/gorilla/mux"
)

var RegisteredRoutes = func(routes *mux.Router) {
	routes.HandleFunc("/books", controllers.CreateBook).Methods("POST")
	routes.HandleFunc("/books", controllers.GetBooks).Methods("GET")
	routes.HandleFunc("/books/{ID}", controllers.GetBookByID).Methods("GET")
	routes.HandleFunc("/books/{ID}", controllers.UpdateBook).Methods("PUT")
	routes.HandleFunc("/books/{ID}", controllers.DeleteBook).Methods("DELETE")
}
