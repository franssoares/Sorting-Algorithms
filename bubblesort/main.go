package main

import (
	"fmt"
)

func main() {

	var bub bubbleSort
	slice := []int{4, 2, 7, 1, 5, 6, 5, 234, -2}

	fmt.Println("Slice original antes:", slice)

	// Chamar o método com a slice
	bub.sortAscending(slice)

	fmt.Println("Slice original depois:", slice)
}
