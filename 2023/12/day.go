// Advent of Code 2023 - Day 12.
package day12

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"sync"
)

type springRecord struct {
	row     string
	damaged []int
}

func (sr springRecord) arrangements() int {
	count := match(sr.row, nil, sr.damaged, false)
	return count
}

func (sr springRecord) unfold(scale int) springRecord {
	var rowParts []string
	newDamaged := make([]int, scale*len(sr.damaged))
	l := len(sr.damaged)
	for i := 0; i < scale; i++ {
		rowParts = append(rowParts, sr.row)
		copy(newDamaged[i*l:(i+1)*l], sr.damaged)
	}

	return springRecord{
		row:     strings.Join(rowParts, "?"),
		damaged: newDamaged,
	}
}

// Add a "." to the end of the initial string before starting.
func dpMatch(s string, groups []int, cache map[string]int) int {
	// fmt.Println(s, groups)

	// Found a possibility.
	if len(s) == 0 && len(groups) == 0 {
		return 1
	}
	// Another possibility.
	if len(groups) == 0 && !strings.Contains(s, "#") {
		return 1
	}

	// String is too short to match remaining group and blank space.
	if len(s) == 0 || len(groups) == 0 || len(s) < groups[0]+1 {
		return 0
	}

	// Cache lookup.
	sGroups := make([]string, len(groups))
	for i, g := range groups {
		sGroups[i] = strconv.Itoa(g)
	}
	key := s + "_" + strings.Join(sGroups, "_")
	if v, ok := cache[key]; ok {
		return v
	}

	current := string(s[0])
	count := 0
	switch current {
	case ".":
		count = dpMatch(s[1:], groups, cache)
	case "?", "#":
		// fmt.Println("Current is", current)
		// Ways if it is a ".".
		if current == "?" {
			// fmt.Println("Treating '?' as '.'")
			count += dpMatch(s[1:], groups, cache)
		}

		// Check to see if the next part of the string could be damaged parts ("#" or "?")
		// and that it then can be a working part to end the group ("." or "?").
		nextGroupLen := groups[0]
		if !strings.Contains(s[:nextGroupLen], ".") && string(s[nextGroupLen]) != "#" {
			// Match a group.
			// fmt.Println("Matched a group!", nextGroupLen, s[:nextGroupLen], string(s[nextGroupLen]), s)
			// fmt.Println("Going to ", s[nextGroupLen+1:], groups[1:])
			count += dpMatch(s[nextGroupLen+1:], groups[1:], cache)
		} else {
			// fmt.Println("No group matched")
		}
	}

	cache[key] = count

	return count
}

func match(s string, got, want []int, onBroken bool) int {
	// fmt.Println("Match start")
	count := 0
	for j, r := range s {
		// fmt.Println(s, j, string(r), got, want)
		switch r {
		case '.':
			onBroken = false
		case '#':
			if !onBroken {
				// Check all previous before adding another.
				for i := range got {
					if got[i] != want[i] {
						return 0
					}
				}
				got = append(got, 0)
			}
			onBroken = true
			i := len(got) - 1
			got[i] += 1
			if i >= len(want) || got[i] > want[i] {
				// fmt.Println("A", got, want)
				return 0
			}
		case '?':
			c1 := make([]int, len(got))
			c2 := make([]int, len(got))
			copy(c1, got)
			copy(c2, got)
			return match("."+s[j+1:], c1, want, onBroken) + match("#"+s[j+1:], c2, want, onBroken)
		}
	}
	if len(got) != len(want) {
		// fmt.Println("B", got, want)
		return 0
	}
	for i := range want {
		if got[i] != want[i] {
			// fmt.Println("C", got, want)
			return 0
		}
	}
	count += 1
	// *total += 1
	return count
}

func parseData(data []byte) []springRecord {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	records := make([]springRecord, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		sr := springRecord{
			row: parts[0],
		}
		for _, s := range strings.Split(parts[1], ",") {
			n, err := strconv.Atoi(s)
			if err != nil {
				log.Fatal(err)
			}
			sr.damaged = append(sr.damaged, n)
		}
		records[i] = sr
	}
	return records
}

func part1(records []springRecord) int {
	total := 0
	var wg sync.WaitGroup
	ch := make(chan int)
	for _, sr := range records {
		wg.Add(1)
		sr := sr
		go func() {
			defer wg.Done()
			ch <- sr.arrangements()
		}()
	}

	var wg2 sync.WaitGroup
	wg2.Add(1)
	go func() {
		defer wg2.Done()
		// total += <-ch
		for c := range ch {
			total += c
		}
	}()
	wg.Wait()
	close(ch)
	wg2.Wait()
	return total
}

func part2(records []springRecord) int {
	newRecords := make([]springRecord, len(records))
	for i, sr := range records {
		newRecords[i] = sr.unfold(5)
	}

	total := 0
	cache := map[string]int{}
	for _, sr := range newRecords {
		total += dpMatch(sr.row+".", sr.damaged, cache)
	}
	return total

	// return part1(newRecords)
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	records := parseData(data)

	p1 := part1(records)
	fmt.Println("Part 1:", p1)

	p2 := part2(records)
	fmt.Println("Part 2:", p2)
}
