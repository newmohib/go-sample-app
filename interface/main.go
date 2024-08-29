package main

import (
	"log"
)

// interface dclaration
type Animal interface {
	Says() string
	NumberOfLegs() int
}

// create a struct
type Dog struct {
	Name  string
	Breed string
}

// create another struct
type Gorilla struct {
	Name          string
	Color         string
	NumberOfTeeth int
}

func main() {
	dog := Dog{
		Name:  "Spot",
		Breed: "German Sheperd",
	}
	PrintInfo(dog)

	gorilla := Gorilla{
		Name:          "Rex",
		Color:         "Black",
		NumberOfTeeth: 30,
	}

	PrintInfo(gorilla)

}

// interface implementation receiver function

func (d Dog) Says() string {
	return "woof"
}

func (d Dog) NumberOfLegs() int {
	return 4
}

func (d Gorilla) Says() string {
	return "grunt"
}

func (d Gorilla) NumberOfLegs() int {
	return 2
}

func PrintInfo(a Animal) {
	log.Println("The animal says ", a.Says(), " and has ", a.NumberOfLegs(), " legs")
	log.Println(a.NumberOfLegs())
}
