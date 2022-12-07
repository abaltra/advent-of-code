package day6

import (
	"bufio"
	"errors"
	"fmt"
	"os"
)

type LimitQueue struct {
	values []byte
	limit  int
}

func NewQueue(limit int) LimitQueue {
	return LimitQueue{
		limit:  limit,
		values: make([]byte, 0),
	}
}

func (q *LimitQueue) Push(v byte) {
	q.values = append(q.values, v)
	if len(q.values) > q.limit {
		q.values = q.values[1:]
	}
}

func (q *LimitQueue) Pop() (byte, error) {
	if len(q.values) == 0 {
		return byte(0), errors.New("empty queue")
	}

	v := q.values[0]
	q.values = q.values[1:]

	return v, nil
}

func (q LimitQueue) IsSizedSet() bool {
	if len(q.values) != q.limit {
		return false
	}

	collisions := make(map[byte]int)

	for _, value := range q.values {
		collisions[value]++
		if collisions[value] > 1 {
			return false
		}
	}

	return true
}

func Run() {
	println("hello 6")

	input, err := os.Open("./day6/input.txt")

	if err != nil {
		panic(err)
	}

	filescanner := bufio.NewScanner(input)
	filescanner.Split(bufio.ScanBytes)

	q := NewQueue(4)
	qMessage := NewQueue(14)

	queue := make([]byte, 0)
	messageQueue := make([]byte, 0)
	counter := 0

	sopIndex := -1
	somIndex := -1

	for filescanner.Scan() {
		c := filescanner.Bytes()

		queue = append(queue, c[0])
		messageQueue = append(messageQueue, c[0])

		q.Push(c[0])
		qMessage.Push(c[0])

		if len(queue) > 4 {
			queue = queue[1:]
		}

		if len(messageQueue) > 14 {
			messageQueue = messageQueue[1:]
		}

		if q.IsSizedSet() && sopIndex == -1 {
			fmt.Printf("Found SOP marker: %s. Data starts at index %d\n", string(q.values), counter+1)
			sopIndex = counter + 1
		}

		if qMessage.IsSizedSet() && somIndex == -1 {
			fmt.Printf("Found SOM marker: %s. Data starts at index %d\n", string(qMessage.values), counter+1)
			somIndex = counter + 1
		}
		counter++
	}
}

func areUnique(set []byte) bool {
	collisions := make(map[byte]int)

	for _, c := range set {
		collisions[c]++
		if collisions[c] > 1 {
			return false
		}
	}

	return true
}
