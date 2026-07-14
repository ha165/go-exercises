package main

import "fmt"

func findStats(nums []int) (min, max int, avg float64, err error) {
	if len(nums) == 0 {
		return 0, 0, 0, fmt.Errorf("slice is empty")
	}

	// Initialize min and max with first element
	min = nums[0]
	max = nums[0]
	sum := 0

	// Loop through all elements
	for _, num := range nums {
		if num < min {
			min = num
		}
		if num > max {
			max = num
		}
		sum += num
	}

	avg = float64(sum) / float64(len(nums))
	return min, max, avg, nil
}

func main() {
	numbers := []int{3, 7, 2, 9, 1, 5}

	min, max, avg, err := findStats(numbers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Min: %d, Max: %d, Average: %.2f\n", min, max, avg)
}
