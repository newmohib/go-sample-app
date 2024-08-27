package main

import (
	"fmt"

	"github.com/newmohib/go-sample-app/helpers"
)

func main()  {
	fmt.Println("Hello, World!")
	var myVar helpers.SomeType
	myVar.TypeName = "hello"
	myVar.TypeNumber = 1
	fmt.Println(myVar.TypeName)
}