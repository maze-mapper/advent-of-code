// Advent of Code 2022 - Day 7.
package day7

import (
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"strconv"
	"strings"
)

type treeNode struct {
	parent   *treeNode
	children map[string]*treeNode
	files    map[string]int
}

func newTreeNode(parent *treeNode) *treeNode {
	return &treeNode{
		parent:   parent,
		children: map[string]*treeNode{},
		files:    map[string]int{},
	}
}

func parseInput(data []byte) *treeNode {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	// First line is root directory.
	if lines[0] != "$ cd /" {
		log.Fatal("Expecting first line to be \"$ cd /\"")
	}
	root := newTreeNode(nil)
	n := root
	for i := 1; i < len(lines); i++ {
		if strings.HasPrefix(lines[i], "$ cd ") {
			dir := strings.TrimPrefix(lines[i], "$ cd ")
			if dir == ".." {
				n = n.parent
			} else {
				n = n.children[dir]
			}
		} else if lines[i] == "$ ls" {
			i += 1
			for ; i < len(lines) && !strings.HasPrefix(lines[i], "$"); i++ {

				if strings.HasPrefix(lines[i], "dir ") {
					dir := strings.TrimPrefix(lines[i], "dir ")
					if _, ok := n.children[dir]; !ok {
						n.children[dir] = newTreeNode(n)
					}
				} else {
					parts := strings.Split(lines[i], " ")
					size, err := strconv.Atoi(parts[0])
					if err != nil {
						log.Fatal(err)
					}
					n.files[parts[1]] = size
				}
			}
			i -= 1
		}

	}
	return root
}

func calculateDirSizes(path string, node *treeNode, sizes map[string]int) {
	totalSize := 0
	for _, size := range node.files {
		totalSize += size
	}
	for k, v := range node.children {
		childPath := path + k + "/"
		calculateDirSizes(childPath, v, sizes)
		totalSize += sizes[childPath]
	}
	sizes[path] = totalSize
}

func part1(sizes map[string]int) int {
	sum := 0
	for _, v := range sizes {
		if v <= 100000 {
			sum += v
		}
	}
	return sum
}

func part2(sizes map[string]int) int {
	totalSpace := 70000000
	requiredSpace := 30000000
	freeSpace := totalSpace - sizes["/"]
	minSpaceToDelete := requiredSpace - freeSpace

	dirSize := int(math.MaxInt)
	for _, v := range sizes {
		if v >= minSpaceToDelete && v < dirSize {
			dirSize = v
		}
	}

	return dirSize
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	directoryTree := parseInput(data)
	sizes := map[string]int{}
	calculateDirSizes("/", directoryTree, sizes)

	p1 := part1(sizes)
	fmt.Println("Part 1:", p1)

	p2 := part2(sizes)
	fmt.Println("Part 2:", p2)
}
