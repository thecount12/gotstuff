package main

import (
	"fmt"
	"reflect"
)


func main() {
	ar := []int{5, 1, -1, 2, 3}
	newAr := []int{}
	fmt.Println(ar[1])
	for i, num := range(ar) {
		var tempVar = num
		fmt.Printf("index: %v, number: %v\n", i, num)
		fmt.Println(reflect.TypeOf(num))
		fmt.Println(num)
		for _, y := range(ar) {
			if num < y {
				ar = append(newAr, tempVar)
			}
		}
	}
	//fmt.Println(newAr[0])
}
