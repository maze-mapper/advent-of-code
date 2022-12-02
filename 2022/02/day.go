// Advent of Code 2022 - Day 2.
package day2

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// Scores for playing each shape.
const (
	shapeRock = 1
	shapePaper = 2
	shapeScissors = 3
)

// Scores for each game outcome.
const (
	outcomeLoss = 0
	outcomeDraw = 3
	outcomeWin = 6
)

func parseInput(data []byte) [][]string {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	output := make([][]string, len(lines))
	for i, line := range lines {
		parts := strings.Split(line, " ")
		output[i] = parts
	}
	return output
}

func part1(guide [][]string) int {
	shapes := map[string]int{
		"A": shapeRock,
		"B": shapePaper,
		"C": shapeScissors,
		"X": shapeRock,
                "Y": shapePaper,
                "Z": shapeScissors,
	}
	score := 0
	for _, g := range guide {
		opponentShape, ok := shapes[g[0]]
		if !ok {
                        log.Fatal("Invalid shape %s", g[0])
                }
		yourShape, ok := shapes[g[1]]
		if !ok {
			log.Fatal("Invalid shape %s", g[1])
		}
		score += yourShape

		switch {
		// Win
                case (opponentShape == shapeRock && yourShape == shapePaper) || (opponentShape == shapePaper && yourShape == shapeScissors) || (opponentShape == shapeScissors && yourShape == shapeRock):
                        score += outcomeWin
		// Draw
                case (opponentShape == shapeRock && yourShape == shapeRock) || (opponentShape == shapePaper && yourShape == shapePaper) || (opponentShape == shapeScissors && yourShape == shapeScissors):
                        score += outcomeDraw
		// Loss
		case (opponentShape == shapeRock && yourShape == shapeScissors) || (opponentShape == shapePaper && yourShape == shapeRock) || (opponentShape == shapeScissors && yourShape == shapePaper):
                        score += outcomeLoss
		}
	}
	return score
}

func part2(guide [][]string) int {
	shapes := map[string]int{
                "A": shapeRock,
                "B": shapePaper,
                "C": shapeScissors,
        }
	outcomes := map[string]int{
		"X": outcomeLoss,
		"Y": outcomeDraw,
		"Z": outcomeWin,
	}
	score := 0
	for _, g := range guide {
		opponentShape, ok := shapes[g[0]]
                if !ok {
                        log.Fatal("Invalid shape %s", g[0])
                }
		outcome, ok := outcomes[g[1]]
		if !ok {
                        log.Fatal("Invalid outcome %s", g[1])
                }
		score += outcome

		switch {
		// You play rock.
                case (opponentShape == shapeRock && outcome == outcomeDraw) || (opponentShape == shapePaper && outcome == outcomeLoss) || (opponentShape == shapeScissors && outcome == outcomeWin):
			score += shapeRock
		// You play paper.
		case (opponentShape == shapeRock && outcome == outcomeWin) || (opponentShape == shapePaper && outcome == outcomeDraw) || (opponentShape == shapeScissors && outcome == outcomeLoss):
			score += shapePaper
		// You play scissors.
		case (opponentShape == shapeRock && outcome == outcomeLoss) || (opponentShape == shapePaper && outcome == outcomeWin) || (opponentShape == shapeScissors && outcome == outcomeDraw):
                        score += shapeScissors
		}
	}
	return score
}

func Run(inputFile string) {
        data, err := ioutil.ReadFile(inputFile)
        if err != nil {
                log.Fatal(err)
        }
        guide := parseInput(data)

        p1 := part1(guide)
        fmt.Println("Part 1:", p1)

        p2 := part2(guide)
        fmt.Println("Part 2:", p2)
}

