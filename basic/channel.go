package main

// go mod init helpers
import (
	"log"
	"foo/helpers"
	)


const numPool = 10
// 2. mkdir helpers and touch helpers.go
// 4. create function to use helpers random number
func CalculateValue(intChan chan int) {
	randomNumber := helpers.RandomNumber(numPool)
	intChan <- randomNumber
}

func main() {
	
	// 1. place to send info this will be integers
	intChan := make(chan int)
	// lets close it after we execute it not needed in main but make a habbit of it
	defer close(intChan)
	
	go CalculateValue(intChan)
	num := <-intChan  //this is listener
	log.Println(num)
	// 5. before we run this want to do 'go mod inito foo
	// 6. type foo/helpers in import section
	
}
