package main

import (
	"fmt"

	"github.com/newmohib/go-sample-app/helpers"
)
const numPool = 10

func CalculateValue(intChan chan int)  {
	randomNumber := helpers.RandomNumber(numPool)

	intChan <- randomNumber


	
}

func main()  {
	fmt.Println("Hello, World!")
	var myVar helpers.SomeType
	myVar.TypeName = "hello"
	myVar.TypeNumber = 1
	fmt.Println(myVar.TypeName)

	intChan := make(chan int)
	defer close(intChan)
	go CalculateValue(intChan)

	num := <- intChan
	fmt.Println(num)
	 

}