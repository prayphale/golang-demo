package main

import "fmt"

func main() {

	for num := 1; num <= 5; num++ {
		if num == 4 {
			break //exit for loop from here
		}
		fmt.Println(num)
	}

	names := []string{"hello", "hii", "bye", "ok"}
	for i := 1; i < len(names); i++ {
		fmt.Println(names[i])
	}

	// Range Keyword
	for index, value := range names {
		fmt.Printf("the value at index %v is %v \n", index, value)
	}

	for _, value := range names {
		fmt.Printf("the value at index %v \n", value)
	}

	fmt.Println(names)

}
