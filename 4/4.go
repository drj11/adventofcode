package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	n_valid := 0

	lines := bufio.NewScanner(inp)
	for lines.Scan() {

		line := lines.Text()
		words := bufio.NewScanner(strings.NewReader(line))
		words.Split(bufio.ScanWords)
		n_valid += passphrase(words)

	}
	fmt.Println(n_valid)
}

func passphrase(words *bufio.Scanner) int {
	m := make(map[string]bool)
	for words.Scan() {
		w := words.Text()
		if m[w] {
			return 0
		}
		m[w] = true
	}
	return 1
}
