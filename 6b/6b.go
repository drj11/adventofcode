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

	var m [16]int
	i := 0

	words := bufio.NewScanner(inp)
	words.Split(bufio.ScanWords)
	for words.Scan() {
		word := words.Text()
		w, err := strconv.Atoi(word)
		if err != nil {
			log.Fatal(err)
		}
		m[i] = w
		i++
	}

	fmt.Println(m)

	seen := make(map[[16]int]int)
	cycles := 0

	for {
		m = redistribute(m)
		fmt.Println(m)
		cycles++
		if seen[m] > 0 {
			fmt.Println(cycles - seen[m])
			return
		}
		seen[m] = cycles
	}
}

func redistribute(m [16]int) [16]int {
	s := 0
	max := 0
	for i := 0; i < 16; i++ {
		if m[i] > max {
			s = i
			max = m[i]
		}
	}

	b := max // amount to redistribute
	m[s] = 0
	i := (s + 1) % 16 // index to start at
	for b > 0 {
		m[i] += 1
		b -= 1
		i = (i + 1) % 16
	}
	return m
}
