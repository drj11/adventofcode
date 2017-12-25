package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const N = 79 // size of grid
type Grid [N][N]bool
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
			g[nr+o][i+o] = c == '#'
		}
		nr += 1
	}
	c = Carrier{N / 2, N / 2, [2]int{0, -1}}
	dump(g)

	infected := 0

	for i := 0; i < 10000; i++ {
		infected += burst(&g, &c)
	}
	fmt.Println(infected)
}

func dump(g Grid) {
	for _, row := range g {
		for _, infected := range row {

			if infected {
				fmt.Print("#")
			} else {
				fmt.Print(".")
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

func (c *Carrier) forward() {
	c.X += c.Heading[0]
	c.Y += c.Heading[1]
}

func burst(g *Grid, c *Carrier) int {
	if g[c.Y][c.X] {
		fmt.Println("infected")
		c.right()
	} else {
		fmt.Println("clean")
		c.left()
	}
	g[c.Y][c.X] = !g[c.Y][c.X]
	v := 0
	if g[c.Y][c.X] {
		v = 1
	}
	c.forward()
	return v
}
