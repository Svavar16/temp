package main

import "fmt"

// in order to start this, we need to write go run main.go model.go
// we need to run both files at the same time.

func main() {
	fmt.Println("Starting this")
	user1 := new(Users)
	user1.name = "Svavar"

	nameToUpdate := new(Users)
	nameToUpdate.name = "Svavar Updated"

	// With this the second name is updated
	user1.update(nameToUpdate)

	fmt.Println(user1)
	fmt.Println(nameToUpdate)

	user2 := new(Users)
	user2.name = "Svavar Test"

	nameToUpdate1 := new(Users)
	nameToUpdate1.name = "Svavar Test Updated"

	// So if I understand this correctly, then user 2 should be updated.
	nameToUpdate1.update(user2)

	// so here the user2 has changed,
	fmt.Println(user2)

	user1.name = "Momoko"

	user2.update2(user1)

	fmt.Println(user1)
	fmt.Println(user2)

	user1.name = "Momoko"

	user1.update(user2)

	fmt.Println(user1)
	fmt.Println(user2)

	// so this will always change the user that comes in, good to know, does this save more memory?
	
	fmt.Println("This is done")
}
