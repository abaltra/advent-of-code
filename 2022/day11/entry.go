package day11

import (
	"fmt"
)

type Operation func(v1 int64) int64
type Test func(v1 int64) bool

type Monkey struct {
	id              int
	items           []int64
	operation       Operation
	test            Test
	flingToIfTrue   int
	flingToIfFalse  int
	inspectionCount int
}

func (m Monkey) ToString() string {
  return fmt.Sprintf("Monkey %d (activityCount: %d): %v\n", m.id, m.inspectionCount, m.items)
}

var STRESS_DIVIDER = 9699690

var MONKEYS = []*Monkey{
	{
		id:             0,
		items:          []int64{84, 66, 62, 69, 88, 91, 91},
		operation:      func(old int64) int64 { return old * 11 },
		test:           func(v int64) bool { return v%2 == 0 },
		flingToIfTrue:  4,
		flingToIfFalse: 7,
		inspectionCount: 0,
	},
	{
		id:             1,
		items:          []int64{98, 50, 76, 99},
		operation:      func(old int64) int64 { return old * old },
		test:           func(v int64) bool { return v%7 == 0 },
		flingToIfTrue:  3,
		flingToIfFalse: 6,
		inspectionCount: 0,
	},
	{
		id:             2,
		items:          []int64{72,56,94},
		operation:      func(old int64) int64 { return old + 1 },
		test:           func(v int64) bool { return v%13 == 0 },
		flingToIfTrue:  4,
		flingToIfFalse: 0,
		inspectionCount: 0,
	},
	{
		id:             3,
		items:          []int64{55,88,90,77,60,67},
		operation:      func(old int64) int64 { return old + 2 },
		test:           func(v int64) bool { return v%3 == 0 },
		flingToIfTrue:  6,
		flingToIfFalse: 5,
		inspectionCount: 0,
	},
	{
		id:             4,
		items:          []int64{69,72,63,60,72,52,63,78},
		operation:      func(old int64) int64 { return old * 13 },
		test:           func(v int64) bool { return v%19 == 0 },
		flingToIfTrue:  1,
		flingToIfFalse: 7,
		inspectionCount: 0,
	},
  {
		id:             5,
		items:          []int64{89,73},
		operation:      func(old int64) int64 { return old + 5 },
		test:           func(v int64) bool { return v%17 == 0 },
		flingToIfTrue:  2,
		flingToIfFalse: 0,
		inspectionCount: 0,
	},
  {
		id:             6,
		items:          []int64{78,68,98,88,66},
		operation:      func(old int64) int64 { return old + 6 },
		test:           func(v int64) bool { return v%11 == 0 },
		flingToIfTrue:  2,
		flingToIfFalse: 5,
		inspectionCount: 0,
	},
  {
		id:             7,
		items:          []int64{70},
		operation:      func(old int64) int64 { return old + 7 },
		test:           func(v int64) bool { return v%5 == 0 },
		flingToIfTrue:  1,
		flingToIfFalse: 3,
		inspectionCount: 0,
	},
}


func Run() {
	println("day 11")

	for i := 0; i < 10000; i++ {
		for _, monkey := range MONKEYS {
      for len(monkey.items) > 0 {
        item := monkey.items[0]
        monkey.items = monkey.items[1:]
        item = monkey.operation(item)
        item %= int64(STRESS_DIVIDER)
        if monkey.test(item) {
          MONKEYS[monkey.flingToIfTrue].items = append(MONKEYS[monkey.flingToIfTrue].items, item)
        } else {
          MONKEYS[monkey.flingToIfFalse].items = append(MONKEYS[monkey.flingToIfFalse].items, item)
        }
        monkey.inspectionCount++
      }
		}
	}

	for _, monkey := range MONKEYS {
	  println(monkey.ToString())
	}
}
