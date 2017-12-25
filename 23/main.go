package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Instruction struct {
	Op      string
	Target  string
	Source  string
	Literal float64
}

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile("(\\w+)\\s*(\\w+)\\s*(\\w+|-?\\d+)")

	program := make([]Instruction, 0)
	lines := bufio.NewScanner(inp)
	for lines.Scan() {
		line := lines.Text()
		match := re.FindStringSubmatch(line)
		fmt.Println(match)

		var i Instruction

		// Instructions of form:
		// Op T S
		// 1  2 3
		i.Op = match[1]
		i.Target = match[2]

		v, err := strconv.ParseFloat(match[3], 64)
		if err != nil {
			i.Source = match[3]
		} else {
			i.Literal = v
		}

		program = append(program, i)
	}
	fmt.Println(program)
}
