package main

import "fmt"

func main() {

	age := 16

	fmt.Println(age <= 50)
	fmt.Println(age >= 50)
	fmt.Println(age == 50)
	fmt.Println(age != 50)

	if age < 30 {
		fmt.Println("age is less than 30")
	} else if age < 40 {
		fmt.Println("age is less than 40")
	} else {
		fmt.Println("age is not less than 40")
	}

	names := []string{"Pooja", "check", "truw", "abc"}

	for index, value := range names {

		if index == 1 {
			fmt.Println("Contin", index)
			continue
			// The continue keyword is used to skip an iteration and then continue with the next iteration.
		}

		if index > 2 {
			fmt.Println("breakiing at po", index)
			break
		}

		fmt.Printf("the  value at pos %v is %v \n ", index, value)

	}
}
