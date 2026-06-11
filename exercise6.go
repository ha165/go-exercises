package main

import "fmt"

// Exercise6 Create a function for Grade Calculator

func gradeCalc(score float64) string {
	var grade string
	if score >= 90 {
		grade = "A"
	} else if score >= 80 {
		grade = "B"
	} else if score >= 70 {
		grade = "C"
	} else if score >= 60 {
		grade = "D"
	} else {
		grade = "F"
	}
	return grade
}
func main() {
	score := 85.0
	grade := gradeCalc(score)
	fmt.Printf("The grade for a score of %.2f is: %s\n", score, grade)
}
