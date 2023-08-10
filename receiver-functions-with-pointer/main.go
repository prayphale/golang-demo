package main

import "fmt"

func main() {
	mybill := newBill("Pooja's bill")

	// mybill.format()
	mybill.updateTip(10)

	fmt.Println(mybill.format())
}
