package main

// go mod init books-list
// go get -u github.com/gorilla/mux
// go get github.com/subosito/gotenv
// pq   go get github.com/lib/pq

/* make one function then copy and paste
 * func getBooks(w http.ResponseWriter, r *http.Request) {
	log.Println("Get all books is called")
}

after basic functions go build run
visit: http://localhost:8000/books
* http://localhost:8000/books/1
2022/08/19 17:29:33 Get all books is called
etc...

i, err := strconv.Atoi("-42")
s := strconv.Itoa(-42)


curl http://localhost:8000/books
curl -d '{json}' -H 'Content-Type: application/json' http://localhost:8000/books
curl -d '{"id": 6, "title": "C++ is old", "author": "Mr. C++", "Year": "2014"}' -H "Content-Type: application/json" -X POST http://localhost:8000/books/6
// let's update
curl -d '{"id":3,"title":"Bazinga","author":"Mr. Sheldon","year":"2012"}' -H "Content-Type: application/json" -X PUT http://localhost:8000/books
curl -X DELETE http://localhost:8000/books/3

*/



import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"reflect"
	"strconv"
)

type Book struct {
	ID 		int		`json:"id"`
	Title 	string	`json:"title"`
	Author	string	`json:"author"`
	Year	string	`json:"year"`
}

var books []Book

func main() {
	fmt.Println("cool")
	
	// copy and paste from gorilla but add Methods("GET")
	router := mux.NewRouter()
	
	books = append(books, 
	Book{ID: 1, Title: "Golang pointers", Author: "Mr. Golang", Year: "2010"},
	Book{ID: 2, Title: "Go Routines", Author: "Mr. Goroutine", Year: "2011"},
	Book{ID: 3, Title: "Golang routers", Author: "Mr. Router", Year: "2012"},
	Book{ID: 4, Title: "Golang concurrency", Author: "Mr. Currency", Year: "2013"},
	Book{ID: 5, Title: "Golang good parts", Author: "Mr. Good", Year: "2014"})
	
	router.HandleFunc("/books", getBooks).Methods("GET")
    router.HandleFunc("/books/{id}", getBook).Methods("GET")
    router.HandleFunc("/books/{id}", addBook).Methods("POST")
    router.HandleFunc("/books", updateBook).Methods("PUT")
    router.HandleFunc("/books/{id}", removeBook).Methods("DELETE")
    
	// start the server
	// http.ListenAndServe(":8000", router)
	log.Fatal(http.ListenAndServe(":8000", router))
}

func getBooks(w http.ResponseWriter, r *http.Request) {
	//log.Println("Get all books is called")
	json.NewEncoder(w).Encode(books)
}

func getBook(w http.ResponseWriter, r *http.Request) {
	// log.Println("Get book is called")
	params := mux.Vars(r)
	log.Println(params) // 2022/08/19 18:58:52 map[id:3]
	log.Println(reflect.TypeOf(params["id"]))
	i, _ := strconv.Atoi(params["id"])
	log.Println(reflect.TypeOf(i))
	
	/*
	for _, book := range books {
		if book.ID == params["id"] {
			json.NewEncoder(w).Encode(&book)
		}
	}
	*/
	for _, book := range books {
		if book.ID == i {
			json.NewEncoder(w).Encode(&book)
		}
	}
	
}

func addBook(w http.ResponseWriter, r *http.Request) {
	// log.Println("Add book is called")
	var book Book
	_ = json.NewDecoder(r.Body).Decode(&book)
	books = append(books, book)
	json.NewEncoder(w).Encode(books)
	log.Println(book)
	
}

func updateBook(w http.ResponseWriter, r *http.Request) {
	// log.Println("Update book is called")
	var book Book
	json.NewDecoder(r.Body).Decode(&book)
	
	for i, item := range books {
		if item.ID == book.ID {
			books[i] = book
		}
	}
	json.NewEncoder(w).Encode(books)
}

func removeBook(w http.ResponseWriter, r *http.Request) {
	// log.Println("Remove book is called")
	params := mux.Vars(r)
	
	id, _ := strconv.Atoi(params["id"])
	for i, item := range books {
		if item.ID == id {
			books = append(books[:i], books[i+1:]...)
		}
	}
	json.NewEncoder(w).Encode(books)
	
}






