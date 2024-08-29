package main

import (
	"log"
)

func main() {
	myNum := 100
	isTrue := false

	if isTrue {
		log.Println("isTrue is ", isTrue)
	} else {
		log.Println("isTrue is ", isTrue)
	}

	if myNum > 99 && !isTrue {
		log.Println("myNum is greater than 99 and isTrue is set to true")
	} else if myNum < 100 && isTrue {
		log.Println("myNum is ", 1)

	} else if myNum == 101 || isTrue {
		log.Println("myNum is ", 2)

	} else if myNum > 100 && !isTrue {
		log.Println("myNum is ", 3)

	}

	// switch statement is use decision structures
	myVar := "cat"

	switch myVar {

	case "cat":
		log.Println("myVar is ", "cat")
		// fallthrough
	case "dog":
		log.Println("myVar is ", "dog")
		// fallthrough
	case "fish":
		log.Println("myVar is ", "fish")
		// fallthrough
	default:
		log.Println("myVar is ", "default")
	}

}
