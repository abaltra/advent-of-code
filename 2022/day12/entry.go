package day12

import (
	"fmt"
	"strings"
)

type Node struct {
	I int `json:"I"`
	J int `json:"J"`
	Value byte `json:"VALUE"`
  Elevation int `json:"Elevation"`
}

func (n Node) ToString() string {
  return fmt.Sprintf("%d,%d (%v)", n.I, n.J, string(n.Value))
}

func lineToNodes(line string, i int) []*Node {
  nodes := make([]*Node, len(line))

  for j, c := range line {
    n := &Node{
      I: i,
      J: j,
      Value: byte(c),
    }

    if n.Value == 'S' {
      n.Elevation = int('a')
    } else if n.Value == 'E' {
      n.Elevation = int('z')
    } else {
      n.Elevation = int(c)
    }
    nodes[j] = n
  }

  return nodes
}

func getAdjacentNodes(graph[][]*Node, baseNode Node) []*Node {
  neighbors := []*Node{}
  //fmt.Printf("Getting neighbors for node in %d,%d\n", baseNode.i, baseNode.j)

  // UP
  if baseNode.I > 0 && graph[baseNode.I - 1][baseNode.J].Elevation - baseNode.Elevation <= 1 {
    neighbors = append(neighbors, graph[baseNode.I - 1][baseNode.J])
  }

  //DOWN
  if baseNode.I < len(graph) - 1 && graph[baseNode.I + 1][baseNode.J].Elevation - baseNode.Elevation <= 1 {
    neighbors = append(neighbors, graph[baseNode.I + 1][baseNode.J])
  }

  // LEFT
  if baseNode.J > 0 && graph[baseNode.I][baseNode.J - 1].Elevation - baseNode.Elevation <= 1 {
    neighbors = append(neighbors, graph[baseNode.I][baseNode.J - 1])
  }

  // RIGHT
  if baseNode.J < len(graph[baseNode.I]) - 1 && graph[baseNode.I][baseNode.J + 1].Elevation - baseNode.Elevation <= 1 {
    neighbors = append(neighbors, graph[baseNode.I][baseNode.J + 1])
  }

  return neighbors
}

func printQueue(queue []*Node) {
  for _, node := range queue {
    println(node.ToString())
  }
}

func printGraph(graph [][]*Node) {
  for _, row := range graph {
    for _, node := range row {
      fmt.Print(string(node.Value))
    }
    println()
  }
}

func BFS(graph [][]*Node, start, target *Node) int {
  visited := map[string]bool{}
  queue := []*Node { start }
  depth := 0

  for len(queue) > 0 {
    nodes := queue[0:]
    queue = []*Node {}

    for _, node := range nodes {

      if node.Value == target.Value {
        return depth
      }
      visited[node.ToString()] = true
      adjacent := getAdjacentNodes(graph, *node)

      for _, aNode := range adjacent {
        if visited[aNode.ToString()] {
          continue
        }

        queue = append(queue, aNode)
        visited[aNode.ToString()] = true


      }
    }
    depth ++
  }

  return -1
}

func Run(lines []string) {
	println("day 12")
	heightMap := [][]*Node{}

	var startPos *Node
	var targetPos *Node

	var bestStartingPoint *Node
	leastSteps := 10000000

	for i, line := range lines {
	  nodes := lineToNodes(line, i)
		heightMap = append(heightMap, nodes)
    sIndex := strings.Index(line, "S")
    if sIndex > -1 {
      println("FOUND START")
      startPos = heightMap[i][sIndex]
    }

    tIndex := strings.Index(line, "E")
    if tIndex > -1 {
      println("FOUND END")
      targetPos = heightMap[i][tIndex]
    }
  }

  //n := getAdjacentNodes(heightMap, *heightMap[0][2])
  //fmt.Printf("Checking node %s\n", heightMap[0][2].ToString())
  //printQueue(n)

  distance := BFS(heightMap, startPos, targetPos)

  for _, row := range heightMap {
    for _, node := range row {
      if node.Elevation != int('a') || node.Value == 'E' {
        continue
      }
      d := BFS(heightMap, node, targetPos)
      if d > -1 && d < leastSteps {
        leastSteps = d
        bestStartingPoint = node
      }

    }
  }

  fmt.Printf("Took %d steps to get to the target\n", distance)
  fmt.Printf("The best starting point is %s, it only takes %d steps\n", bestStartingPoint.ToString(), leastSteps)
}
