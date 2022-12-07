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
	bytes, err := ioutil.ReadFile("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	parts := strings.Split(string(bytes[:]), "\n\n")

	diagram := parts[0]
	moves := parts[1]

	fmt.Printf("Part one: %v\n", partOne(diagram, moves)) // part one: ZBDRNPMVH
	fmt.Printf("Part two: %v\n", partTwo(diagram, moves)) // part two: WDLPFNNNB
}

func partOne(diagram string, moves string) (tops string) {
	stackMap := parseDiagram(diagram)
	parsedMoves := parseMoves(moves)
	applyMovesToGrid(parsedMoves, &stackMap, true)

	return getTopsListFromGrid(stackMap)
}

func partTwo(diagram string, moves string) (tops string) {
	stackMap := parseDiagram(diagram)
	parsedMoves := parseMoves(moves)
	applyMovesToGrid(parsedMoves, &stackMap, false)

	return getTopsListFromGrid(stackMap)
}

type cratesGrid = map[string][]string

func getTopsListFromGrid(stackMap cratesGrid) (tops string) {
	cols := sort.IntSlice{}

	for k := range stackMap {
		if i, err := strconv.Atoi(k); err == nil {
			cols = append(cols, i)
		}
	}
	cols.Sort()

	for _, v := range cols {
		stack := stackMap[fmt.Sprint(v)]
		tops += stack[len(stack)-1]
	}

	return tops
}

func parseDiagram(diagram string) (stackMap cratesGrid) {
	strata := strings.Split(diagram[:], "\n")
	stackKey := strata[len(strata)-1]

	keyIndexMap := make(map[string]int)
	stackMap = make(cratesGrid)

	for i, v := range stackKey {
		if string(v) != " " {
			keyIndexMap[string(v)] = i
		}
	}

	for k := range keyIndexMap {
		stackMap[k] = make([]string, 0)
	}

	strata = (strata[:len(strata)-1])
	for i := len(strata) - 1; i >= 0; i-- {
		for k, v := range keyIndexMap {
			if string(strata[i][v]) != " " {
				stackMap[k] = append(stackMap[k], string(strata[i][v]))
			}
		}
	}
	return stackMap
}

type moveMap = map[string]int

func parseMoves(moveList string) (parsedMoves []moveMap) {
	moves := strings.Split(moveList, "\n")

	for _, move := range moves {
		parsedMoves = append(parsedMoves, parseMove(move))
	}
	return parsedMoves
}

func parseMove(move string) (parsedMove moveMap) {
	s := strings.Split(move, " ")
	ints := []int{}

	for _, v := range s {
		if val, err := strconv.Atoi(v); err == nil {
			ints = append(ints, val)
		}
	}

	return moveMap{"count": ints[0], "from": ints[1], "to": ints[2]}
}

func applyMovesToGrid(moves []moveMap, stacks *cratesGrid, reverse bool) {
	for _, move := range moves {
		applyMoveToGrid(move, stacks, reverse)
	}
}

func applyMoveToGrid(move moveMap, stacks *cratesGrid, reverse bool) {
	s := *stacks
	crates, fromCol := popN(s[fmt.Sprint(move["from"])], move["count"])

	if reverse {
		for i, j := 0, len(crates)-1; i < j; i, j = i+1, j-1 {
			crates[i], crates[j] = crates[j], crates[i]
		}
	}

	s[fmt.Sprint(move["from"])] = fromCol
	s[fmt.Sprint(move["to"])] = append(s[fmt.Sprint(move["to"])], crates...)

	stacks = &s
}

func popN(a []string, n int) (x []string, arr []string) {
	return a[len(a)-n:], a[:len(a)-n]
}
