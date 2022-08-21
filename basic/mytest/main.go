package main


// 2. new files main_test.go
import (
	"errors"
	"log"
)
	
// 1. build main
func main() {
	
	result, err := divide(100.0, 0)
	if err != nil {
		log.Println(err)
		return
	}
	log.Println("result of division is", result)
}
// 1. build main
func divide(x, y float32) (float32, error) {
	var result float32
	if y == 0 {
		return result, errors.New("cannot divid by 0")
	}
	result = x/y
	return result, nil
}
