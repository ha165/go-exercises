// You can edit this code!
// Click here and start typing.
package main

import (
	"errors"
	"fmt"
)

func filterEven(numbers []int) ([]int, error) {
	if len(numbers) == 0 {
		return []int{}, errors.New("Empty SLice")
	}
	evens := make([]int, 0, len(numbers)/2)
	for _, num := range numbers {
		if num%2 == 0 {
			evens = append(evens, num)
		}
	}
	return evens, nil
}

func main() {
	numbers := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	evens, err := filterEven(numbers)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Even numbers:", evens)
}
