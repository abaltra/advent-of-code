package day10

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type OPFunc func(a int) int

type Operation struct {
	OPTYPE     string
	cyclesLeft int
	function   OPFunc
	operand    int
}

func GetOpFromLine(line string) Operation {
	parts := strings.Split(line, " ")
	ret := Operation{OPTYPE: parts[0]}
	if ret.OPTYPE == "noop" {
		ret.cyclesLeft = 1
		ret.function = func(a int) int { return a }
	} else {
		ret.cyclesLeft = 2
		ret.operand, _ = strconv.Atoi(parts[1])
		ret.function = func(a int) int { return a + ret.operand }
	}

	return ret
}

func PrintScreen(screen []byte) {
	for i := 0; i < 6; i++ {
		println(string(screen[40*i : 40*i+40]))
	}
	println()
}

func Run(lines []string) {
	println("day 10")

	cycle := 0
	signal := 1

	lIndex := 0
	currentOp := GetOpFromLine("noop")
	currentOp.cyclesLeft = 0
	signalStrengthAggregate := 0
	meaningfulCycle := 20

	// SCREEN_WIDTH := 40
	// SCREEN_HEIGHT := 6
	screen := bytes.Repeat([]byte("."), 240)
	// currentScreenLine := 0

	for {
		cycle += 1
		// fmt.Printf("Signal at start of cycle %d: %d\n", cycle, signal)
		if currentOp.cyclesLeft == 0 {
			if lIndex >= len(lines) {
				break
			}
			currentOp = GetOpFromLine(lines[lIndex])
			lIndex++
		}

		if cycle == meaningfulCycle {
			fmt.Printf("Meaningfule cycle %d. Signal strength: %d\n", cycle, cycle*signal)
			meaningfulCycle += 40
			signalStrengthAggregate += cycle * signal
			break
		}

		if signal-1 <= cycle && signal+1 >= cycle {
			fmt.Printf("Drawing in cycle %d because signal is %d\n", cycle, signal)
			screen[cycle-1] = byte('#')
		}

		if currentOp.cyclesLeft == 1 {
			signal = currentOp.function(signal)
		}

		currentOp.cyclesLeft--
		// fmt.Printf("Signal at end of cycle %d: %d\n", cycle, signal)
	}
	fmt.Printf("The aggregate of meaningful signal strengths is: %d\n", signalStrengthAggregate)
	PrintScreen(screen)
}
