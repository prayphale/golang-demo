package main

import "fmt"

func main() {

	// strings
	var nameOne string = "check" // Can't use single quotes also can't use any other values like int

	var nameTwo = "Pooja" // Also string type
	var nameThree string  // Just a declaration

	fmt.Println(nameOne, nameTwo, nameThree)

	nameOne = "test"
	nameThree = "Ray" // Here we can't changes the type we just able to pass the value

	fmt.Println(nameOne, nameTwo, nameThree)

	nameFour := "Go lang"
	fmt.Println(nameFour)

	// ints
	var ageOne int = 20
	var ageTwo = 30
	ageThree := 40

	fmt.Println(ageOne, ageTwo, ageThree)

	// bits & memory
	var numOne int8 = 100 //  Range: -128 through 127.
	var numTwo int8 = -128
	var numThree uint = 25 // Only use Range: 0 through 255.

	fmt.Println(numOne, numTwo, numThree)

	// float
	var scoreOne float32 = 25.98
	var scoreTwo float64 = 565656565656565656565656565.5
	scoreThree := 1.5

	fmt.Println(scoreOne, scoreTwo, scoreThree)
}
