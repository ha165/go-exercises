package main

import (
	"errors"
	"fmt"
)

// Exercise8 user loggin system
type User struct {
	Username string
	Password string
}

func login(user User, username, password string) error {
	if username == user.Username && password == user.Password {
		return nil
	}
	return errors.New("Login Failed")
}

func main() {
	user := User{Username: "admin", Password: "password123"}
	username := "admin"
	password := "password124"

	err := login(user, username, password)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Login Successful!")
	}
}
