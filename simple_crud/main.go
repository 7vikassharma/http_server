package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Book struct {
	BookId string `json:"id,omitempty"`
	Title  string `json:"title"`
	Author string `json:"author"`
}

var books []Book

func main() {
	fmt.Println("vikas")
	router := mux.NewRouter()

	router.HandleFunc("/books", GetBooks).Methods("GET")
	router.HandleFunc("/books/{id}", GetBook).Methods("GET")
	router.HandleFunc("/books", CreateBook).Methods("POST")
	router.HandleFunc("/books/{id}", UpdateBook).Methods("PUT")
	router.HandleFunc("/books/{id}", DeleteBook).Methods("DELETE")
	fmt.Println("started")
	log.Fatal(http.ListenAndServe(":8000", router))

}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	if books == nil {
		http.Error(w, "Data does not exist", http.StatusNotFound)
		return

	}
	json.NewEncoder(w).Encode(books)

}

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	for _, book := range books {
		if book.BookId == params["id"] {
			json.NewEncoder(w).Encode(book)
		}
	}

}

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	for _, item := range books {
		if item.BookId == book.BookId {
			http.Error(w, "Data already exist", http.StatusAlreadyReported)
			return

		}
	}
	books = append(books, book)
	json.NewEncoder(w).Encode(book)

}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range books {
		if item.BookId == params["id"] {
			books = append(books[:index], books[index+1:]...)
			var book Book
			_ = json.NewDecoder(r.Body).Decode(&book)
			book.BookId = params["id"]
			books = append(books, book)
			json.NewEncoder(w).Encode(book)
			return
		}
	}
	json.NewEncoder(w).Encode(books)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	for index, item := range books {
		if item.BookId == params["id"] {
			books = append(books[:index], books[index+1:]...)
			break
		}
	}
	json.NewEncoder(w).Encode(books)
}
