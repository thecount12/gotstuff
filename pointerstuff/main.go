package main

import "log"

func main() {
	var myString string
	myString = "Green"
	
	log.Println("myString is set to", myString)
	ChangeUsingPointer(&myString)
	log.Println("after func myString is now set to", myString)
}

func ChangeUsingPointer(s *string) {
	log.Println("s is set to", s)
	newValue := "Red"
	*s = newValue  // you just don't need to return anything
}
