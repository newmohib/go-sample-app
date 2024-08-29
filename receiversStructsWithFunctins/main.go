package main

import (
	"log"
)

type User struct {
	FirstName string
}

// Receiver (m *User):
// The receiver m *User means this method is defined on a pointer to User. This allows the method to modify the User instance if needed (though it's not modifying anything in this particular method).
// Using a pointer receiver is common in Go when working with structs, especially when you want to avoid copying the struct or modify the original struct.

// Structs:
// Structs in Go are used to group together fields. In this case, User groups fields related to a user, such as FirstName.

// Methods with Pointer Receivers:
// The method printFirstName has a receiver *User, meaning it's a method that operates on pointers to User.
// This allows methods to modify the original struct if needed or efficiently handle large structs without copying.

// Using a pointer receiver (m *User) is useful when:

// You want to avoid copying the struct each time the method is called, especially if the struct is large.
// also we can usth receivers fucntions for business logic

func (m *User) printFirstName() string {
	return m.FirstName
}

func main() {

	user := User{
		FirstName: "John",
	}

	log.Println(user.printFirstName())
}
