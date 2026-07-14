// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {

	phoneBook := make(map[string]string)

	// Add entries
	phoneBook["Alice"] = "555-1234"
	phoneBook["Bob"] = "555-5678"
	phoneBook["Charlie"] = "555-9012"

	//search

	name := "Alice"
	phone, exists := phoneBook[name]

	if exists {
		fmt.Println(name, phone)
	} else {
		fmt.Println("Not Found")
	}

	// DELETE
	delete(phoneBook, "Bob")

	//add
	phoneBook["Harmony"] = "0702486902"

	for name, phone := range phoneBook {
		fmt.Printf("Name:%s and PhoneNumber:%s\n", name, phone)
	}

}
