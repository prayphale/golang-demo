package main

import (
	"fmt"
)

func updateName(x *string) {
	*x = "Good"
}

func main() {
	name := "Morning"

	// updateName(name)

	// fmt.Println("memry add:", &name)

	m := &name

	// fmt.Println(m)

	// fmt.Println("value if memory address:", *m)

	fmt.Println(name)
	updateName(m)

	fmt.Println(name)
}
