package day3

import (
	"bufio"
	"fmt"
	"os"
)

func findDuplicates(set1 string, set2 string) []byte {
	collisionMap := make(map[byte]int)
	foundCollisions := make(map[byte]int)
	collisions := make([]byte, 0)

	for i := 0; i < len(set1); i++ {
		collisionMap[set1[i]] = 1
	}

	for i := 0; i < len(set2); i++ {
		if collisionMap[set2[i]] == 1 && foundCollisions[set2[i]] == 0 {
			collisions = append(collisions, set2[i])
			foundCollisions[set2[i]] = 1
		}
	}

	return collisions
}

func Run() {

	input, err := os.Open("./day3/input.txt")

	if err != nil {
		panic(err)
	}

	filescanner := bufio.NewScanner(input)
	filescanner.Split(bufio.ScanLines)
	score := 0
	groupScore := 0

	rucksackGroups := make([]string, 0)
	for filescanner.Scan() {
		rucksack := filescanner.Text()
		set1 := rucksack[0 : len(rucksack)/2]
		set2 := rucksack[len(rucksack)/2:]

		duplicates := findDuplicates(set1, set2)
		score += calculateScore(duplicates)

		rucksackGroups = append(rucksackGroups, rucksack)
		if len(rucksackGroups) == 3 {
			d := findDuplicatesIngGroup(rucksackGroups[0], rucksackGroups[1], rucksackGroups[2])
			groupScore += calculateScore(d)
			rucksackGroups = make([]string, 0)
		}
	}
	fmt.Printf("Aggregate score: %d\n", score)
	fmt.Printf("Aggregate score for groups: %d\n", groupScore)
}

func removeDuplicates(s string) string {
	collMap := make(map[byte]int)
	uniq := make([]byte, 0)

	for i := 0; i < len(s); i++ {
		collMap[s[i]]++
	}

	for key := range collMap {
		uniq = append(uniq, key)
	}

	return string(uniq)
}

func findDuplicatesIngGroup(g1, g2, g3 string) []byte {
	collisions := make([]byte, 0)
	collMap := make(map[byte]int)

	g1_unique := removeDuplicates(g1)
	g2_unique := removeDuplicates(g2)
	g3_unique := removeDuplicates(g3)

	for i := 0; i < len(g1_unique); i++ {
		collMap[g1_unique[i]]++
	}

	for i := 0; i < len(g2_unique); i++ {
		collMap[g2_unique[i]]++
	}

	for i := 0; i < len(g3_unique); i++ {
		collMap[g3_unique[i]]++

		if collMap[g3_unique[i]] == 3 {
			collisions = append(collisions, g3_unique[i])
		}
	}

	return collisions
}

func calculateScore(values []byte) int {
	score := 0

	for i := 0; i < len(values); i++ {
		score += byte2score(values[i])
	}

	return score
}

func byte2score(c byte) int {
	ascii := int(c)
	if ascii > 96 && ascii < 123 {
		// lowercase
		return ascii - 96
	} else {
		return ascii - 64 + 26
	}
}
