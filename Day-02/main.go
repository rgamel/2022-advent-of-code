package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const rock = 1
const paper = 2
const scissors = 3

const win = 6
const lose = 0
const draw = 3

func main() {
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}

	rounds := strings.Split(string(bytes[:]), "\n")

	m1 := map[string]int{"A X": rock + draw, "A Y": paper + win, "A Z": scissors + lose, "B X": rock + lose, "B Y": paper + draw, "B Z": scissors + win, "C X": rock + win, "C Y": paper + lose, "C Z": scissors + draw}
	m2 := map[string]int{"A X": lose + scissors, "A Y": draw + rock, "A Z": win + paper, "B X": lose + rock, "B Y": draw + paper, "B Z": win + scissors, "C X": lose + paper, "C Y": draw + scissors, "C Z": win + rock}

	fmt.Printf("ansOne: %v\nansTwo: %v\n", sumRounds(rounds, m1), sumRounds(rounds, m2))
}

func sumRounds(rounds []string, dict map[string]int) (total int) {
	for _, round := range rounds {
		total += dict[round]
	}
	return total
}
