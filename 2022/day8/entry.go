package day8

import (
	"fmt"
	"strconv"
)

type Tree struct {
  height int
  isVisibile bool
  scenicScore int
  northScore int
  southScore int
  eastScore int
  westScore int
}

func Run(lines []string) {
  forest := make([][]*Tree, len(lines))
  width := len(lines[0])
  height := len(lines)
  for i, line := range lines {
    forest[i] = make([]*Tree, len(line))
    for j, char := range line {
      height, _ := strconv.Atoi(string(char))
      forest[i][j] = &Tree{
        height: height,
        isVisibile: false,
      }
    }
  }

  flagAsVisible(forest, height, width)
  printForest(forest)
  maxTree := calculateScores(forest, height, width)
  fmt.Printf("There are %d visible trees, and the best scored tree has %d points\n", countVisible(forest), maxTree.scenicScore)
}

func calculateScores(forest [][]*Tree, height, width int) *Tree {
  bestTree := &Tree{}

  for i, row := range forest {
    for j := range row {
      calculateScore(forest, i, j, height, width)
      if forest[i][j].scenicScore > bestTree.scenicScore {
        bestTree = forest[i][j]
      }
    }
  }

  return bestTree
}

func calculateScore(forest [][]*Tree, initialI, initialJ, maxI, maxJ int) {
  initialTree := forest[initialI][initialJ]

  // to the top
  for i := initialI - 1; i >= 0; i-- {
    initialTree.northScore ++
    if forest[i][initialJ].height >= initialTree.height {
      break
    }
  }

  // to the bottom
  for i := initialI + 1; i < maxI; i++ {
    initialTree.southScore ++
    if forest[i][initialJ].height >= initialTree.height {
      break
    }
  }

  // to the left
  for j := initialJ - 1; j >= 0; j-- {
    initialTree.westScore ++
    if forest[initialI][j].height >= initialTree.height {
      break
    }
  }
  
  // to the right
  for j := initialJ + 1; j < maxJ; j++ {
    initialTree.eastScore ++
    if forest[initialI][j].height >= initialTree.height {
      break
    }
  }

	initialTree.scenicScore = initialTree.northScore * initialTree.southScore * initialTree.westScore * initialTree.eastScore
}

func flagAsVisible(forest [][]*Tree, height, width int) {
  tallest := 0

  // FROM THE LEFT
  for i := 0; i < height; i++ {
    tallest = -1
    for j := 0; j < width; j++ {
      tree := forest[i][j]
      if tree.height > tallest {
        tree.isVisibile = true
        tallest = tree.height
      }
    }
  } 
  
  // FROM THE TOP
  for j := 0; j < width; j++ {
    tallest = -1
    for i := 0; i < height; i++ {
      tree := forest[i][j]
      if tree.height > tallest {
        tree.isVisibile = true
        tallest = tree.height
      }
    }
  }

  // FROM THE RIGHT
  for i := 0; i < height; i++ {
    tallest = -1
    for j := width - 1; j >= 0; j-- {
      tree := forest[i][j]
      if tree.height > tallest {
        tree.isVisibile = true
        tallest = tree.height
      }
    }
  }

  // FROM THE BOTTOM
  for j := 0; j < width; j++ {
    tallest = -1
    for i := height - 1; i >= 0; i-- {
      tree := forest[i][j]
      if tree.height > tallest {
        tree.isVisibile = true
        tallest = tree.height
      }
    }
  }
}

func countVisible(forest [][]*Tree) int {
  count := 0
  for _, row := range forest {
    for _, tree := range row {
      if tree.isVisibile {
        count ++
      }
    }
  }

  return count
}

func printForest(forest [][]*Tree) {
  for _, l := range forest {
    for _, t := range l {
      fmt.Print(t.height)
    }
    fmt.Print("\n")
  }
}
