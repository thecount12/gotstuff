package main

import "fmt"
import "math"

func main() {
	
	// last after creating the function
	rect := Rectangle{20,50}
	circ := Circle{4}
	
	fmt.Println("Rectangle Area=", getArea(rect))
	fmt.Println("Circle Area=", getArea(circ))
}

type Shape interface {
	area() float64
}

type Rectangle struct {
	height float64
	width float64
}

type Circle struct {
	radius float64
}

func(r Rectangle) area() float64 {
	return r.height * r.width
}

func(c Circle) area() float64 {
	return math.Pi * math.Pow(c.radius, 2)
}

func getArea(shape Shape) float64{
	return shape.area()
}
