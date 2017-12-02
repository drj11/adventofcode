package main

import (
        "bufio"
	"fmt"
        "os"
	"log"
)

func main() {
	inp, err := os.Open("input")
	if err != nil {
		log.Fatal(err)
	}

        scanner := bufio.NewScanner(inp)
        for scanner.Scan() {
          fmt.Println(scanner.Text())
        }
}
