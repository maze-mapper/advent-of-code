// Advent of Code 2015 - Day 5
package day5

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strings"
	"sync"
)

var mutex sync.Mutex
var bannedStrings = regexp.MustCompile(`ab|cd|pq|xy`)

func isVowel(r rune) bool {
	return strings.ContainsAny(string(r), "aeiou")
}

func containsNVowels(s string, n int) bool {
	vowelCount := 0
	for _, c := range s {
		if isVowel(c) {
			vowelCount++
		}
		if vowelCount >= n {
			return true
		}
	}
	return false
}

func part1(line string, nice *int, wg *sync.WaitGroup) {
	if !bannedStrings.MatchString(line) {
		vowelCount := 0
		doubleLetter := false
		for i, c := range line {
			if isVowel(c) {
				vowelCount++
			}
			if i != 0 {
				if c == rune(line[i-1]) {
					doubleLetter = true
				}
			}

			if doubleLetter == true && vowelCount >= 3 {
				mutex.Lock()
				*nice++
				mutex.Unlock()
				break
			}
		}
	}
	wg.Done()
}

func repeatedLetterPair(s, p string) bool {
	if strings.Count(s, p) > 1 {
		return true
	}
	return false
}

func findSameLetterOneApart(s string, r rune) bool {
	s2 := string(r)
	matched, err := regexp.MatchString(s2+"[^"+s2+"]"+s2, s)
	if err != nil {
		fmt.Println(err)
	}
	return matched
}

func part2(line string, nice *int, wg *sync.WaitGroup) {
	foundRepeatedLetterPair := false
	foundSameLetterOneApart := false
	for i, c := range line {
		if !foundRepeatedLetterPair && i != len(line)-1 {
			foundRepeatedLetterPair = repeatedLetterPair(line, line[i:i+2])
		}

		if !foundSameLetterOneApart && i != len(line)-2 {
			foundSameLetterOneApart = findSameLetterOneApart(line, c)
		}
		if foundRepeatedLetterPair == true && foundSameLetterOneApart == true {
			mutex.Lock()
			*nice++
			mutex.Unlock()
			break
		}

	}
	wg.Done()
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	lines := strings.Split(string(data), "\n")

	var wg sync.WaitGroup
	var nice int
	for _, line := range lines {
		wg.Add(1)
		go part1(line, &nice, &wg)
	}
	wg.Wait()

	fmt.Println(nice)

	// Reset nice count for part 2
	nice = 0
	for _, line := range lines {
		wg.Add(1)
		go part2(line, &nice, &wg)
	}
	wg.Wait()

	fmt.Println(nice)
}
