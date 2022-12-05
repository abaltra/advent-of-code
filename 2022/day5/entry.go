package day5

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func Run() {
	input, err := os.Open("./day5/input.txt")

	stacks := make(map[int][]byte)
	newStacks := make(map[int][]byte)

	if err != nil {
		panic(err)
	}

	filescanner := bufio.NewScanner(input)
	filescanner.Split(bufio.ScanLines)

	processingMoves := false
	maxCrates := 0
	for filescanner.Scan() {
		l := filescanner.Text()

		if l == "" {
			processingMoves = true
			continue
		}

		if l[1] == '1' {
			t := strings.TrimSpace(l)
			maxCrates, _ = strconv.Atoi(t[len(t)-1:])
			continue
		}

		if processingMoves {
			processMoves(stacks, l)
			processNewMoves(newStacks, l)
		} else {
			processStacks(stacks, l)
			processStacks(newStacks, l)
		}
	}

	result := getCrates(stacks, maxCrates)

	printStacks(stacks)
	println(result)
	printStacks(newStacks)
	println(getCrates(newStacks, maxCrates))
}

func getCrates(stacks map[int][]byte, maxCrates int) string {
	ret := make([]byte, 0)
	for i := 1; i <= maxCrates; i++ {
		if len(stacks[i]) > 0 {
			ret = append(ret, stacks[i][0])
		}
	}

	return string(ret)
}

func printStacks(stacks map[int][]byte) {
	for stackId, crates := range stacks {
		fmt.Printf("Stack %d has these crates: %s\n", stackId, string(crates))
	}
}

func processStacks(stacks map[int][]byte, line string) {
	crateId := 0
	for i := 0; i < len(line); i += 4 {
		crateId++
		crate := line[i : i+3]

		if strings.TrimSpace(crate) == "" {
			continue
		}

		id := crate[1]
		if stacks[crateId] == nil {
			stacks[crateId] = make([]byte, 0)
		}

		stacks[crateId] = append(stacks[crateId], id)
	}
}

func processNewMoves(stacks map[int][]byte, line string) {
	parts := strings.Split(line, " ")
	amountS, fromS, toS := parts[1], parts[3], parts[5]

	amount, _ := strconv.Atoi(amountS)
	from, _ := strconv.Atoi(fromS)
	to, _ := strconv.Atoi(toS)

	//fmt.Printf("Moving %d elements from stack %d (%s) to stack %d (%s)\n", amount, from, string(stacks[from]), to, string(stacks[to]))

	e := make([]byte, amount)
	copy(e, stacks[from][0:amount])
	if amount == len(stacks[from]) {
		stacks[from] = []byte{}
	} else {
		stacks[from] = stacks[from][amount:]
	}

	stacks[to] = append(e, stacks[to]...)
	//fmt.Printf("Resulting stack from (%d): %s, to (%d): %s\n", from, string(stacks[from]), to, string(stacks[to]))
}

func processMoves(stacks map[int][]byte, line string) {
	parts := strings.Split(line, " ")
	amountS, fromS, toS := parts[1], parts[3], parts[5]

	amount, _ := strconv.Atoi(amountS)
	from, _ := strconv.Atoi(fromS)
	to, _ := strconv.Atoi(toS)

	for i := 0; i < amount; i++ {
		e := stacks[from][0]
		stacks[from] = stacks[from][1:]

		stacks[to] = append([]byte{e}, stacks[to]...)
	}
}
