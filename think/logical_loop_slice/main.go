package main

import "fmt"

func main() {
	fmt.Println("true && false = ", true && false)
	fmt.Println("true || false = ", true || false)
	fmt.Println("!true = ", !true)
	// loops
	i := 1
	for i < 10 {
		fmt.Println(i)
		i++
	}
	fmt.Println("--")
	for j:= 0; j<5; j++ {
		fmt.Println(j)
	}
	Age := 18
	if Age >= 16 {
		fmt.Println("You can Drive")
	} else if Age >= 18 {
		fmt.Println("You can vote") // even though it hits 18
	} else {
		fmt.Println("you Can Have Fun")
	}
	yourAge := 5
	switch yourAge {
		case 16 : fmt.Println("Go Drive")
		case 18 : fmt.Println("Go Vote")
		default : fmt.Println("Go Have Fun")
	}
	// array
	var favNum2[5] float64
	favNum2[0] = 163
	favNum2[1] = 78557
	favNum2[2] = 691
	favNum2[3] = 3.141
	favNum2[4] = 1.618
	
	fmt.Println(favNum2[3])
	
	favNum := [5]float64 {1,2,3,4,5}
	for i, value := range(favNum) {
		fmt.Printf("index: %d value: %v\n", i, value)
		// or fmtPrintln(value, i)
	}
	// slize is an array without size
	numSlice := []int {5,4,3,2,1} // 0..5 it ignores 5 
	numSlice2 := numSlice[3:5]
	fmt.Println("numSlice2[0] =", numSlice2[0]) //2
	fmt.Println("numSlice2[1] =", numSlice2[1]) //1
	fmt.Println("numSlice[0:2] =", numSlice[:2]) // [5 4]
	// not defined set of values
	numSlice3 := make([]int, 5, 10) // int array default val 5 max size 10
	copy(numSlice3, numSlice)
	fmt.Println(numSlice3[0])
	fmt.Println(numSlice3[1])
	numSlice3 = append(numSlice3, 0, -1) // append 0 and -1
	fmt.Println(numSlice3[6])
	
}


