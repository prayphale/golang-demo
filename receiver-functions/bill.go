package main

import "fmt"

type bill struct {
	name  string
	items map[string]int
	tip   int
}

func newBill(name string) bill {
	b := bill{
		name:  name,
		items: map[string]int{"cake": 4, "chocolate": 5},
		tip:   0,
	}

	return b
}

// limited only for bill object
func (b bill) format() string {
	fs := "bill breakdown: \n"
	var total int = 0

	for k, v := range b.items {
		fs += fmt.Sprintf("%v ...%v \n", k+":", v)
		total += v
	}

	fs += fmt.Sprintf("%v ...%v", "total:", total)

	return fs
}
