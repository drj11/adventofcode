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
	b, c := run(program)
	composites := 0
	for ; b <= c; b += 17 {
		composites += is_composite(b)
	}
	fmt.Println(composites)
}

// Return the number of composites in the set{b}
func is_composite(b float64) int {
	a := int(b)
	if a < 3 {
		log.Fatal("too small")
	}
	if a%2 == 0 {
		return 1
	}
	for f := 3; f*f <= a; f += 2 {
		if (a/f)*f == a {
			return 1
		}
	}
	return 0
}

func run(P []Instruction) (float64, float64) {
	var R [8]float64
	i := 0
	muls := 0
	clock := 0
	// For Part Two
	R[0] = 1

	for i < len(P) {
		I := P[i]
		switch I.Op {
		case "set":
			R[I.Target] = val(R, I)
			// For Part Two
			// When F is set,
			// return with B and C, the loop limits.
			if I.Target == 5 {
				return R[1], R[2]
			}
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
			fmt.Println("PC =", i, R)
		}
	}
	fmt.Println(muls)
	return -1, -1
}

func val(R [8]float64, I Instruction) float64 {
	if I.Source < 0 {
		return I.Literal
	}
	return R[I.Source]
}
