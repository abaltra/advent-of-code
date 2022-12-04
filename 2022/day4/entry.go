package day4

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Range struct {
	start int
	end   int
}

func (r Range) contains(value int) bool {
	return value >= r.start && value <= r.end
}

func Run() {
	input, err := os.Open("./day4/input.txt")

	if err != nil {
		panic(err)
	}

	filescanner := bufio.NewScanner(input)
	filescanner.Split(bufio.ScanLines)

	overlaps := 0
	partialOverlaps := 0
	for filescanner.Scan() {
		l := filescanner.Text()
		parts := strings.Split(l, ",")
		elf1 := parts[0]
		elf2 := parts[1]

		overlaps += isContained(elf1, elf2)
		partialOverlaps += isPartiallyContained(elf1, elf2)
	}

	fmt.Printf("There are %d fully overlapping sets\n", overlaps)
	fmt.Printf("There are %d partially overlapping sets\n", partialOverlaps)
}

func isPartiallyContained(elf1, elf2 string) int {
	elf1RangeLimits := strings.Split(elf1, "-")
	elf2RangeLimits := strings.Split(elf2, "-")

	range1 := Range{}
	range2 := Range{}

	range1.start, _ = strconv.Atoi(elf1RangeLimits[0])
	range1.end, _ = strconv.Atoi((elf1RangeLimits[1]))

	range2.start, _ = strconv.Atoi(elf2RangeLimits[0])
	range2.end, _ = strconv.Atoi((elf2RangeLimits[1]))

	if range1.contains(range2.start) || range1.contains(range2.end) {
		return 1
	}

	if range2.contains(range1.start) || range2.contains(range1.end) {
		return 1
	}

	return 0
}

func isContained(elf1, elf2 string) int {
	elf1RangeLimits := strings.Split(elf1, "-")
	elf2RangeLimits := strings.Split(elf2, "-")

	range1 := Range{}
	range2 := Range{}

	range1.start, _ = strconv.Atoi(elf1RangeLimits[0])
	range1.end, _ = strconv.Atoi((elf1RangeLimits[1]))

	range2.start, _ = strconv.Atoi(elf2RangeLimits[0])
	range2.end, _ = strconv.Atoi((elf2RangeLimits[1]))

	if range1.contains(range2.start) && range1.contains(range2.end) {
		return 1
	}

	if range2.contains(range1.start) && range2.contains(range1.end) {
		// range 2 contained in range 1
		return 1
	}

	return 0
}
