package main

import "fmt"

type Student struct {
	Name  string
	Grade string
	Age   int
}

func report(s Student) {
	fmt.Printf("Name: %s, Grade: %s, Age:%d", s.Name, s.Grade, s.Age)
}

func main() {
	s1 := Student{
		Name:  "Harmony",
		Grade: "A",
		Age:   25,
	}
	s2 := Student{
		Name:  "Wallace",
		Grade: "C",
		Age:   15,
	}

	report(s1)
	report(s2)
}
