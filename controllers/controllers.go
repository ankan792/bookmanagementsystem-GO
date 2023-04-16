package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/ankan792/bookmanagementsystem-GO/models"
	"github.com/gorilla/mux"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book *models.Book
	err := json.NewDecoder(r.Body).Decode(&book)
	if err != nil {
		log.Fatalln(err)
	}
	createdBook := models.Create(book)
	res, _ := json.Marshal(createdBook)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAll()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.Atoi(params["ID"])
	if err != nil {
		log.Println("error converting string to int")
	}
	book, err := models.GetByID(bookID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	//convert the json body request into golang struct format
	var updatedBook *models.Book
	err := json.NewDecoder(r.Body).Decode(&updatedBook)
	if err != nil {
		log.Fatalln(err)
	}

	//get the book using the book ID
	params := mux.Vars(r)
	bookID, err := strconv.Atoi(params["ID"])
	if err != nil {
		log.Println("error converting string to int")
	}
	book, err := models.GetByID(bookID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	//update the corresponding book
	if updatedBook.Name != "" {
		book.Name = updatedBook.Name
	}
	if updatedBook.Author != "" {
		book.Author = updatedBook.Author
	}
	if updatedBook.Publication != "" {
		book.Publication = updatedBook.Publication
	}

	//respond back the updated book to the user
	res, _ := json.Marshal(book)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	bookID, err := strconv.Atoi(params["ID"])
	if err != nil {
		log.Println("error converting string to int")
	}
	err = models.DeleteByID(bookID)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	GetBooks(w, r)
}
