package main

import (
	"bufio"
	"os"

	"github.com/abaltra/aoc2022/day1"
	"github.com/abaltra/aoc2022/day10"
	"github.com/abaltra/aoc2022/day12"
	"github.com/abaltra/aoc2022/day2"
	"github.com/abaltra/aoc2022/day3"
	"github.com/abaltra/aoc2022/day4"
	"github.com/abaltra/aoc2022/day5"
	"github.com/abaltra/aoc2022/day6"
	"github.com/abaltra/aoc2022/day7"
	"github.com/abaltra/aoc2022/day8"
	"github.com/abaltra/aoc2022/day9"
)

func readInput(path string) []string {
	input, err := os.Open(path)

	if err != nil {
		panic(err)
	}

	filescanner := bufio.NewScanner(input)
	filescanner.Split(bufio.ScanLines)

	lines := make([]string, 0)
	for filescanner.Scan() {
		l := filescanner.Text()
		lines = append(lines, l)
	}

	input.Close()
	return lines
}

func main() {
	day1.Run(readInput("./day1/input.txt"))
	day2.Run(readInput("./day2/input.txt"))
	day3.Run(readInput("./day3/input.txt"))
	day4.Run(readInput("./day4/input.txt"))
	day5.Run(readInput("./day5/input.txt"))
	day6.Run()
	day7.Run(readInput("./day7/input.txt"))
	day8.Run(readInput("./day8/input.txt"))
	day9.Run(readInput("./day9/input.txt"))
	day10.Run(readInput("./day10/input.txt"))
	day12.Run(readInput("./day12/input.txt"))
}
