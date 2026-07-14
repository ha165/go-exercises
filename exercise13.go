package main

import "fmt"

func average(numbers ...float64) float64 {
	if len(numbers) == 0 {
		return 0.0
	}
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return (float64(sum)) / float64(len(numbers))
}

func main() {
	result := average(1, 5, 6, 1, 7, 8, 6, 9)
	fmt.Println(result)
}
