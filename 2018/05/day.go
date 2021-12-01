// Advent of Code 2018 - Day 5
package day5

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"sync"
)

func abs(i int) int {
	if i < 0 {
		return -i
	}
	return i
}

func checkOffset(a, b byte) bool {
	if abs(int(b)-int(a)) == 32 {
		return true
	}
	return false
}

func reduce(b []byte) ([]byte, bool) {
	for i := 0; i < len(b)-1; i++ {
		if checkOffset(b[i], b[i+1]) {
			return append(b[:i], b[i+2:]...), true
		}
	}
	return b, false
}

func part1(b []byte) int {
	bb := make([]byte, len(b))
	copy(bb, b)
	for canReduce := true; canReduce; {
		bb, canReduce = reduce(bb)
	}
	return len(bb)
}

func part2(b []byte) int {
	var wg sync.WaitGroup
	var mu sync.Mutex
	result := int(^uint(0) >> 1) // Initialise to high number

	for letter := []byte("A")[0]; letter <= []byte("Z")[0]; letter++ {
		wg.Add(1)
		bb := bytes.ReplaceAll(b, []byte{letter}, []byte(""))
		bb = bytes.ReplaceAll(bb, []byte{letter + 32}, []byte(""))
		go func() {
			defer wg.Done()
			r := part1(bb)
			mu.Lock()
			if r < result {
				result = r
			}
			mu.Unlock()
		}()
	}
	wg.Wait()
	return result
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	// Strip trailing new line
	data = bytes.TrimSuffix(data, []byte("\n"))

	p1 := part1(data)
	fmt.Println("Part 1:", p1)

	p2 := part2(data)
	fmt.Println("Part 2:", p2)
}
