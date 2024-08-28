package helpers

import (
	"math/rand"
	"time"
)

// Declare a struct type
type SomeType struct {
	TypeName string
	TypeNumber int
}

// Declare a Channel

func RandomNumber(n int) int {
	rand.Seed(time.Now().UnixNano())
	value := rand.Intn(n)
	return value
}

type Person struct {
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	HairColor string `json:"hair_color"`
	HasDog bool `json:"has_dog"`
}

