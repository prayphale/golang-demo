package main

import (
	"fmt"
)

func updateName(x string) string {
	x = "chekc"
	return x
}

func main() {
	name := "test"

	// name = updateName(name)
	updateName(name)
	fmt.Println(name)

}
