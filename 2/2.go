package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	sum := 0.0

	lines := bufio.NewScanner(inp)
	for lines.Scan() {
		min := math.MaxFloat64
		max := -math.MaxFloat64

		line := lines.Text()
		words := bufio.NewScanner(strings.NewReader(line))
		words.Split(bufio.ScanWords)
		for words.Scan() {
			w := words.Text()
			f, err := strconv.ParseFloat(w, 64)
			if err != nil {
				log.Fatal(err)
			}
			min = math.Min(min, f)
			max = math.Max(max, f)
		}
		sum += max - min
		fmt.Println(min, max)
	}
	fmt.Println(sum)
}
