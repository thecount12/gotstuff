package render


// get bootstrap.com
import (
	"fmt"
	"html/template"
	"net/http"
	"log"
)

// render function
/*
func RenderTemplate(w http.ResponseWriter, tmpl string) {
	parsedTemplate, _ := template.ParseFiles("./templates/" + tmpl, "./templates/base.layout.tmpl")
	err := parsedTemplate.Execute(w, nil)
	if err != nil {
		fmt.Println("error parsing template:", err)
		return 
	}
}
*/

// template cache
var tc = make(map[string]*template.Template)


func RenderTemplate(w http.ResponseWriter, t string) {
	var tmpl *template.Template
	var err error
	
	//check to see if we have template in cache
	_, inMap := tc[t]
	if !inMap {
		// need to create template
		err = createTemplateCache(t)
		if err != nil {
			log.Println(err)
		}
	} else {
		// we have the template in cache
		log.Println("using cache tempalte")
	}
	tmpl = tc[t]
	err = tmpl.Execute(w, nil)
	if err != nil {
			log.Println(err)
		}	
}

func createTemplateCache(t string) error {
	templates := []string{
		fmt.Sprintf("./templates/%s", t),
		"./templates/base.layout.tmpl",
	}
	// parse the template
	tmpl, err := template.ParseFiles(templates...)
	if err != nil {
		return err
	}
	
	// add tempalte to cache (map)
	tc[t] = tmpl
	return nil
}



