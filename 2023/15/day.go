// Advent of Code 2023 - Day 15.
package day15

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"strconv"
)

func hashAlgortihm(input []byte) int {
	value := 0
	for _, b := range input {
		value += int(b)
		value *= 17
		value %= 256
	}
	return value
}

func part1(input [][]byte) int {
	total := 0
	for _, bs := range input {
		total += hashAlgortihm(bs)
	}
	return total
}

type lens struct {
	label       string
	focalLength int
}

type actionType int

const (
	actionUnknown actionType = iota
	actionDash
	actionEquals
)

func focusingPower(m map[int][]*lens) int {
	power := 0
	for box, lenses := range m {
		for j, l := range lenses {
			power += (box + 1) * (j + 1) * l.focalLength
		}
	}
	return power
}

func part2(input [][]byte) int {
	m := map[int][]*lens{}
	for _, bs := range input {
		var label []byte
		var action actionType
		var focalLength int
		if bytes.HasSuffix(bs, []byte{'-'}) {
			label = bytes.TrimSuffix(bs, []byte{'-'})
			action = actionDash
		} else {
			parts := bytes.Split(bs, []byte{'='})
			label = parts[0]
			n, err := strconv.Atoi(string(parts[1]))
			if err != nil {
				log.Fatal(err)
			}
			focalLength = n
			action = actionEquals
		}
		box := hashAlgortihm(label)

		switch action {
		case actionDash:
			for i, l := range m[box] {
				if l.label == string(label) {
					m[box] = append(m[box][:i], m[box][i+1:]...)
				}
			}
		case actionEquals:
			found := false
			for _, l := range m[box] {
				if l.label == string(label) {
					found = true
					l.focalLength = focalLength
					break
				}
			}
			if !found {
				m[box] = append(m[box], &lens{label: string(label), focalLength: focalLength})
			}
		}
	}
	return focusingPower(m)
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	input := bytes.Split(bytes.TrimSuffix(data, []byte{'\n'}), []byte{','})

	p1 := part1(input)
	fmt.Println("Part 1:", p1)

	p2 := part2(input)
	fmt.Println("Part 2:", p2)
}
