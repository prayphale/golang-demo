package main

import "fmt"

func greet(name string) {
	fmt.Printf("Good Morning %v \n", name)
}

// Function with Input Parameters and Return Value
func Sum1(x int, y int) int {
	return x + y
}

func cycleNames(n []string, f func(string)) {
	for _, v := range n {
		f(v)
	}
}

func main() {
	// greet("Pooja")

	result := Sum1(5, 10)
	fmt.Println(result)

	cycleNames([]string{"abc", "qwe", "poo"}, greet)
}

// multiple arguments
