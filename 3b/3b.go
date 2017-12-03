package main

import (
	"fmt"
	"log"
	"math"
)

func main() {
	input := 312051

	x, y := cartesian(input)

	fmt.Println(x, y)

	fmt.Println(math.Abs(float64(x)) + math.Abs(float64(y)))
}

// Convert from spiral system (starting at 1) to cartesian.
// Where starting cell 1 is at (0,0)
// and y increases down the page.
func cartesian(spiral int) (x, y int) {
	// Imagine the grid of memory cells is expanded
	// until it is a square of cells just big enough
	// to include the input cell.

	// The input lies on on the outside edge
	// of this square grid, which we will say is of size n.
	// n is odd, and the bottom-right cell is n*n.
	// Find n.
	n := int(math.Ceil(math.Sqrt(float64(spiral))))
	n |= 1

	// The 4 corners of the square grid are:
	//                 side 2
	//     n*n - 2*n  ~~~~~~~~  n*n - 3*n
	//         '                    '
	//         '                    '
	// side 1  '                    '  side 3
	//         '                    '
	//         '                    '
	//     n*n - n  ~~~~~~~~~~~~~  n*n
	//                 side 0

	// Counting backwards along the spiral,
	// starting from the bottom-right cell,
	// the distance to our input cell is d.
	d := n*n - spiral
	if d >= 4*n {
		log.Fatal("oh no")
	}

	// side is 0 (bottom) to 3 (right), as per diagram.
	side := d / n
	// Convert d to a distance from the corner
	// that is at least as large as spiral.
	d = d % n

	// Now we are nearly done.
	// n and d are the coordinates of our input cell,
	// but shifted from what we need to compute distance to
	// cell 1.

	fmt.Println(n, d)

	p := n / 2
	q := d - p

	// Convert to x,y.
	switch side {
	case 0:
		y = p
		x = -q
	case 1:
		x = -p
		y = -q
	case 2:
		y = -p
		x = q
	case 3:
		x = p
		y = q
	}

	return x, y
}
