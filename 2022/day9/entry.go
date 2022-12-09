package day9

import (
	"fmt"
	"strconv"
	"strings"
)

type Position struct {
  x int
  y int
}

func (p *Position) Add (p2 Orientation){
  p.x += p2.x
  p.y += p2.y
}

func (p Position) ToString() string {
  return fmt.Sprintf("%d,%d", p.x, p.y)
}

func (p Position) Equals(target Position) bool {
	return p.x == target.x && p.y == target.y
}

func (p Position) IsAdjacent(target Position) bool {
	return p.Equals(target) ||
					Position{ x: p.x+1, y: p.y}.Equals(target) ||
					Position{ x: p.x-1, y: p.y}.Equals(target) ||
					Position{ x: p.x, y: p.y+1}.Equals(target) ||
					Position{ x: p.x, y: p.y-1}.Equals(target) ||
					Position{ x: p.x+1, y: p.y+1}.Equals(target) ||
					Position{ x: p.x+1, y: p.y-1}.Equals(target) ||	
					Position{ x: p.x-1, y: p.y-1}.Equals(target) ||	
					Position{ x: p.x-1, y: p.y+1}.Equals(target)
}

func (p *Position) Follow(target Position) bool {
  if p.IsAdjacent(target) {
  	return false
  }

	if p.x == target.x {
		// just move over the Y axis
		if p.y < target.y {
			p.y ++
		} else {
			p.y --
		}
	} else if p.y == target.y {
		// or just over Y
		if p.x < target.x {
			p.x ++
		} else {
			p.x --
		}
	} else {
		// they're both different, move in diagonal
		if p.x < target.x {
			p.x ++
		} else {
			p.x --
		}

		if p.y < target.y {
			p.y ++
		} else {
			p.y --
		}
	}


	return true
}

type Orientation Position

func NewOrientation(o string) Orientation {
  if o == "U" {
    return Orientation { x: 0, y: 1 }
  } else if (o == "D") {
    return Orientation{ x: 0, y: -1 }
  } else if (o == "L") {
    return Orientation { x: -1, y: 0 }
  }

  return Orientation { x: 1, y: 0 }
}

type Knot struct {
	position *Position
	next *Knot
}

type Rope struct {
	head *Knot
	tail *Knot
	length int
}

func (r *Rope) Init(length int) {
	for i := 0; i < length; i++ {
		r.AddKnot(&Position{})
	}
}

func (r Rope) ToString() string {
	out := ""

	ptr := r.head
	for ; ptr != nil ; ptr = ptr.next {
		out += fmt.Sprintf("(%s) ->", ptr.position.ToString())
	}

	return out
}

func (r *Rope) AddKnot(p *Position) {
	if r.head == nil {
		r.head = &Knot{
			position: p,
		}

		r.tail = r.head
	} else {
		ptr := r.head
		for ; ptr.next != nil; {
			ptr = ptr.next
		}
		ptr.next = &Knot {
			position: p,
		}

		r.tail = ptr.next
	}
	r.length++
}

func (r *Rope) MoveHead(p Orientation) {
	mvr := r.head
	mvr.position.Add(p)
	for ; mvr != nil; {
		follower := mvr.next
		if follower != nil {
			moved := follower.position.Follow(*mvr.position)
			if !moved {
				return
			}
		}
		
		mvr = mvr.next
	}
}

func Run(lines []string) {
  seenPositions := map[string]bool { "0,0": true }
  seenRopeTailPositions := map[string]bool { "0,0": true }

	shortRope := Rope{}
	shortRope.Init(2)

  rope := Rope{}
  rope.Init(10)

  for _, line := range lines {
    parts := strings.Split(line, " ")
    orientation := NewOrientation(parts[0])
    steps, _ := strconv.Atoi(parts[1])
    for i := 0; i < steps; i ++ {
      shortRope.MoveHead(orientation)
      rope.MoveHead(orientation)
      seenRopeTailPositions[fmt.Sprintf("%s", rope.tail.position.ToString())] = true
      seenPositions[fmt.Sprintf("%s", shortRope.tail.position.ToString())] = true
    } 
  }

  fmt.Printf("We've visited %d different spaces\n", len(seenPositions))
  fmt.Printf("The tail of the rope has visited %d different spaces\n", len(seenRopeTailPositions))
}
