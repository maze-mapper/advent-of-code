// Advent of Code 2021 - Day 23
package day23

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// moveCosts are the costs to move each type of amphipod
var moveCosts = map[string]int{
	"A": 1,
	"B": 10,
	"C": 100,
	"D": 1000,
}

const (
	burrowA int = 2
	burrowB int = 4
	burrowC int = 6
	burrowD int = 8
)

var burrowLocations = map[string]int{
	"A": burrowA,
	"B": burrowB,
	"C": burrowC,
	"D": burrowD,
}

var burrowIndex = map[string]int{
	"A": 0,
	"B": 1,
	"C": 2,
	"D": 3,
}

// burrowLen is the length of the hallway
const burrowLen int = 11

type burrow []string

func parseData(data []byte) [4]burrow {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	burrows := [4]burrow{}
	for i := 0; i < len(burrows); i++ {
		burrows[i] = make(burrow, 2)
	}
	for level, line := range lines[2:4] {
		b := 0
		for _, r := range line {
			switch r {
			case 'A', 'B', 'C', 'D':
				burrows[b][level] = string(r)
				b += 1
			}
		}
	}
	return burrows
}

func hallwayClear(i, j int, hallway [burrowLen]string) (int, bool) {
	if i < j {
		for _, s := range hallway[i+1 : j+1] {
			if s != "" {
				return 0, false
			}
		}
	} else {
		for _, s := range hallway[j:i] {
			if s != "" {
				return 0, false
			}
		}
	}

	diff := 0
	if j > i {
		diff = j - i
	} else {
		diff = i - j
	}

	return diff, true
}

func hallwayMoves(i int, hallway [burrowLen]string) []int {
	availableMoves := []int{}

	// To the left
	for j := i - 1; j >= 0; j-- {
		if hallway[j] != "" {
			break
		}
		switch j {
		case burrowA, burrowB, burrowC, burrowD:
			// Not a valid move
		default:
			availableMoves = append(availableMoves, j)
		}
	}

	// To the right
	for j := i + 1; j < len(hallway); j++ {
		if hallway[j] != "" {
			break
		}
		switch j {
		case burrowA, burrowB, burrowC, burrowD:
			// Not a valid move
		default:
			availableMoves = append(availableMoves, j)
		}
	}

	fmt.Println("availableMoves from", i, availableMoves)
	return availableMoves
}

func nextAmphipod(b burrow, s string) (int, bool) {
	lastAmphipodFixed := true
	available := false
	idx := 0
	for i := len(b) - 1; i >= 0; i-- {
		// Check to see if the current stack is incomplete
		if lastAmphipodFixed && b[i] != s {
			lastAmphipodFixed = false
		}

		if !lastAmphipodFixed && b[i] != "" {
			idx = i
			available = true
		}
	}
	return idx, available
}

func burrowAvailable(burrows [4]burrow, s string) (int, bool) {
	freeSpaces := 0
	for _, b := range burrows[burrowIndex[s]] {
		if b == "" {
			freeSpaces += 1
		}
		if b != s && b != "" {
			return freeSpaces, false
		}
	}
	return freeSpaces, true
}

func won(burrows [4]burrow) bool {
	for s, idx := range burrowIndex {
		for _, b := range burrows[idx] {
			if b != s {
				return false
			}
		}
	}
	return true
}

func moveFromHallwayToBurrows(burrows *[4]burrow, hallway *[burrowLen]string, score *int) bool {
	// Check if there are any amphipods in the hallway that could be moved in to a burrow
	moved := false
	for i, s := range hallway {
		switch s {
		case "A", "B", "C", "D":
			hSteps, hClear := hallwayClear(i, burrowLocations[s], *hallway)
			bSteps, bClear := burrowAvailable(*burrows, s)
			fmt.Println(hClear, bClear)
			if hClear && bClear {
				// Move amphipod
				moved = true
				burrows[burrowIndex[s]][bSteps-1] = s
				hallway[i] = ""
				*score += (hSteps + bSteps) * moveCosts[s]
				fmt.Println("moving amphipod from hallway to burrow with", hSteps+bSteps, "steps, score now", *score)
				printBurrow(*burrows, *hallway)
			}
		}
	}
	return moved
}

func moveFromRoomToRoom(rooms *[4]burrow, hallway *[burrowLen]string, score *int) bool {
	moved := false
	for room, i := range burrowIndex {
		// Get next amphipod that can move from this room
		amphipodIdx, ok := nextAmphipod(rooms[i], room)
		if !ok {
			continue
		}

		amphipod := rooms[i][amphipodIdx]
		hSteps, hClear := hallwayClear(burrowLocations[room], burrowLocations[amphipod], *hallway)
		bSteps, bClear := burrowAvailable(*rooms, amphipod)
		if hClear && bClear {
			// Move amphipod
			moved = true
			rooms[burrowIndex[amphipod]][bSteps-1] = amphipod
			rooms[i][amphipodIdx] = ""

			*score += (amphipodIdx + 1 + hSteps + bSteps) * moveCosts[amphipod]
			fmt.Println("moving amphipod from room to room with", hSteps+bSteps+amphipodIdx+1, "steps, score now", *score)
			printBurrow(*rooms, *hallway)
		}

	}
	return moved
}

func move(burrows [4]burrow, hallway [burrowLen]string, score int, minScore *int) {
	printBurrow(burrows, hallway)

	// Try to move any amphipods in burrows to their destination burrow
	canMoveFromBurrow1 := true
	for canMoveFromBurrow1 {
		canMoveFromBurrow1 = moveFromRoomToRoom(&burrows, &hallway, &score)
	}

	// Check if there are any amphipods in the hallway that could be moved in to a burrow
	canMoveFromHallway := true
	for canMoveFromHallway {
		canMoveFromHallway = moveFromHallwayToBurrows(&burrows, &hallway, &score)
	}

	// Try to move any amphipods in burrows to their destination burrow
	canMoveFromBurrow := true
	for canMoveFromBurrow {
		canMoveFromBurrow = moveFromRoomToRoom(&burrows, &hallway, &score)
	}

	// Abort if current score is already worse than the min score found so far
	if score > *minScore {
		return
	}

	// Check if won
	if won(burrows) {
		fmt.Println("======================================== WON")
		*minScore = score
		return
	}

	// Find amphipods in burrows that can be moved out to hallway
	for a, i := range burrowIndex {
		// Get amphipod
		amphipodIdx, ok := nextAmphipod(burrows[i], a)
		if !ok {
			continue
		}

		// Get possible move locations
		hMoves := hallwayMoves(burrowLocations[a], hallway)

		for _, hMove := range hMoves {
			amphipod := burrows[i][amphipodIdx]

			newBurrows := [4]burrow{}
			for i, b := range burrows {
				newBurrows[i] = make(burrow, len(burrows[i]))
				copy(newBurrows[i], b)
			}
			fmt.Println("Burrows:", burrows, newBurrows)
			newBurrows[i][amphipodIdx] = ""

			newHallway := hallway
			newHallway[hMove] = amphipod

			newScore := score
			newScore += (amphipodIdx + 1) * moveCosts[amphipod]
			hallwaySteps := 0
			if hMove > burrowLocations[a] {
				hallwaySteps = hMove - burrowLocations[a]
			} else {
				hallwaySteps = burrowLocations[a] - hMove
			}
			newScore += hallwaySteps * moveCosts[amphipod]

			fmt.Println("Moving to", hMove, "(", hallwaySteps, "steps)")
			fmt.Println("Burrows:", burrows, newBurrows)
			fmt.Println("Hallway:", hallway, newHallway, len(hallway), len(newHallway))
			fmt.Println(score, newScore)

			move(newBurrows, newHallway, newScore, minScore)
		}
	}

}

func printBurrow(burrows [4]burrow, hallway [burrowLen]string) {
	fmt.Println("#############")

	fmt.Print("#")
	for _, s := range hallway {
		if s == "" {
			print(".")
		} else {
			print(s)
		}
	}
	fmt.Print("#\n")

	fmt.Print("###")
	for _, b := range burrows {
		if b[0] == "" {
			print(".")
		} else {
			fmt.Print(b[0])
		}
		fmt.Print("#")
	}
	fmt.Print("##\n")

	for i := 1; i < len(burrows[0]); i++ {
		fmt.Print("  #")
		for j := 0; j < len(burrows); j++ {
			if burrows[j][i] == "" {
				print(".")
			} else {
				fmt.Print(burrows[j][i])
			}
			fmt.Print("#")
		}
		fmt.Print("\n")
	}

	fmt.Println("  #########")
}

func solve(burrows [4]burrow) int {
	hallway := [burrowLen]string{}
	minScore := int(^uint(0) >> 1)
	move(burrows, hallway, 0, &minScore)
	return minScore
}

func part1(burrows [4]burrow) int {
	return solve(burrows)
}

func part2(burrows [4]burrow) int {
	insert := [4][]string{
		{"D", "D"},
		{"C", "B"},
		{"B", "A"},
		{"A", "C"},
	}
	newBurrows := [4]burrow{}
	for i, _ := range burrows {
		newBurrows[i] = []string{}
		newBurrows[i] = append(newBurrows[i], burrows[i][0])
		newBurrows[i] = append(newBurrows[i], insert[i]...)
		newBurrows[i] = append(newBurrows[i], burrows[i][1])
	}
	fmt.Println(newBurrows)
	return solve(newBurrows)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	burrows := parseData(data)

	p1 := part1(burrows)
	fmt.Println("Part 1:", p1)

	p2 := part2(burrows)
	fmt.Println("Part 2:", p2)
}
