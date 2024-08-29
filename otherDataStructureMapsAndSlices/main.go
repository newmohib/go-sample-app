package main

import (
	"log"
	"sort"
)

type User struct {
	FirstName string
	LastName  string
}

func main() {

	// create a map
	myMap := make(map[string]string)

	myMap["dog"] = "Samsomn"
	myMap["other-dog"] = "Cassie"

	log.Println(myMap["dog"])
	log.Println(myMap["other-dog"])

	myMap2 := make(map[string]User)

	//This directly initializes the map with values.
	me := User{
		FirstName: "John",
		LastName:  "Doe",
	}

	myMap2["me"] = me
	log.Println(myMap2["me"].FirstName)

	// create a slice as string
	var myString []string
	myString = append(myString, "Test user")
	myString = append(myString, "John")

	// create a slice as int
	var myInt []int

	myInt = append(myInt, 3)
	myInt = append(myInt, 1)
	myInt = append(myInt, 2)

	sort.Ints(myInt)

	log.Println(myString)
	log.Println(myInt)

	// create a slice as int with declaration fixed
	myInt2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	// slice from index 2 to 3 range:
	log.Println(myInt2[2:3])

	// create a slice as string with declaration fixed
	myString2 := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j"}

	// slice from index 2 to 3 range:
	log.Println(myString2[0:1])

}
