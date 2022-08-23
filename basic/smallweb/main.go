package main
// could use 'go mod init myapp'
// go run *.go
import (
	"fmt"
	"net/http"
	//"html/template"
	//"errors"
)

const portNumber = ":8080"


// main just main
func main() {
	http.HandleFunc("/", Home)
	http.HandleFunc("/about", About)
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
