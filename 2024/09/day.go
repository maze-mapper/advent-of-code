// Advent of Code 2024 - Day 9.
package day9

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func parseData(data []byte) ([]int, error) {
	s := strings.TrimSuffix(string(data), "\n")
	var arr []int
	var isFree bool
	var id int
	for _, r := range s {
		n, err := strconv.Atoi(string(r))
		if err != nil {
			return nil, err
		}
		for i := 0; i < n; i++ {
			if isFree {
				arr = append(arr, -1)
			} else {
				arr = append(arr, id)
			}
		}
		if isFree {
			id += 1
		}
		isFree = !isFree
	}
	return arr, nil
}

func part1(arr []int) int {
	l := 0
	r := len(arr) - 1
	for l < r {
		if arr[l] >= 0 {
			l++
			continue
		}
		if arr[r] < 0 {
			r--
			continue
		}
		arr[l] = arr[r]
		arr[r] = -1
		r--
	}
	return checksum(arr)
}

func part2(arr []int) int {
	// Last element is part of a file. The first file never needs to move.
	for fileID := arr[len(arr)-1]; fileID > 0; fileID-- {
		leftPos, rightPos := findFile(arr, fileID)
		length := rightPos - leftPos + 1
		pos, ok := findFreeSpace(arr[:leftPos], length)
		if ok {
			for i := leftPos; i <= rightPos; i++ {
				arr[i] = -1
			}
			for i := pos; i < pos+length; i++ {
				arr[i] = fileID
			}
		}
	}
	return checksum(arr)
}

func checksum(arr []int) int {
	var sum int
	for i, n := range arr {
		if n < 0 {
			continue
		}
		sum += i * n
	}
	return sum
}

func findFile(arr []int, id int) (int, int) {
	var leftPos int
	for i := 0; i < len(arr); i++ {
		if arr[i] == id {
			leftPos = i
			break
		}
	}
	var rightPos int
	for i := len(arr) - 1; i >= 0; i-- {
		if arr[i] == id {
			rightPos = i
			break
		}
	}
	return leftPos, rightPos
}

func findFreeSpace(arr []int, l int) (int, bool) {
	var currentFree int
	for i, n := range arr {
		if n >= 0 {
			currentFree = 0
			continue
		}
		currentFree += 1
		if currentFree == l {
			return i - currentFree + 1, true
		}
	}
	return 0, false
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	arr, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	arr2 := make([]int, len(arr))
	copy(arr2, arr)

	p1 := part1(arr)
	fmt.Println("Part 1:", p1)

	p2 := part2(arr2)
	fmt.Println("Part 2:", p2)
}
