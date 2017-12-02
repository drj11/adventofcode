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
		p, q := find_pair(fs)
		sum += q - p
		fmt.Println(p, q)
	}
	fmt.Println(sum)
}

func find_pair(fs []float64) (float64, float64) {
	for i := 0; i < len(fs); i++ {
		for j := 0; j < len(fs); j++ {
			a := fs[i] // dividend
			b := fs[j] // divisor
			if a == b {
				continue
			}
			if math.Mod(a, b) == 0.0 {
				return b, a
			}
		}
	}
	log.Fatal("halp")
	return math.NaN(), math.NaN()
}
