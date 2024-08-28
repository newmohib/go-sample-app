package main

import (
	"log"
)

func main()  {
	myString := "hello"

	log.Println(myString)

	changeUsingPointer(&myString)
	
	log.Println(myString)
}

func changeUsingPointer(s *string)  {

	log.Println("s is set to ", s)

	newValue := "goodbye"
	
	*s = newValue
	
}