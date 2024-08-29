package main

import (
	"log"
)

// crate a struc for User
type User struct {
	FirstName string
	LastName  string
}

func main() {
	// for loop

	for i := 0; i <= 10; i++ {
		//log.Println(i)
	}

	// string slice
	mySlice := []string{"a", "b", "c", "d", "e"}

	// crate a slice using make and map with sample animal names

	mySlice2 := make([]string, 5)

	mySlice2[0] = "dog"
	mySlice2[1] = "cat"
	mySlice2[2] = "fish"
	mySlice2[3] = "bird"
	mySlice2[4] = "rabbit"

	for i, x := range mySlice {
		log.Println(i, x)
	}
	// skip the index using _ or blank identifier
	for _, x := range mySlice {
		log.Println(x)
	}
	for i, x := range mySlice2 {
		log.Println(i, x)
	}

	// crate a map with sample animal names

	myMap := make(map[string]string)
	myMap["dog"] = "dog name"
	myMap["cat"] = "cat name"
	myMap["fish"] = "fish name"
	myMap["bird"] = "bird name"
	myMap["rabbit"] = "rabbit name"

	for key, value := range myMap {
		log.Println(key, value)
	}

	var mySlice3 []User

	u1 := User{
		FirstName: "John",
		LastName:  "Doe",
	}

	u2 := User{
		FirstName: "Jane2",
		LastName:  "Doe2",
	}

	mySlice3 = append(mySlice3, u1, u2)

	for i, x := range mySlice3 {
		log.Println(i, x.FirstName, x.LastName)
	}

}
