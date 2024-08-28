package main

import (
	"errors"
	"fmt"
)


func main()  {
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