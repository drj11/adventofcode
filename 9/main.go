package main

import (
	"bufio"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(
		"(\\w+)\\s*(inc|dec)\\s*(-?\\d+)\\s*if\\s*(\\w+)\\s*([<>!=]+)\\s*(-?\\d+)")
	register := make(map[string]float64)
	var max float64

	lines := bufio.NewScanner(inp)
	for lines.Scan() {
		line := lines.Text()
		match := re.FindStringSubmatch(line)
		fmt.Println(match)
		// Instructions of form:
		// T Op V if R Cmp C
		// 1 2  3    4 5   6
		R := match[4]
		v := register[R]
		c, err := strconv.ParseFloat(match[6], 64)
		if err != nil {
			log.Fatal("Threshold bad")
		}
		Cmp := match[5]
		var modify bool
		switch Cmp {
		case "<":
			modify = v < c
		case ">":
			modify = v > c
		case "<=":
			modify = v <= c
		case ">=":
			modify = v >= c
		case "==":
			modify = v == c
		case "!=":
			modify = v != c
		default:
			log.Fatal("Comparison bad")
		}
		if modify {
			T := match[1]
			addend, err := strconv.ParseFloat(match[3], 64)
			if err != nil {
				log.Fatal("Addend bad")
			}
			if match[2] == "dec" {
				addend = -addend
			}
			register[T] += addend
			max = math.Max(max, register[T])
		}
	}
	fmt.Println(register)
	fmt.Println(max)
}
