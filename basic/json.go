package main

import (
	"log"
	"encoding/json"
	"fmt"
)

// 2. create type
type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}


func main() {
	// 1. Create json data
	myJson := `
[
	{
		"first_name": "Clark",
		"last_name": "Kent",
		"hair__color": "black",
		"has_dog": true
	},
	{
		"first_name": "Bruce",
		"last_name": "Wayne",
		"hair__color": "black",
		"has_dog": false
	}
]`	

// 3. we need to unmarshall the data
	var unmarshalled []Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled)
	if err != nil {
		log.Println("Error unmarshalling json", err)
	}
	log.Printf("unmarshalled: %v", unmarshalled)
	
	// 4. write json from struct
	var mySlice []Person
	var m1 Person
	m1.FirstName = "Wally"
	m1.LastName = "West"
	m1.HairColor = "red"
	m1.HasDog = false
	mySlice = append(mySlice, m1)
	
	newJson, err := json.MarshalIndent(mySlice, "", "		")
	if err != nil {
		log.Println("error marshalling", err)
	}
	
	fmt.Println(string(newJson))
}




