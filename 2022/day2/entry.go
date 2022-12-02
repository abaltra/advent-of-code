package day2

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var SCORES = map[string]int{
  "X": 1,
  "Y": 2,
  "Z": 3,
  "A": 1,
  "B": 2,
  "C": 3,
  "LOSS": 0,
  "DRAW": 3,
  "WIN": 6,
}

var VALUES_TO_RESULTS = map[string]string {
  "X": "LOSS",
  "Y": "DRAW",
  "Z": "WIN",
}

//ROCK: A,X
//PAPER: B,Y
//SCISSORS: C,Z
var BEATS = map[string]string{
  "A": "Z", // 
  "B": "X",
  "C": "Y",
}

var LOSSES = map[string]string{
  "A": "Y",
  "B": "Z",
  "C": "X",
}

var EQUALITIES = map[string]string{
  "A": "X",
  "B": "Y",
  "C": "Z",
}

func processPlayMethod2(p1 string, result string) int {
  score_for_round := 0
  r := VALUES_TO_RESULTS[result]

  score_for_round += SCORES[r]

  var p2Hand string

  if r == "DRAW" {
    p2Hand = p1
  } else if (r == "WIN") {
    p2Hand = LOSSES[p1]
  } else {
    p2Hand = BEATS[p1]
  }

  score_for_round += SCORES[p2Hand]

  return score_for_round
}

func processPlay(p1 string, p2 string) int {
  score_for_round := SCORES[p2]
  var result string

  if EQUALITIES[p1] == p2 {
    result = "DRAW"
  } else if (BEATS[p1] == p2){
    // p1 wins
    result = "LOSS"
  } else {
    result = "WIN"
  }

  score_for_round += SCORES[result]

  return score_for_round

}

func Run() {
  println("hello world!")
  
  input, err := os.Open("./day2/input.txt")

	if err != nil {
		panic(err)
	}

	filescanner := bufio.NewScanner(input)
	filescanner.Split(bufio.ScanLines)

  running_score := 0
  running_score_method2 := 0
	for filescanner.Scan() {
		l := filescanner.Text()
    parts := strings.Split(l, " ")
    p1, p2 := parts[0], parts[1]
    running_score += processPlay(p1, p2)
    running_score_method2 += processPlayMethod2(p1, p2)
  }

  fmt.Printf("Total score is: %v\n", running_score)
  fmt.Printf("Total score with method 2: %v\n", running_score_method2)
}
