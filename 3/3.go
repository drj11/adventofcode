package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	input := 312051

	// Imagine the grid of memory cells is expanded
	// until it is a square of cells just big enough
	// to include the input cell.

	// The input lies on on the outside edge
	// of this square grid, which we will say is of size n.
	// n is odd, and the bottom-right cell is n*n.
	// Find n.
	n := int(math.Ceil(math.Sqrt(float64(input))))
	n |= 1

	// The 4 corners of the square grid are:
	// n*n - 2*n  ~~~~~~~~  n*n - 3*n
	//     '                    '
	//     '                    '
	//     '                    '
	//     '                    '
	// n*n - n  ~~~~~~~~~~~~~  n*n

	// Counting backwards along the spiral,
	// starting from the bottom-right cell,
	// the distance to our input cell is d.
	d := n*n - input
	if d >= 4*n {
		log.Fatal("oh no")
	}

	// Convert d to a distance from its nearest corner.
	d = d % n

	// Now we are nearly done.
	// n and d are the coordinates of our input cell,
	// but shifted from what we need to compute distance to
	// cell 1.

	fmt.Println(n, d)

	p := n / 2
	q := d - p
	if q < 0 {
		q = -q
	}

	fmt.Println(p + q)
}
