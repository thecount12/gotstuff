package main

import "fmt"

func main() {
	
	fmt.Println("hello world")
	var age int = 40 // could also int64
	var favNum float64 = 1.6180339
	randNum := 1
	fmt.Println(age, " ", favNum, "rand", randNum)
	var num1 = 1.000
	var num99 = 2.23
	fmt.Println("new Num:", num1 + num99)
	const pi float64 = 3.14159265
	var (
		varA = 2
		varB = 3
		)
	fmt.Println(varA + varB)
	var myName string = "William Gunnells"
	fmt.Println(len(myName))
	var isOver40 = true
	fmt.Printf("%.3f \n", pi)
	fmt.Printf("%T \n", pi)
	fmt.Printf("%t \n", isOver40)
	fmt.Printf("%d \n ", 100) // integer
	fmt.Printf("%b \n ", 100)  // binary
	fmt.Printf("%c \n ", 44)  // comma represents 44
	fmt.Printf("%x \n ", 100)  // hex
	fmt.Printf("%e \n", pi) // scientific

	
}
