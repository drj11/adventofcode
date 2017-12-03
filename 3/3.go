package main

import (
	"fmt"
	"math"
)

func main() {
	input := 312051

	// The input lies on on the outside edge
	// of square grid of size n,
	// where n is odd, and the bottom-right
	// cell is n*n.
	// Find n
	n := int(math.Ceil(math.Sqrt(float64(input))))
	n |= 1

	fmt.Println(n)
}
