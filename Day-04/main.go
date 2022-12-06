package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	elves := strings.Split(string(bytes), "\n")

	count1, _ := partOne(elves)
	fmt.Printf("count1: %v\n", count1)
	count2, _ := partTwo(elves)
	fmt.Printf("count2: %v\n", count2)
}

func partOne(elves []string) (count int, err error) {
	for _, pair := range elves {
		m, err := createElfMap(pair)
		if err != nil {
			return -1, err
		}
		oneContainsTwo := m["start0"] >= m["start1"] && m["end0"] <= m["end1"]
		twoContainsOne := m["start0"] <= m["start1"] && m["end0"] >= m["end1"]

		if oneContainsTwo || twoContainsOne {
			count++
		}

	}
	return count, nil
}

func partTwo(elves []string) (count int, err error) {
	for _, pair := range elves {
		m, err := createElfMap(pair)
		if err != nil {
			return -1, err
		}

		if m["end0"] >= m["start1"] && m["start0"] <= m["end1"] || m["end0"] <= m["start1"] && m["start0"] >= m["end1"] {
			count++
		}

	}
	return count, nil
}

func createElfMap(pair string) (m map[string]int, err error) {
	m = make(map[string]int)

	for i, elf := range strings.Split(pair, ",") {
		bounds := strings.Split(elf, "-")
		if m[fmt.Sprintf("start%d", i)], err = strconv.Atoi(bounds[0]); err != nil {
			return m, err
		}
		if m[fmt.Sprintf("end%d", i)], err = strconv.Atoi(bounds[1]); err != nil {
			return m, err
		}
	}
	return m, nil
}
