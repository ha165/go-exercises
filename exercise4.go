package main

import "fmt"

// Exercisse 4 Largest Number
func largest(a, b, c, d int) int {
	var largest int
	if a > b && a > c && a > d {
		largest = a
	} else if b > a && b > c && b > d {
		largest = b
	} else if c > a && c > b && c > d {
		largest = c
	} else {
		largest = d
	}
	return largest
}
func main() {
	num1 := 10
	num2 := 25
	num3 := 15
	num4 := 5

	result := largest(num1, num2, num3, num4)
	fmt.Printf("The largest number among %d, %d, %d, and %d is: %d\n", num1, num2, num3, num4, result)
}
