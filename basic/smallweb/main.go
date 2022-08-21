package main
// could use 'go mod init myapp'
import (
	"fmt"
	"net/http"
)

const portNumber = ":8080"

// Home is the home page handler
func Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "this is the home page")
}

// About is the about page handler
func About(w http.ResponseWriter, r *http.Request) {
	sum := AddValues(2, 2)
	// fmt.Fprintf(w, "this is the about page")
	_, _ = fmt.Fprintf(w, fmt.Sprintf("This is the about page and 2+2 is %d", sum))
}

// AddValues just add 
func AddValues(x, y int) int {
	return x + y
}

// main just main
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
