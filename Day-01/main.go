package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

func main() {
	elves := []int{}

	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	for _, elf := range strings.Split(string(bytes[:]), "\n\n") {
		elfTotal := 0
		err := sumElf(elf, &elfTotal)
		if err != nil {
			log.Fatal(err)
		}
		elves = append(elves, elfTotal)
	}

	sort.Ints(elves)
	fmt.Printf("1a: %d\n", elves[len(elves)-1]) // get last elf

	answer := sum(elves[len(elves)-3:]) // sum last 3 elves
	fmt.Printf("1b: %d\n", answer)
}

func sumElf(elf string, elfTotal *int) (err error) {
	for _, cal := range strings.Split(elf, "\n") {
		if calInt, err := strconv.Atoi(cal); err != nil {
			return err
		} else {
			*elfTotal += calInt
		}
	}
	return nil
}

func sum(ints []int) (sum int) {
	for _, v := range ints {
		sum += v
	}
	return sum
}
