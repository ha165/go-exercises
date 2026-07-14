package main

import "fmt"

func countWords(words []string) map[string]int {
	wordCount := make(map[string]int)

	for _, word := range words {
		wordCount[word]++
	}

	return wordCount
}

func main() {
	words := []string{"Go", "Go", "Java", "Go", "Rust", "Rust"}

	counts := countWords(words)

	fmt.Println("Word Count:")
	for word, count := range counts {
		fmt.Printf("%s -> %d\n", word, count)
	}
}
