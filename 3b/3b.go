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

	for i := 1; i < 10; i++ {
		fmt.Println(cartesian(i))
	}
}

// Convert from spiral system (starting at 1) to cartesian.
// Where starting cell 1 is at (0,0)
// and y increases down the page.
func cartesian(spiral int) (x, y int) {
	// Imagine the grid of memory cells is expanded
	// until it is a square of cells just big enough
	// to include the input cell.

	if spiral == 1 {
		return 0, 0
	}

	// The input lies on on the outside edge
	// of this square grid, which we will say is of size n.
	// n is odd, and the bottom-right cell is n*n.
	// Find n.
	n := int(math.Ceil(math.Sqrt(float64(spiral))))
	n |= 1
	k := n - 1

	// The 4 corners of the square grid are:
	//                 side 2
	//     n*n - 2*k  ~~~~~~~~  n*n - 3*k
	//         '                    '
	//         '                    '
	// side 1  '                    '  side 3
	//         '                    '
	//         '                    '
	//     n*n - k  ~~~~~~~~~~~~~  n*n
	//                 side 0
	// Where k is n-1.

	// Counting backwards along the spiral,
	// starting from the bottom-right cell,
	// the distance to our input cell is d.
	d := n*n - spiral
	if d > 4*k {
		log.Fatal("oh no")
	}

	// side is 0 (bottom) to 3 (right), as per diagram.
	side := d / k
	// Convert d to a distance from the corner
	// that is at least as large as spiral.
	d = d % k

	// Now we are nearly done.
	// n and d are the coordinates of our input cell,
	// but shifted from what we need to compute distance to
	// cell 1.

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
