package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bs, err := ioutil.ReadFile("input")
	if err != nil {
		log.Fatal(err)
	}

        // Trim away trailing non-digits.
	l := len(bs)
	for !('0' <= bs[l-1] && bs[l-1] <= '9') {
		l = l - 1
	}

        bs = bs[:l]
	fmt.Printf("%v\n", len(bs))

        // Convert each digit to its numerical value
        for i := 0; i < l; i++ {
          bs[i] -= '0'
        }

        // Copy entire buffer onto end.
        bs = append(bs, bs...)

        fmt.Printf("%v\n", bs)

        fmt.Printf("%v\n", advent_sum(bs, l, l/2))

}

func advent_sum(bs []byte, l, o int) int {
  s := 0
  for i := 0; i < l; i++ {
    if bs[i] == bs[i+o] {
      s += int(bs[i])
    }
  }
  return s
}
