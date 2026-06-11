package main

import "fmt"

// Exercise8 user loggin system
type User struct {
	Username string
	Password string
}

func login(user User, username, password string) bool {
	if username == user.Username && password == user.Password {
		return true
	}
	return false
}

func main() {
	user := User{Username: "admin", Password: "password123"}
	username := "admin"
	password := "password124"

	if login(user, username, password) {
		fmt.Println("Login successful!")
	} else {
		fmt.Println("Login failed. Invalid username or password.")
	}

}
