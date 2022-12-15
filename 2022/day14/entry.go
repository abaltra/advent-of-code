package day14

import (
	"bytes"
	"fmt"
	"math"
	"strconv"
	"strings"
	//"time"
)

type Position struct {
	x int
	y int
}

type Automata struct {
	position          *Position
	moved             bool
	numberOfMovements int
}

func IsAvailable(grid [][]byte, p Position) bool {
	return p.y < 0 || p.y >= len(grid) || p.x < 0 || p.x >= len(grid[0]) || grid[p.y][p.x] == byte('.')
}

func Run(lines []string) {
	println("dat14")
	largestX := 0
	largestY := 0
	structures := make([][]*Position, len(lines))
	for i, line := range lines {
		positions := strings.Split(line, "->")
		structure := []*Position{}
		for _, position := range positions {
			position = strings.TrimSpace(position)
			indices := strings.Split(position, ",")
			x, _ := strconv.Atoi(indices[0])
			y, _ := strconv.Atoi(indices[1])

			largestY = int(math.Max(float64(y), float64(largestY)))
			largestX = int(math.Max(float64(x), float64(largestX)))

			p := &Position{x: x, y: y}
			structure = append(structure, p)
		}
		structures[i] = structure
	}

	fmt.Printf("largest x: %d, largest y: %d\n", largestX, largestY)

	for _, structure := range structures {
		l := ""
		for _, position := range structure {
			l += fmt.Sprintf("(%d, %d) -> ", position.x, position.y)
		}
		println(l)
	}

	grid := make([][]byte, largestY + 3)
	for i := range grid {
		grid[i] = bytes.Repeat([]byte("."), 1000)
	}
	grid[largestY + 2] = bytes.Repeat([]byte("#"), 1000)

	//printGrid(grid)
	sourcePosition := Position{x: 500, y: 0}

	grid[sourcePosition.y][sourcePosition.x] = byte('+')
	for _, s := range structures {
		drawStructure(grid, s)
	}

	//padEdges(grid)

	//printGrid(grid)
	//println(len(grid))

	automata := []*Automata{}
	for {
		automaton := &Automata{position: &Position{x: sourcePosition.x, y: sourcePosition.y}, moved: true}
		automata = append(automata, automaton)
		for automaton.moved {
			automaton.moved = false

			newPosDown := &Position{x: automaton.position.x, y: automaton.position.y + 1}
			newPosLeft := &Position{x: automaton.position.x - 1, y: automaton.position.y + 1}
			newPosRight := &Position{x: automaton.position.x + 1, y: automaton.position.y + 1}

			var newPos *Position = nil

			if IsAvailable(grid, *newPosDown) {
				newPos = newPosDown
			} else if IsAvailable(grid, *newPosLeft) {
				newPos = newPosLeft
			} else if (IsAvailable(grid, *newPosRight)) {
				newPos = newPosRight
			}

			if newPos != nil {
				if IsOutOfBounds(grid, newPos) {
					//fmt.Printf("Grid is %d wide and %d deep. New pos is %+v\n", len(grid), len(grid[0]), newPos)
					//println("dobne")
					goto END
				} else {
					automaton.position = newPos
					automaton.moved = true
				}
			} else {
				automaton.moved = false
				if automaton.position.x == sourcePosition.x && automaton.position.y == sourcePosition.y {
					println("done")
					goto END
				}
				grid[automaton.position.y][automaton.position.x] = byte('o')
			}

			//time.Sleep(10 * time.Millisecond)
			//printGridWithAutomatas(grid, automata)
			 
			//if len(automata) == 100 {
			//	println("bye")
			//	goto END
			//}
		}
	}

	END: {
		//printGrid(grid)
		printGridWithAutomatas(grid, automata)
		fmt.Printf("Created %d automaton before falling into the void\n", len(automata))
		//printGrid(grid)
	}
}

func IsOutOfBounds(grid [][]byte, p *Position) bool {
	return p.y < 0 || p.y >= len(grid) || p.x < 0 || p.x >= len(grid[0])
}

func printGridWithAutomatas(grid [][]byte, automata []*Automata) {
	automataPositions := map[string]bool{}

	for _, a := range automata {
		automataPositions[fmt.Sprintf("%d,%d", a.position.x, a.position.y)] = true
	}

	for y, row := range grid {
		l := ""
		for x, cell := range row {
			if automataPositions[fmt.Sprintf("%d,%d", x, y)] {
				l += "o"
			} else {
				l += string(cell)
			}
		}
		println(l)
	}

	println()
}

func padEdges(grid [][]byte) {
	for i, row := range grid {
		grid[i] = append([]byte("~"), row...)
		grid[i] = append(grid[i], byte('~'))
	}
	grid[len(grid)-1] = bytes.Repeat([]byte("~"), len(grid[0]))
}

func printGrid(grid [][]byte) {
	for _, row := range grid {
		println(string(row))
	}
	println()
}

func drawStructure(grid [][]byte, structure []*Position) {
	previousPos := structure[0]
	grid[previousPos.y][previousPos.x] = byte('#')
	for _, pos := range structure[1:] {
		grid[pos.y][pos.x] = byte('#')
		if previousPos.x == pos.x {
			// we're moving on Y
			for i := math.Min(float64(previousPos.y), float64(pos.y)); i < math.Max(float64(previousPos.y), float64(pos.y)); i++ {
				grid[int(i)][pos.x] = byte('#')
			}
		} else {
			// we're moving on X
			for i := math.Min(float64(previousPos.x), float64(pos.x)); i < math.Max(float64(previousPos.x), float64(pos.x)); i++ {
				grid[pos.y][int(i)] = byte('#')
			}
		}
		previousPos = pos
	}
}
