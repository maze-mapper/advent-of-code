// Advent of Code 2018 - Day 15
package day15

import (
	"container/list"
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strconv"
	"strings"
)

// coord holds grid coordinates
type coord [2]int

// Define constants for unit type
const (
	Elf = iota
	Goblin
)

// unit holds information for an Elf or a Goblin
type unit struct {
	hp, attack, unitType int
}

// parseData reads the input text file and returns data structures
func parseData(file string) ([][]rune, map[coord]*unit) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	area := make([][]rune, len(lines))
	units := map[coord]*unit{}

	for i, line := range lines {
		area[i] = make([]rune, len(line))
		for j := 0; j < len(line); j++ {
			c := coord{j, i}
			switch line[j] {
			case 'E':
				area[i][j] = '.'
				units[c] = &unit{
					hp:       200,
					attack:   3,
					unitType: Elf,
				}
			case 'G':
				area[i][j] = '.'
				units[c] = &unit{
					hp:       200,
					attack:   3,
					unitType: Goblin,
				}
			default:
				area[i][j] = rune(line[j])
			}
		}
	}

	return area, units
}

// sortCoords sorts a slice of coordinates
func sortCoords(coords []coord) {
	sort.Slice(
		coords,
		func(i, j int) bool {
			if coords[i][1] == coords[j][1] {
				// Sort by x if y is the same
				return coords[i][0] < coords[j][0]
			} else {
				// Sort by y
				return coords[i][1] < coords[j][1]
			}
		},
	)
}

// getMoveOrder returns the coordinates of the units in the order they will move
func getMoveOrder(units map[coord]*unit) []coord {
	// Get the coordinates of all units
	coords := make([]coord, len(units))
	i := 0
	for c := range units {
		coords[i] = c
		i++
	}
	// Sort the coordinates
	sortCoords(coords)
	return coords
}

// getTargets returns all units of a different type
func getTargets(units map[coord]*unit, unitType int) []coord {
	coords := []coord{}
	for k, v := range units {
		if v.unitType != unitType {
			coords = append(coords, k)
		}
	}
	return coords
}

// getAdjacentCoordinates returns the adjacent coordinates in the cardinal directions
func getAdjacentCoordinates(area [][]rune, c coord) []coord {
	coords := []coord{}
	// Add coordinates in tie-break order
	// Coordinate above
	if c[1] != 0 {
		coords = append(coords, coord{c[0], c[1] - 1})
	}
	// Coordinate to the left
	if c[0] != 0 {
		coords = append(coords, coord{c[0] - 1, c[1]})
	}
	// Coordinate to the right
	if c[0] != len(area[c[1]]) {
		coords = append(coords, coord{c[0] + 1, c[1]})
	}
	// Coordinate below
	if c[1] != len(area) {
		coords = append(coords, coord{c[0], c[1] + 1})
	}
	return coords
}

// getInRange finds all in range coordinates next to the input coordinates
func getInRange(area [][]rune, units map[coord]*unit, unitType int, targets []coord) []coord {
	inRange := []coord{}
	for _, c := range targets {
		// Get adjacent coordinates
		cAdj := getAdjacentCoordinates(area, c)

		for _, d := range cAdj {
			// Check if the coordinate is empty
			if area[d[1]][d[0]] == '.' {
				inRange = append(inRange, d)
			}
		}
	}

	return inRange
}

// getAttackTarget finds the adjacent enemy with the lowest HP
func getAttackTarget(area [][]rune, units map[coord]*unit, c coord) (coord, bool) {
	minTargetHP := int(^uint(0)>>1) - 1 // initialise to max value
	target := coord{}
	foundTarget := false

	adjacent := getAdjacentCoordinates(area, c)
	for _, a := range adjacent {
		if u, ok := units[a]; ok {
			// Check the unit found is an enemy and see if their HP is lower than the min found
			if u.unitType != units[c].unitType && u.hp < minTargetHP {
				target = a
				foundTarget = true
				minTargetHP = u.hp
			}
		}
	}
	return target, foundTarget
}

// Node is used in the BFS algorithm
type Node struct {
	c          coord
	pathLength int
}

// bfs implements a BFS algorithm to find the shortest path to spaces in range on an enemy
func bfs(area [][]rune, units map[coord]*unit, c coord, targets []coord) (coord, bool) {
	inRange := getInRange(area, units, units[c].unitType, targets)

	// Initialise queue with current coordinate
	queue := list.New()
	node := Node{c: c, pathLength: 0}
	queue.PushBack(node)

	// Track visited nodes and add root node
	visited := map[coord]coord{c: c}

	end := []coord{}
	foundPath := false
	shortestPathFound := 0

	for queue.Len() > 0 {
		nextElement := queue.Front()
		nextNode := nextElement.Value.(Node)

		// If we have already found a path to an in range location, only check nodes at this same distance
		if foundPath && nextNode.pathLength > shortestPathFound {
			break
		}

		// Check if we have reached our target
		for _, t := range inRange {
			if nextNode.c == t {
				end = append(end, nextNode.c)
				foundPath = true
				shortestPathFound = nextNode.pathLength
			}
		}

		adjacent := getAdjacentCoordinates(area, nextNode.c)
		for _, a := range adjacent {
			// Must be an empty space
			if area[a[1]][a[0]] != '.' {
				continue
			}
			// Cannot have another unit on it
			if _, ok := units[a]; ok {
				continue
			}
			// Don't backtrack
			if _, ok := visited[a]; ok {
				continue
			}

			// Add node to queue and mark as visited
			visited[a] = nextNode.c

			aNode := Node{c: a, pathLength: nextNode.pathLength + 1}
			queue.PushBack(aNode)

		}
		queue.Remove(nextElement)
	}

	if foundPath {
		// Choose the in range coordinate which is first in reading order
		sortCoords(end)
		chosen := end[0]

		// Backtrack up the path taken to find the first node on the path
		for {
			if previousNode := visited[chosen]; previousNode == c {
				return chosen, foundPath
			} else {
				chosen = previousNode
			}
		}
	}
	return c, foundPath
}

// doRound carries out all actions in a round
func doRound(area [][]rune, units map[coord]*unit) (bool, bool) {
	roundComplete := false
	battleOver := false
	// Keep track of original positions of fallen units
	fallenUnits := []coord{}

	moveOrder := getMoveOrder(units)
turn:
	for _, c := range moveOrder {
		// Check that unit at this starting coordinate has not already fallen
		// Need to do this as other units may move in to a space previously occupied by a unit
		for _, f := range fallenUnits {
			if c == f {
				continue turn
			}
		}

		// Find targets and check if the battle is over
		targets := getTargets(units, units[c].unitType)
		if len(targets) == 0 {
			battleOver = true
			return battleOver, roundComplete
		}

		// Move if unit is not already adjacent to an enemy and has a clear path to an enemy
		if nextStep, foundPath := bfs(area, units, c, targets); foundPath && nextStep != c {
			u := units[c]
			delete(units, c)
			units[nextStep] = u
			c = nextStep
		}

		// Attack
		if target, found := getAttackTarget(area, units, c); found {
			units[target].hp -= units[c].attack
			if units[target].hp <= 0 {
				delete(units, target)
				// Mark location where unit fell
				fallenUnits = append(fallenUnits, target)
			}
		}

	}
	roundComplete = true
	return battleOver, roundComplete
}

// determineOutcome returns the outcome of a battle
func determineOutcome(rounds int, units map[coord]*unit) int {
	totalHP := 0
	for _, u := range units {
		totalHP += u.hp
	}
	return totalHP * rounds
}

// doCombat will perform the combat in an area for the given units
func doCombat(area [][]rune, units map[coord]*unit) int {
	round := 0
	combatOver := false
	for !combatOver {
		roundComplete := false
		combatOver, roundComplete = doRound(area, units)
		if roundComplete {
			round += 1
		}
	}
	outcome := determineOutcome(round, units)
	return outcome
}

// copyUnits deep copies a set of units
func copyUnits(units map[coord]*unit) map[coord]*unit {
	newUnits := map[coord]*unit{}
	for k, v := range units {
		newUnits[k] = &unit{hp: v.hp, attack: v.attack, unitType: v.unitType}
	}
	return newUnits
}

// setElfAttackPower updates the attack power of elves
func setElfAttackPower(units map[coord]*unit, attack int) {
	for _, v := range units {
		if v.unitType == Elf {
			v.attack = attack
		}
	}
}

// countUnits counts the number of units of the given type
func countUnits(units map[coord]*unit, unitType int) int {
	count := 0
	for _, v := range units {
		if v.unitType == unitType {
			count += 1
		}
	}
	return count
}

// printState prints the current game state for help with debugging
func printState(area [][]rune, units map[coord]*unit) {
	for i, row := range area {
		endStr := " "
		for j, val := range row {
			c := coord{j, i}
			if u, ok := units[c]; ok {
				switch u.unitType {
				case Elf:
					fmt.Print("E")
					endStr += "E(" + strconv.Itoa(u.hp) + ") "
				case Goblin:
					fmt.Print("G")
					endStr += "G(" + strconv.Itoa(u.hp) + ") "
				default:
					fmt.Print(" ")
				}
			} else {
				fmt.Print(string(val))
			}
		}
		fmt.Print(endStr)
		fmt.Print("\n")
	}
}

func Run(inputFile string) {
	area, originalUnits := parseData(inputFile)

	// Part 1
	units := copyUnits(originalUnits)
	outcome := doCombat(area, units)
	fmt.Println("Part 1:", outcome)

	// Part 2
	totalElves := countUnits(originalUnits, Elf)
	for elfAttack := 4; elfAttack < 201; elfAttack++ {
		units := copyUnits(originalUnits)
		setElfAttackPower(units, elfAttack)
		outcome := doCombat(area, units)
		if elves := countUnits(units, Elf); elves == totalElves {
			fmt.Println("Part 2:", outcome)
			break
		}
	}
}
