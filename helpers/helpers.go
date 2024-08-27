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