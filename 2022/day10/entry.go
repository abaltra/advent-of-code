package day10

import (
	"bytes"
	"fmt"
	"strconv"
	"strings"
)

type OPFunc func(a int) int

type Operation struct {
	raw string
	OPTYPE     string
	cyclesLeft int
	function   OPFunc
	operand    int
}

func GetOpFromLine(line string) Operation {
	parts := strings.Split(line, " ")
	ret := Operation{OPTYPE: parts[0], raw: line}
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

func PrintScreen(screen [][]byte, height int) {
	for i := 0; i < height ; i++ {
		println(string(screen[i]))
	}
	println()
}

func spriteString(xPos, sWidth int) []byte{

	str := bytes.Repeat([]byte("."), sWidth)
	if xPos >= 0 && xPos < sWidth {
		str[xPos] = byte('#')
	}

	if xPos - 1 >= 0 && xPos - 1 < sWidth {
		str[xPos - 1] = byte('#')
	}

	if xPos + 1 >= 0 && xPos + 1 < sWidth {
		str[xPos + 1] = byte('#')
	}

	return str
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

	SCREEN_WIDTH := 40
	SCREEN_HEIGHT := 6
	screen := make([][]byte, SCREEN_HEIGHT)
	for i := 0; i < SCREEN_HEIGHT; i++ {
		screen[i] = bytes.Repeat([]byte("."), SCREEN_WIDTH)
	}

	for {
		cycle += 1

		// START CYCLE
		// START NEW OP IF WE'RE IDLE
		if currentOp.cyclesLeft == 0 {
			if lIndex >= len(lines) {
				break
			}
			currentOp = GetOpFromLine(lines[lIndex])
			lIndex++
		}

		// DURING CYCLE
		if cycle == meaningfulCycle {
			meaningfulCycle += 40
			signalStrengthAggregate += cycle * signal
		}

		spriteString := spriteString(signal, SCREEN_WIDTH)
		if spriteString[(cycle - 1) % SCREEN_WIDTH] == byte('#') {
			screen[(cycle - 1) / 40][(cycle - 1) % SCREEN_WIDTH] = byte('#')
		}


		// END CYCLE
		if currentOp.cyclesLeft == 1 {
			signal = currentOp.function(signal)
		}

		currentOp.cyclesLeft--
	}
	fmt.Printf("The aggregate of meaningful signal strengths is: %d\n", signalStrengthAggregate)
	PrintScreen(screen, SCREEN_HEIGHT)
}
