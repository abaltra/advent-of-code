	package day6

	import (
		"bufio"
		"fmt"
		"os"
	)

	func Run() {
  	println("hello 6")

  	
  	input, err := os.Open("./day6/input.txt")

		if err != nil {
		panic(err)
	}

	filescanner := bufio.NewScanner(input)
	filescanner.Split(bufio.ScanBytes)


	queue := make([]byte, 0)
	messageQueue := make([]byte, 0)
	counter := 0

	sopIndex := -1
	somIndex := -1

	for filescanner.Scan() {
		c := filescanner.Bytes()

		queue = append(queue, c[0])
		messageQueue = append(messageQueue, c[0])

		if len(queue) > 4 {
			temp := make([]byte, 4)
			copy(temp, queue[1:])
			queue = temp
		}

		if len(messageQueue) > 14 {
			temp := make([]byte, 14)
			copy(temp, messageQueue[1:])
			messageQueue = temp
		}

		if len(queue) == 4 && sopIndex == -1 {
			if areUnique(queue) {
				fmt.Printf("Found SOP marker: %s. Data starts at index %d\n", string(queue), counter + 1)
				sopIndex = counter + 1
			}
		}

		if len(messageQueue) == 14 && somIndex == -1 {
			if areUnique(messageQueue) {
				fmt.Printf("Found SOM marker: %s. Data starts at index %d\n", string(messageQueue), counter + 1)
				somIndex = counter + 1
			}
		}
  	counter ++
  }
}

func areUnique(set []byte) bool {
	collisions := make(map[byte]int)

	for _, c := range set {
		collisions[c] ++
		if collisions[c] > 1 {
			return false
		}
	}

	return true
}
