package main

import (
	"fmt"
	"log"
	"math"
)

const N = 13 // size of grid

func main() {
	input := 312051
	answer := 0

	var g [N][N]int

	// get element from g
	G := func(x, y int) int {
		x += N / 2
		y += N / 2
		return g[y][x]
	}
	// set element in g
	S := func(x, y, v int) {
		x += N / 2
		y += N / 2
		g[y][x] = v
	}
	// print
	P := func() {
		for y := 0; y < N; y++ {
			fmt.Println(g[y])
		}
	}

	S(0, 0, 1)
	for i := 2; i < 82; i++ {
		x, y := cartesian(i)
		ns := neighbours(x, y)
		sum := 0
		for _, neighbour := range ns {
			nx, ny := neighbour[0], neighbour[1]
			sum += G(nx, ny)
		}
		S(x, y, sum)
		if sum > input && answer == 0 {
			answer = sum
		}
	}
	P()
	fmt.Println(answer)
}

// return coordinates of neighbours
func neighbours(x, y int) [][2]int {
	var r [][2]int
	for j := -1; j < 2; j++ {
		for i := -1; i < 2; i++ {
			r = append(r, [2]int{x + i, y + j})
		}
	}
	return r
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
