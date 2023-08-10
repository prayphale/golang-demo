package main

import "fmt"

func main() {
	//using var
	var emplist = map[int]string{1: "Steve", 2: "Roy", 3: "Amar", 4: "Nancy", 5: "Abdul"}
	fmt.Println(emplist)

	//shorthand syntax
	empsal := map[string]int{"Steve": 60000, "Roy": 70000, "Amar": 75000, "Nancy": 60000, "Abdul": 65000}
	fmt.Println(empsal)

	// Check If Key Exists
	emp := map[int]string{1: "Steve", 2: "Roy", 3: "Amar", 4: "Nancy", 5: "Abdul"}

	val, ok := emp[3]
	fmt.Println(val, ok)

	val2, ok := emp[7]
	fmt.Println(val2, ok)

	// delete
	emp2 := map[int]string{1: "Steve", 2: "Roy", 3: "Amar", 4: "Nancy", 5: "Abdul"}
	delete(emp2, 1)
	delete(emp2, 2)

	fmt.Println(emp2)

}
