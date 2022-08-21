package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName string
}

func main() {
	
	// make built in keyword
	myMap := make(map[string]string)
	
	myMap["dog"] = "samson"
	myMap["other-dog"] = "Cassie"
	myMap["dog"] = "fido"
	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])
	
	newMap := make(map[string]int)
	newMap["first"] = 1
	newMap["second"] = 2
	log.Println(newMap["first"])
	
	userMap := make(map[string]User)
	me := User {FirstName: "will", LastName: "Gunn",}
	userMap["me"] = me
	log.Println(userMap["me"].FirstName)
	
	// slices
	var mySlice []string
	mySlice = append(mySlice, "Will")
	mySlice = append(mySlice, "John")
	log.Println(mySlice)
	sort.Strings(mySlice)
	log.Println(mySlice)
	
}
