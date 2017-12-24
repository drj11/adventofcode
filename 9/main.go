package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(
		"(\\w+)\\s*(inc|dec)\\s*(-?\\d+)\\s*if\\s*(\\w+)\\s*([<>!=]+)\\s*(-?\\d+)")

	lines := bufio.NewScanner(inp)
	for lines.Scan() {
		line := lines.Text()
		match := re.FindStringSubmatch(line)
		fmt.Println(match)
	}
}
