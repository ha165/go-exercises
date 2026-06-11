package main

import "fmt"

// Exercise7 Create a function for Fibonacci sequence

func fibonacci(n int) []int {
	sequence := make([]int, n)
	for i := 0; i < n; i++ {
		if i == 0 {
			sequence[i] = 0
		} else if i == 1 {
			sequence[i] = 1
		} else {
			sequence[i] = sequence[i-1] + sequence[i-2]
		}
	}
	return sequence
}
func main() {
	n := 10
	sequence := fibonacci(n)
	fmt.Printf("The Fibonacci sequence for %d terms is: %v\n", n, sequence)
}
