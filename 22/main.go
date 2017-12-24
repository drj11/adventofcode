package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const N = 79 // size of grid
var g [N][N]bool

func main() {

	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

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
	dump()
}

func dump() {
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
