package main

import "fmt"

//Exercise 3 Rectangle Area

type Rectangle struct {
	length float64
	width  float64
}

func (r Rectangle) Area() float64 {
	return r.length * r.width
}
func main() {
	rect := Rectangle{length: 5.0, width: 3.0}
	area := rect.Area()
	fmt.Printf("The area of the rectangle is: %.2f\n", area)
}
