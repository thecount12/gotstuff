package main

import (
	//"log"
	"fmt"
)

// list of the functions that satisfies the animal type must have
type Animal interface {
	Says() string
	NumberOfLegs() int
	
}

type Dog struct {
	Name string
	Breed string
}

type Gorilla struct {
	Name string
	Color string
	NumberOfTeeth int
}

func main() {
	// 1. create Animal interface
	// 2. create simple structs
	// 3. add data to either of the structs
	dog := Dog{Name: "Samson", Breed: "German Shephered",}
	// 5. will error becuase dog is messed up
	PrintInfo(&dog) // will give error without the receiver and no pointers
	// 9 add data
	gorilla := Gorilla{Name: "Jock", Color: "grey", NumberOfTeeth: 38,}
	// 10, print  it will give error without reciver for gorilla
	PrintInfo(&gorilla)
}

// 4. create a function that uses interface
func PrintInfo(a Animal) {
	fmt.Println("this animal says", a.Says(), "and has", a.NumberOfLegs(), "legs")
}
	
// 6. receiver for dog says without Pointeras
func (d *Dog) Says() string {
	return "Woof"
}
// 7. receiver for dog legs without Pointers
// 8. now make Dog *Dog for both and &dereference above
func (d *Dog) NumberOfLegs() int {
	return 4
}
// 11 both below
func (g *Gorilla) Says() string {
	return "Ugh"
}

func (g *Gorilla) NumberOfLegs() int {
	return 2
}
