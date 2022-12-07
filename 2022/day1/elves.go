package day1

import (
	"container/heap"
	"fmt"
	"strconv"
)

type Elf struct {
	index     int
	calories  int
	itemCount int
}

type ElfHeap []*Elf

func (h ElfHeap) Len() int           { return len(h) }
func (h ElfHeap) Less(i, j int) bool { return h[i].calories < h[j].calories }
func (h ElfHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *ElfHeap) Push(x any) {
	elf := x.(*Elf)
	*h = append(*h, elf)
}

func (h *ElfHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	old[n-1] = nil
	*h = old[0 : n-1]
	return x
}

func Run(lines []string) {

	highest_calorie_count := -1
	elf_index := 0
	number_of_items := 0
	elf_heap := ElfHeap{}
	heap.Init(&elf_heap)

	running_count := 0
	for _, l := range lines {
		if l == "" {
			if running_count > highest_calorie_count {
				highest_calorie_count = running_count
			}

			e := &Elf{
				index:     elf_index,
				calories:  running_count,
				itemCount: number_of_items,
			}
			heap.Push(&elf_heap, e)
			if elf_heap.Len() > 3 {
				heap.Pop(&elf_heap)
			}

			running_count = 0
			elf_index += 1
		}

		v, _ := strconv.Atoi(l)
		running_count += v
		number_of_items += 1
	}

	fmt.Printf("Highest calorie count: %d\n", highest_calorie_count)

	top_carried := 0
	for elf_heap.Len() > 0 {
		elf := heap.Pop(&elf_heap).(*Elf)
		fmt.Printf("Elf: index -> %v, calories -> %v, items -> %v\n", elf.index, elf.calories, elf.itemCount)
		top_carried += elf.calories
	}
	fmt.Printf("Total carried by top 3: %v\n", top_carried)
}
