package main

import (
	"log"
	"time"
)

var s = "Seven"

type User struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

func main()  {
	s2 := "Six"
	 s := "Eight"
	log.Println("s is",s)
	log.Println("s is",s2)
	saySomething("Hello World")

	user := User {
		FirstName: "John",
		LastName: "Doe",
		PhoneNumber: "555-555-5555",
		Age: 30,
		BirthDate: time.Now(),
	}

	log.Println(user.FirstName, user.LastName, user.PhoneNumber, user.Age, "Birth Date:", user.BirthDate)
}

func saySomething(s3 string) (string, string)  {
	log.Println("s from the saySomethign func is ",s)
	return s, "hello"
}

// public variable will start with small letter
// private variable will start with capital letter

// public variable

var TestVar = "hello"
// public function
func Public()  {
	
}

// public struct
type PublicType struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

// Private variable

var testVar = "hello"
// private function
func Private()  {
	
}

// private struct
type PrivateType struct {
	FirstName string
	LastName string
	PhoneNumber string
	Age int
	BirthDate time.Time
}

