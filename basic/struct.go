package main

import (
	"time"
	"fmt"
	"log"
)

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	Birthday time.Time	
}

type myStruct struct {
	FirstName string
}

// receiver
func (m *myStruct) printFirstname() string {
	return m.FirstName
	
}

func main() {
	fmt.Println("asdf")
	user := User{
		FirstName: "Will", 
		LastName: "Gunn",
		}
	log.Println(user.FirstName, user.LastName, user.Birthday)
	whatever()
	Whatever()
	phone := User{PhoneNumber: "555-1212"} // different set of records
	log.Println(phone.PhoneNumber)
	
	
	// ther struct
	var myVar myStruct
	myVar.FirstName = "John"
	myVar2 := myStruct{FirstName: "Mary", }
	log.Println("myVar is set to ", myVar.FirstName)
	log.Println("myVar2 is set to", myVar2.FirstName)
	// receiver functions hand be handy
	log.Println("myVar is set to ", myVar.printFirstname())
	log.Println("myVar2 is set to", myVar2.printFirstname())
}

func whatever() { // available inside
	
}

func Whatever() {  // available outside
	
}

// camelCase is important Outside, inside

