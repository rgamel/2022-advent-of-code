package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	packs := strings.Split(string(bytes[:]), "\n")

	sum1 := partOne(packs)
	fmt.Printf("sum1: %v\n", sum1)

	sum2 := partTwo(packs)
	fmt.Printf("sum2: %v\n", sum2)

}

func partOne(packs []string) (result int) {
	for _, pack := range packs {
		sackVal := getPriority(findInBothCompartments(partitionString(getSplitPoint(pack))))
		result += sackVal
	}
	return result
}

func partTwo(packs []string) (result int) {
	for i := 0; i < len(packs); i += 3 {
		result += getPriority(findInThreeBags(packs[i : i+3]))
	}
	return result
}

func getSplitPoint(str string) (s string, p int) {
	p = len(str) / 2
	return str, p
}

func partitionString(str string, i int) (partA, partB string) {
	partA = str[:i]
	partB = str[i:]
	return partA, partB
}

func findInBothCompartments(str1, str2 string) (r rune) {
	m := make(map[rune]int)

	for _, c := range str1 {
		m[c] = 0
	}

	for _, c := range str2 {
		if _, ok := m[c]; ok {
			return c
		}
	}
	return r
}

func findInThreeBags(bags []string) (r rune) {
	m := make(map[rune]int)

	for _, c := range bags[0] {
		m[c] = 0
	}

	for i, bag := range bags[1:] {
		for _, c := range bag {
			if val, ok := m[c]; ok {
				if val == i {
					m[c] = i + 1
				}
				if m[c] == 2 {
					return c
				}
			}
		}
	}

	return r
}

func getPriority(r rune) (p int) {
	switch {
	case 122 >= r && r >= 97:
		return int(r) - 96
	case r >= 65:
		return int(r) - 38
	default:
		log.Fatal("Invalid rune value")
	}
	return
}
