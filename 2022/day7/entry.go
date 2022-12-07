package day7

import (
	"fmt"
	"strconv"
	"strings"
)

type NodeType int
const (
	FILE NodeType = iota
	DIRECTORY
)

const MAX_SPACE = 70000000
const TARGET_SPACE = 30000000

type Node struct {
	name string
	size int
	nodeType NodeType
	children map[string]*Node
	parent *Node
}

func newNode(name string, size int, nodeType NodeType, parent *Node) *Node {
	return &Node{
		name: name,
		size: size,
		nodeType: nodeType,
		parent: parent,
		children: make(map[string]*Node),
	}
}

func Run(lines []string) {

	var root *Node
	var activeNode *Node
	for _, l := range lines {
		if strings.HasPrefix(l, "$ cd") {
			parts := strings.Split(l, " ")
			if root == nil {
				root = newNode(parts[2], 0, DIRECTORY, nil)
				activeNode = root
				continue
			}
			if parts[2] == ".." {
				activeNode = activeNode.parent
			} else {
				activeNode = activeNode.children[parts[2]]
			}
		} else if strings.HasPrefix(l, "$ ls") {
			// do nothing
		} else {
			parts := strings.Split(l, " ")
			if parts[0] == "dir" {
				activeNode.children[parts[1]] = newNode(parts[1], 0, DIRECTORY, activeNode)
			} else {
				size, _ := strconv.Atoi(parts[0])
				activeNode.children[parts[1]] = newNode(parts[1], size, FILE, activeNode)
			}
		}
	}

	totalSizeOfLargeFolders(root, 100000)
	printTree(root, 0)
	freeSpace := MAX_SPACE - root.size
	if freeSpace < TARGET_SPACE {
		fmt.Printf("We currently have %d free space, which is less than the target of %d. We need to free at least %d\n", freeSpace, TARGET_SPACE, TARGET_SPACE - freeSpace)
		dirToDelete := findDirectoryToDelete(root, TARGET_SPACE - freeSpace)
		fmt.Printf("We'll delete dir %s, freeing up %d bytes \n", dirToDelete.name, dirToDelete.size)
	}
}

var overSizedTotal int = 0
var bytesToRemove int = -1
var spaceToFree int = 0

func totalSizeOfLargeFolders(node *Node, limit int) int {
	if node == nil {
		return 0
	}

	if node.nodeType == FILE {
		return node.size
	}

	sizeOfChildren := 0
	for _, child := range node.children {
		sizeOfChildren += totalSizeOfLargeFolders(child, limit)
	}

	if sizeOfChildren < limit {
		overSizedTotal += sizeOfChildren
	}

	node.size = sizeOfChildren
	return sizeOfChildren
}

func findDirectoryToDelete(node *Node, target int) *Node {
	if node == nil {
		return nil
	}

	if node.nodeType != DIRECTORY {
		return nil
	}

	var toDelete *Node = nil
	for _, child := range node.children {
		option := findDirectoryToDelete(child, target)
		if option != nil && (toDelete == nil || toDelete.size > option.size) {
			toDelete = option
		}
	}

	if toDelete != nil {
		return toDelete
	}

	if node.size < target {
		return nil
	}

	return node
}

func printTree(node *Node, depth int) {
	if node == nil {
		return
	}
	typeString := fmt.Sprintf("dir, size=%d", node.size)

	if node.nodeType == FILE {
		typeString = fmt.Sprintf("file, size=%d", node.size)
	}
	
	fmt.Printf("%s- %s (%s)\n", addTabs(depth), node.name, typeString)

	for _, child := range node.children {
		printTree(child, depth + 1)
	}
}

func addTabs(total int) string {
	tabs := ""
	for i := 0; i < total; i++ {
		tabs += "\t"
	}

	return tabs
}
