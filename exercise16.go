package main

import (
	"errors"
	"fmt"
)

type Employee struct {
	Name       string
	Salary     float64
	Department string
}

func stats(emp []Employee) (Employee, float64, error) {
	if len(emp) == 0 {
		return Employee{}, 0, errors.New("Empty slice")
	}

	// Initialize with first employee
	maxSalary := emp[0]
	totalSalary := 0.0

	// Find highest salary and calculate total
	for _, e := range emp {
		if e.Salary > maxSalary.Salary {
			maxSalary = e
		}
		totalSalary += e.Salary
	}

	// Calculate average
	averageSalary := totalSalary / float64(len(emp))

	return maxSalary, averageSalary, nil
}

func main() {
	employees := []Employee{
		{Name: "Alice", Salary: 75000, Department: "Engineering"},
		{Name: "Bob", Salary: 85000, Department: "Marketing"},
		{Name: "Charlie", Salary: 92000, Department: "Engineering"},
		{Name: "Diana", Salary: 68000, Department: "HR"},
		{Name: "Eve", Salary: 78000, Department: "Marketing"},
	}

	highest, avg, err := stats(employees)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Printf("Highest Salary: $%.2f (Employee: %s, Dept: %s)\n",
		highest.Salary, highest.Name, highest.Department)
	fmt.Printf("Average Salary: $%.2f\n", avg)
}
