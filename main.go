package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"

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

	myJson := `[
		{
			"first_name": "John",
			"last_name": "Doe",
			"hair_color": "brown",
			"has_dog": true
		},
		{
			"first_name": "John 2",
			"last_name": "Doe 2",
			"hair_color": "brown 2",
			"has_dog": false
		}
	]`

	var unmarshalled []helpers.Person
	err := json.Unmarshal([]byte(myJson), &unmarshalled) // Unmarshal => parsing

	if err != nil {
		log.Println("Error unmarshalling json", err)
	}else {
		log.Printf("unmarshalled: %v", unmarshalled)
	}

	//write json form a struct
	var mySLice []helpers.Person

	//write Single json
	var m1 helpers.Person
	
	m1.FirstName = "John 3"
	m1.LastName = "Doe 3"
	m1.HairColor = "brown 3"
	m1.HasDog = true

	mySLice = append(mySLice, m1)

	// write Single json
	var m2 helpers.Person
	
	m2.FirstName = "John 4"
	m2.LastName = "Doe 4"
	m2.HairColor = "brown 4"
	m2.HasDog = true

	mySLice = append(mySLice, m2)

	// append multiple slice in unmarshalled
	unmarshalled = append(unmarshalled, m1, m2)

	log.Printf("mySLice: %v", mySLice)
	log.Printf("unmarshalled: %v", unmarshalled)

	newJson, err := json.MarshalIndent(unmarshalled, "", " ")
	if err != nil {
		log.Println("Error marshalling json", err)
	}else{
		fmt.Println(string(newJson))
	}
	
	// Loop through the unmarshalled data
	// _: _ is a blank identifier for ignoring the index

	for _, person := range unmarshalled {
		fmt.Printf("First Name: %s, Last Name: %s, Hair Color: %s, Has Dog: %t\n",
			person.FirstName, person.LastName, person.HairColor, person.HasDog)
	}

	// divide is a fucntions

	result, err := divide(5, 2)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("Divided Resutl", result)

}


func divide(x, y float32) (float32, error) {
	var result float32

	if y  == 0 {
		return result, errors.New("cannot divide by zero")
	}
	result = x / y
	return result, nil
	
}