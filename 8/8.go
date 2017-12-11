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
		"(\\w+)\\s*\\((\\d+)\\)(?:\\s*->\\s*(.*))?")

        cre := regexp.MustCompile("\\w+")

	lines := bufio.NewScanner(inp)
	for lines.Scan() {
		line := lines.Text()
		match := re.FindStringSubmatch(line)
                parent := match[1]
                weight := match[2]
                children := cre.FindAllString(match[3], -1)
                fmt.Println(parent, weight, children)
	}
}
