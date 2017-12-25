package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const N = 1079      // size of grid
type Grid [N][N]int // 0: Clean, 1: Weakened, 2: Infected, 3: Flagged
type Carrier struct {
	X       int    // +ve to the right
	Y       int    // +ve down
	Heading [2]int // unit vector in direction of heading
}

func main() {

	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var g Grid
	var c Carrier

	lines := bufio.NewScanner(inp)
	nr := 0
	for lines.Scan() {
		line := lines.Text()
		// Assume square grid
		o := (N - len(line)) / 2
		for i, c := range line {
			if c == '#' {
				g[nr+o][i+o] = 2
			}
		}
		nr += 1
	}
	c = Carrier{N / 2, N / 2, [2]int{0, -1}}
	dump(g)

	infected := 0

	for i := 0; i < 10000000; i++ {
		infected += burst(&g, &c)
	}
	fmt.Println(infected)
}

func dump(g Grid) {
	for _, row := range g {
		for _, state := range row {
			switch state {
			case 0:
				fmt.Print(".")
			case 1:
				fmt.Print("W")
			case 2:
				fmt.Print("#")
			case 3:
				fmt.Print("#")
			}
		}
		fmt.Println("")
	}
}

func (c *Carrier) left() {
	c.Heading[0], c.Heading[1] = c.Heading[1], -c.Heading[0]
}

func (c *Carrier) right() {
	c.left()
	c.left()
	c.left()
}

func (c *Carrier) reverse() {
	c.left()
	c.left()
}

func (c *Carrier) forward() {
	c.X += c.Heading[0]
	c.Y += c.Heading[1]
}

func burst(g *Grid, c *Carrier) int {
	switch g[c.Y][c.X] {
	case 0:
		c.left()
	case 1:

	case 2:
		c.right()
	case 3:
		c.reverse()
	}
	g[c.Y][c.X] = (g[c.Y][c.X] + 1) % 4
	v := 0
	if g[c.Y][c.X] == 2 {
		v = 1
	}
	c.forward()
	return v
}
