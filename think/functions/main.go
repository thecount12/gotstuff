package main 

import "fmt"

func main() {
	// sum all numbers
	listNums := []float64{1,2,3,4,5}
	fmt.Println("Sum: ", addThemUp(listNums))
	
	// interesting add function
	num1, num2 := next2Values(5)
	fmt.Println(num1, num2)
	
	// undefined or variatic
	fmt.Println(subtractThem(1,2,3,4,5))
	
	// function inside of function or closure
	num3 := 3
	doubleNum := func() int {
		num3 *= 2 // has access to num3
		return num3
	}
	fmt.Println(doubleNum())
	fmt.Println(doubleNum()) // double again
	
	// recursion factorial example
	fmt.Println(factorial(5)) //4=24 5=120
	
	// defer
	defer printTwo() // prints last after everything
	printOne()
	
	// safe division
	fmt.Println(safeDiv(3,0))
	fmt.Println(safeDiv(3,2))
	
	// demonstrate panic
	demPanic()
}

func addThemUp(numbers []float64) float64 {
	sum := 0.0
	for _, val := range numbers {
		sum += val
	}
	return sum // should be 15
}

func next2Values(number int) (int, int) {
	return number+1, number+2  // should return 6 7
	
}

func subtractThem(args ...int) int {
	finalValue := 0
	for _, value := range args {
		finalValue -= value
	}
	return finalValue // should be -15
}

func factorial(num int) int {
	if num == 0 {
		return 1
	}
	return num * factorial(num -1)
	
}

func printOne() { fmt.Println(1)}
func printTwo() { fmt.Println(2)}

func safeDiv(num1, num2 int) int {
	defer func() {
		fmt.Println(recover()) // catch if error occur
		// recover() // remove Println() program will just continue
	}()

	solution := num1 / num2 // if 2/0 error recover
	return solution

}

func demPanic() {
	
	defer func() {
		fmt.Println(recover())
	}()
	panic("PANIC")
}






