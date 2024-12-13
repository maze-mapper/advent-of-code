// Advent of Code 2024 - Day 13.
package day13

import (
	"fmt"
	"log"
	"math"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

type machine struct {
	a, b, prize coordinates.Coord
}

// aPressesToPrize calculates the number of A button presses to reach the prize.
// The returned boolean indicates if reaching the prize is possible.
// The formula is derived from the equations:
//
//	(A presses)*(A X distance) + (B presses)*(B X distance) = (Prize X position)
//	(A presses)*(A Y distance) + (B presses)*(B Y distance) = (Prize Y position)
//
// It fails if any of A, B or Prize have the same unit vector.
func (m machine) aPressesToPrize() (int, bool) {
	numerator := m.prize.Y*m.b.X - m.prize.X*m.b.Y
	denominator := m.a.Y*m.b.X - m.a.X*m.b.Y
	if denominator == 0 {
		return 0, false
	}
	if numerator%denominator != 0 {
		return 0, false
	}
	return numerator / denominator, true
}

// bPressesToPrize calculates the number of B button presses to reach the prize.
// The returned boolean indicates if reaching the prize is possible.
// The formula is derived from the equations:
//
//	(A presses)*(A X distance) + (B presses)*(B X distance) = (Prize X position)
//	(A presses)*(A Y distance) + (B presses)*(B Y distance) = (Prize Y position)
//
// It fails if any of A, B or Prize have the same unit vector.
func (m machine) bPressesToPrize() (int, bool) {
	numerator := m.prize.Y*m.a.X - m.prize.X*m.a.Y
	denominator := m.a.X*m.b.Y - m.a.Y*m.b.X
	if denominator == 0 {
		return 0, false
	}
	if numerator%denominator != 0 {
		return 0, false
	}
	return numerator / denominator, true
}

func parseData(data []byte) ([]machine, error) {
	sections := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	machines := make([]machine, len(sections))
	for i, section := range sections {
		section = strings.ReplaceAll(section, "+", " ")
		section = strings.ReplaceAll(section, ",", "")
		section = strings.ReplaceAll(section, "=", " ")
		m := machine{}
		n, err := fmt.Sscanf(section, "Button A: X %d Y %d\nButton B: X %d Y %d\nPrize: X %d Y %d", &m.a.X, &m.a.Y, &m.b.X, &m.b.Y, &m.prize.X, &m.prize.Y)
		if err != nil {
			return nil, err
		}
		if n != 6 {
			return nil, fmt.Errorf("expected to parse 6 items, instead parsed %d", n)
		}
		machines[i] = m
	}
	return machines, nil
}

func part1(machines []machine) int {
	var totalTokens int
	maxButtonPresses := 100
	for _, m := range machines {
		minTokens := math.MaxInt
		var isWinnable bool
		for aPresses := 0; aPresses < maxButtonPresses; aPresses++ {
			for bPresses := 0; bPresses < maxButtonPresses; bPresses++ {
				pos := coordinates.Coord{
					X: aPresses*m.a.X + bPresses*m.b.X,
					Y: aPresses*m.a.Y + bPresses*m.b.Y,
				}
				if pos == m.prize {
					isWinnable = true
					tokens := 3*aPresses + bPresses
					if tokens < minTokens {
						minTokens = tokens
					}
				}
			}
		}
		if isWinnable {
			totalTokens += minTokens
		}
	}
	return totalTokens
}

func part2(machines []machine) int {
	var totalTokens int
	for _, m := range machines {
		m.prize.X += 10000000000000
		m.prize.Y += 10000000000000

		aPresses, ok := m.aPressesToPrize()
		if !ok {
			continue
		}
		bPresses, ok := m.bPressesToPrize()
		if !ok {
			continue
		}

		totalTokens += 3*aPresses + bPresses
	}
	return totalTokens
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	machines, err := parseData(data)
	if err != nil {
		log.Fatal(err)
	}

	p1 := part1(machines)
	fmt.Println("Part 1:", p1)

	p2 := part2(machines)
	fmt.Println("Part 2:", p2)
}
