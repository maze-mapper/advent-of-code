package day25

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

const (
	eastMoving string = ">"
	southMoving string = "v"
)

func parseData(data []byte) [][]string {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	seaCucumbers := make([][]string, len(lines))
	for i, line := range lines {
		seaCucumbers[i] = make([]string, len(line))
		for j, r := range line {
			s := string(r)
			if s == eastMoving || s == southMoving {
				seaCucumbers[i][j] = s
			}
		}
	}

	return seaCucumbers
}

func moveHerd(seaCucumbers [][]string, moving string) ([][]string, bool) {
	newSeaCucumbers := make([][]string, len(seaCucumbers))
	for i, row := range seaCucumbers {
		newSeaCucumbers [i] = make([]string, len(row))
	}

	moved := false
	for i, row := range seaCucumbers {
		for j, elem := range row {
			switch elem {
			case eastMoving:
				if elem == moving {
					newJ := (j + 1) % len(row)
					if seaCucumbers[i][newJ] == "" {
						newSeaCucumbers[i][newJ] = elem
						moved = true
					} else {
						newSeaCucumbers[i][j] = elem
					}
				} else {
					newSeaCucumbers[i][j] = elem
				}

			case southMoving:
				if elem == moving {
					newI := (i + 1) % len(seaCucumbers)
					if seaCucumbers[newI][j] == "" {
						newSeaCucumbers[newI][j] = elem
						moved = true
					} else {
						newSeaCucumbers[i][j] = elem
					}
				} else {
					newSeaCucumbers[i][j] = elem
				}
			}
		}
	}
	return newSeaCucumbers, moved
}

func solve(seaCucumbers [][]string) int {
	moved := true
	steps := 0
	for moved {
		var movedEast, movedSouth bool
		seaCucumbers, movedEast = moveHerd(seaCucumbers, eastMoving)
		seaCucumbers, movedSouth = moveHerd(seaCucumbers, southMoving)
		steps += 1
		moved = movedEast || movedSouth
	}
	return steps
}

func part1(seaCucumbers [][]string) int {
	return solve(seaCucumbers)
}

func Run(inputFile string) {
        data, err := ioutil.ReadFile(inputFile)
        if err != nil {
                log.Fatal(err)
        }
        seaCucumbers := parseData(data)

        p1 := part1(seaCucumbers)
        fmt.Println("Part 1:", p1)
}

