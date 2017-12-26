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
	Op      string  // Operation
	Target  int     // Register on Left
	Source  int     // Register on Right or -ve
	Literal float64 // Literal on Right
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
		i.Target = int(match[2][0] - 'a')

		v, err := strconv.ParseFloat(match[3], 64)
		if err != nil {
			i.Source = int(match[3][0] - 'a')
		} else {
			i.Source = -1
			i.Literal = v
		}

		program = append(program, i)
	}
	fmt.Println(program)
	run(program)
}

func run(P []Instruction) {
	var R [8]float64
	i := 0
	muls := 0
	clock := 0

	for i < len(P) {
		I := P[i]
		switch I.Op {
		case "set":
			R[I.Target] = val(R, I)
		case "sub":
			R[I.Target] -= val(R, I)
		case "mul":
			R[I.Target] *= val(R, I)
			muls += 1
		case "jnz":
			// Special case for when left-hand
			// operand is "1" (an unconditional
			// jump)
			if 7 < I.Target ||
				R[I.Target] != 0.0 {
				i += int(val(R, I)) - 1
			}
		}
		i += 1
		clock += 1
		if clock%1000 == 0 {
			fmt.Println(i)
			fmt.Println(R)
		}
	}
	fmt.Println(muls)
}

func val(R [8]float64, I Instruction) float64 {
	if I.Source < 0 {
		return I.Literal
	}
	return R[I.Source]
}
