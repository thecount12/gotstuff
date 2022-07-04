package main

import "fmt"

func main() {
	// or 0, 50, 10, 10 if you know the order
	rect1 := Rectangle{leftX: 0, topY: 50, height: 10, width:10}
	fmt.Println("Rec is", rect1.width, "wide")
	fmt.Println("Area of rectang=", rect1.area())
}

type Rectangle struct {
	
	leftX float64
	topY float64
	height float64
	width float64
}

// *Rectangle is receiver
func (rect *Rectangle) area() float64 {
	return rect.width * rect.height
}

