package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Component struct{ a, b int }

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var components []Component

	re := regexp.MustCompile("(\\d+)/(\\d+)")

	lines := bufio.NewScanner(inp)
	for lines.Scan() {
		line := lines.Text()
		match := re.FindStringSubmatch(line)
		// Components of form
		// PortA / PortB
		// 1       2
		a, err := strconv.Atoi(match[1])
		if err != nil {
			log.Fatal("bad port")
		}
		b, err := strconv.Atoi(match[2])
		if err != nil {
			log.Fatal("bad port")
		}
		fmt.Println(a, "/", b)
		component := Component{a, b}
		components = append(components, component)
	}
	fmt.Println(components)
	bridge(0, components)
}

var M int
var L int

// cs are all the components
// i is the number of components that
// have been made into a bridge (the first i of cs).
func bridge(i int, cs []Component) {
	port := 0
	if i > 0 {
		port = cs[i-1].b
	}

	// extend the bridge by scanning forwards to find
	// a component matching the port.
	// When found, swap it with the component at index i.
	found := false
	for j := i; j < len(cs); j++ {
		c := cs[j]
		if c.b == port {
			c.a, c.b = c.b, c.a
			cs[j] = c
		}
		if c.a == port {
			cs[i], cs[j] = cs[j], cs[i]
			bridge(i+1, cs)
			found = true
		}
	}

	// End of bridge?
	if !found {
		s := strength(cs[:i])
		if i > L ||
			(i == L && s > M) {
			M = s
			L = i
			fmt.Println(cs[:i])
		}
		fmt.Println(L, M, s)
	}
}

func strength(cs []Component) int {
	a := 0
	for _, c := range cs {
		a += c.a
		a += c.b
	}
	return a
}
