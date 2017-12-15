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
	Name      string
	Weight    int
	SumWeight int // weight of me plus children
	Children  []string
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
		node := Node{parent, weight, 0, children}
		fmt.Println(node)
		nodes[parent] = node
		for i := 0; i < len(children); i++ {
			child := children[i]
			cs[child] = true
		}

	}

	var root string
	for p, _ := range nodes {
		if !cs[p] {
			root = p
			fmt.Println(p)
		}
	}

	compute_sum_weight(nodes, root)
}

func compute_sum_weight(nodes map[string]Node, root string) {
	node := nodes[root]
	if node.Name == "" {
		log.Fatal("no node")
	}

	weights := make(map[int]bool)
	weight := 0
	for _, child := range node.Children {
		compute_sum_weight(nodes, child)
		cw := nodes[child].SumWeight
		weights[cw] = true
		weight += cw
	}
	node.SumWeight = node.Weight + weight
	nodes[root] = node
	if len(weights) > 1 {
		fmt.Println(node)
		for _, child := range node.Children {
			fmt.Println(" ", child, nodes[child].Weight, nodes[child].SumWeight)
		}
	}
}
