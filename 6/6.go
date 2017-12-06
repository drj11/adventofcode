package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

	var m []int

	words := bufio.NewScanner(inp)
	words.Split(bufio.ScanWords)
	for words.Scan() {
		word := words.Text()
		i, err := strconv.Atoi(word)
		if err != nil {
			log.Fatal(err)
		}
		m = append(m, i)
	}

	fmt.Println(m)
}
