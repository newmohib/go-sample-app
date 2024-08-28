package main

import (
	"log"
)


func main()  {

	var saySomethingElse string
	var i int
	var world string
	
	
	wahtToSay, _ := saySomething("Hello World")

	log.Println(saySomethingElse)
	log.Println(wahtToSay)
	saySomethingElse, world = saySomething("Goodbye")

	i = 7
	i = 8
	log.Println(saySomethingElse)
	log.Println(i)
	log.Println(world)
	
}

func saySomething(s string) (string, string)  {
	return s, "hello"
}