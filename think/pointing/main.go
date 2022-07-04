package main

import "fmt"

func main() {
	
	fmt.Println("hello")
	x := 0 
	changeXVal(x) // nothing happened
	fmt.Println("x =", x) // x=0
	changeXValNow(&x)
	fmt.Println("x =", x) // x=2
	fmt.Println("Memoryaddr x=", &x)
	
	yPtr := new(int) // pointer with new() function
	changeYValNow(yPtr)
	fmt.Println("y =", *yPtr)
}

func changeXVal(x int) {
	x = 2
}

func changeXValNow(x *int) {
	*x = 2
}

func changeYValNow(yPtr *int) {
	*yPtr = 100
}
