package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {

	// Strings packages
	greeting := "hello all friends!"

	fmt.Println(strings.Contains(greeting, "hello"))

	fmt.Println(strings.ReplaceAll(greeting, "hello", "hi"))

	fmt.Println(strings.ToUpper(greeting))

	fmt.Println(strings.Index(greeting, "ll"))

	fmt.Println(strings.Split(greeting, " "))

	fmt.Println("Original value = ", greeting)

	// Sort Package
	ages := []int{45, 20, 35, 75, 60, 50, 35}

	sort.Ints(ages) // altered the original slice
	fmt.Println(ages)

	index := sort.SearchInts(ages, 50)

	fmt.Println(index)

	// Slice string

	names := []string{"box", "max", "fish"}

	sort.Strings(names)
	fmt.Println(names)

	fmt.Println(sort.SearchStrings(names, "max"))
}
