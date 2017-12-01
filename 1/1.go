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

        // Put initial digit at end.
        bs = append(bs, bs[0])

        fmt.Printf("%v\n", bs)

        fmt.Printf("%v\n", advent_sum(bs))

}

func advent_sum(bs []byte) int {
  s := 0
  for i := 0; i < len(bs)-1; i++ {
    if bs[i] == bs[i+1] {
      s += int(bs[i])
    }
  }
  return s
}
