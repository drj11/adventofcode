package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strconv"
)

type Node struct {
	Name     string
	Weight   int
	Children []string
}

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	re := regexp.MustCompile(
		"(\\w+)\\s*\\((\\d+)\\)(?:\\s*->\\s*(.*))?")

	cre := regexp.MustCompile("\\w+")

	nodes := make(map[string]Node)
	cs := make(map[string]bool)

	lines := bufio.NewScanner(inp)
	for lines.Scan() {
		line := lines.Text()
		match := re.FindStringSubmatch(line)
		parent := match[1]
		weight_s := match[2]
		children := cre.FindAllString(match[3], -1)

		weight, err := strconv.Atoi(weight_s)
		if err != nil {
			log.Fatal("not an integer")
		}
		node := Node{parent, weight, children}
		fmt.Println(node)
		nodes[parent] = node
		for i := 0; i < len(children); i++ {
			child := children[i]
			cs[child] = true
		}

	}

	for p, _ := range nodes {
		if !cs[p] {
			fmt.Println(p)
		}
	}
}
