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
		var fs []float64

		line := lines.Text()
		words := bufio.NewScanner(strings.NewReader(line))
		words.Split(bufio.ScanWords)
		for words.Scan() {
			w := words.Text()
			f, err := strconv.ParseFloat(w, 64)
			if err != nil {
				log.Fatal(err)
			}
			fs = append(fs, f)
		}
		p, q := minmax(fs)
		sum += q - p
		fmt.Println(p, q)
	}
	fmt.Println(sum)
}

func minmax(fs []float64) (float64, float64) {
	min := math.MaxFloat64
	max := -math.MaxFloat64
	for i := 0; i < len(fs); i++ {
		f := fs[i]
		min = math.Min(min, f)
		max = math.Max(max, f)
	}
	return min, max
}
