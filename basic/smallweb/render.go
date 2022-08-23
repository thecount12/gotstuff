package main


// get bootstrap.com
import (
	"fmt"
	"html/template"
	"net/http"
)

// render function
func renderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl)
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return 
	}
}
