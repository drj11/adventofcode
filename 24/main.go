package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

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
	}
}
