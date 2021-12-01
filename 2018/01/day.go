// Advent of Code 2018 - Day 1
package day1

import (
	"container/ring"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

type void struct{}

func parseInput(file string) []int {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	changes := make([]int, len(lines))
	for i, line := range lines {
		if v, err := strconv.Atoi(line); err == nil {
			changes[i] = v
		} else {
			log.Fatal(err)
		}
	}

	return changes
}

func sliceToRing(sl []int) *ring.Ring {
	r := ring.New(len(sl))
	for i := 0; i < r.Len(); i++ {
		r.Value = sl[i]
		r = r.Next()
	}
	return r
}

func part1(changes []int) int {
	freq := 0
	for _, c := range changes {
		freq += c
	}
	return freq
}

func part2(changes []int) int {
	freq := 0
	prevFreq := map[int]void{freq: void{}}
	r := sliceToRing(changes)

	for {
		freq += r.Value.(int)
		if _, ok := prevFreq[freq]; ok == true {
			return freq
		} else {
			prevFreq[freq] = void{}
		}
		r = r.Next()
	}
}

func Run(inputFile string) {
	changes := parseInput(inputFile)

	p1 := part1(changes)
	fmt.Println("Part 1:", p1)

	p2 := part2(changes)
	fmt.Println("Part 2:", p2)
}
