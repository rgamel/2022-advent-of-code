package main

import (
	"fmt"
	"io/ioutil"
	"log"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	input := string(bytes)
	one, two := findUniqueMarker(input, 4), findUniqueMarker(input, 14)
	fmt.Printf("one: %v\ntwo: %v\n", one, two) // 1912 | 2122
}

func findUniqueMarker(input string, size int) (index int) {
	if len(input) >= size {
		for i := size; i <= len(input); i++ {
			if isNonRepeating(input[i-size : i]) {
				return i
			}
		}
	}
	return -1
}

func isNonRepeating(str string) bool {
	m := make(map[rune]int)
	for _, b := range str {
		if _, ok := m[b]; ok {
			return false
		}
		m[b] = 0
	}
	return true
}
