package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var m []int

	lines := bufio.NewScanner(inp)
	for lines.Scan() {
		line := lines.Text()
		i, err := strconv.Atoi(line)
		if err != nil {
			log.Fatal(err)
		}
		m = append(m, i)
	}
	fmt.Println(len(m))

	cycles := escape(m)
	fmt.Println(cycles)
}

func escape(m []int) int {
	clock := 0
	i := 0

	for {
		clock++
		n := i + m[i]
		if m[i] >= 3 {
			m[i] -= 1
		} else {
			m[i] += 1
		}
		i = n
		if i >= len(m) {
			return clock
		}
	}
}
