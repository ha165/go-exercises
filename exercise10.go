package main

import "fmt"

// EXercise10 Interfaces

type Animal interface {
	makeSound()
}
type Dog struct {
	Name string
}
type Cat struct {
	Name string
}

func (d Dog) makeSound() {
	fmt.Println("woof")
}
func (c Cat) makeSound() {
	fmt.Println("meow")
}
func makeAnimalSound(a Animal) {
	a.makeSound()
}
func main() {
	d := Dog{Name: "Rex"}
	c := Cat{Name: "Whiskers"}
	makeAnimalSound(d)
	makeAnimalSound(c)
}
