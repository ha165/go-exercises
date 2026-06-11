package main

import "fmt"

// Exercisse 2 Adult checker

type Person struct {
	Name string
	Age int
}

 func isAdult(p Person) bool {
	return p.Age >= 18
}

func main() {
person := Person{
	Name: "Alice",
	Age: 20,
}

if isAdult(person) {
	fmt.Printf("%s is an adult.\n", person.Name)
} else {
	fmt.Printf("%s is not an adult.\n", person.Name)
 }
}