package main

import "fmt"
// Exercise  create a student struct with name age and grade

type Student struct {
    Name string
    Age int
    Grade string
}
func main() {
    
    std := Student{
        Name: "Harmony",
        Age : 24,
        Grade : "A",
    }

    fmt.Printf("Student Name: %s\n", std.Name)
    fmt.Printf("Student Age: %d\n", std.Age)
    fmt.Printf("Student Grade: %s\n", std.Grade)
} 