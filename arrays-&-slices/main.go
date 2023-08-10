package main

import "fmt"

func main() {

	// array
	var ages [3]int = [3]int{20, 25, 30}

	fmt.Println(ages)

	var ages1 = [3]int{20, 25, 30}

	// Array Shorthand
	names := [4]string{"tt", "ty", "oo", "poo"}

	names[1] = "yyy"

	fmt.Println(ages, len(ages1))
	fmt.Println(names, len(names))

	// slices

	var scores = []int{100, 50, 60}
	scores[2] = 25
	scores = append(scores, 85)

	fmt.Println(scores, len(scores))

	// slice range
	// A slice can be created by slicing a given array and by specifying the lower and higher bound.

	rangeOne := names[1:3]
	rangeTwo := names[2:]
	rangeThree := names[:3]

	fmt.Println(rangeOne, rangeTwo, rangeThree)

	rangeOne = append(rangeOne, "test")

	fmt.Println(rangeOne)

}
