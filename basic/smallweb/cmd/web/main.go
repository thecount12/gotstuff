package main
// could use 'go mod init myapp'
// go run *.go
// move stuff around now
// create cmd/web/main.go
// create pkg/handler/handler.go
// create pkg/render/render.go

import (
	"fmt"
	"net/http"
	//"html/template"
	//"errors"
	"myapp/pkg/handlers"
)

const portNumber = ":8080"


// main just main
func main() {
	http.HandleFunc("/", handlers.Home)
	http.HandleFunc("/about", handlers.About)
	fmt.Println(fmt.Sprintf("Starting app on port %s", portNumber))
	http.ListenAndServe(portNumber, nil)
}
