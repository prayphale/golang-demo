package main

import "fmt"

func main() {

	// Print
	fmt.Print("Hello, ")
	fmt.Print("Pooja \n")

	fmt.Println("Hello, Pooja")
	fmt.Println("Good, Morning")

	age := 16
	name := "Pooja"
	fmt.Println("My age is", age, "and my name is", name)

	// Formatted String (Printf)

	fmt.Printf("My age is %v and my name is %v \n", age, name)

	fmt.Printf("My age is %q and my name is %q \n", age, name)

	fmt.Printf("Age is of type %T \n", age)

	// Float
	fmt.Printf("Scored %f \n", 225.5)

	fmt.Printf("Scored %0.1f \n", 225.55) // limit

	// SprintF (save Formatted Strings)

	var str = fmt.Sprintf("My age is %v and my name is %v \n", age, name)

	fmt.Println("Saved string is: ", str)

}
