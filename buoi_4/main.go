package main

import (
	"fmt"
)

func swap(a, b *int) {
	*a, *b = *b, *a
}
func main() {
	a := 5
	b := 10
	fmt.Println("Before swap:", a, b)
	swap(&a, &b)
	fmt.Println("After swap:", a, b)
	fmt.Print("Swapping completed successfully.\n")
}
