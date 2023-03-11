// Advent of Code 2022 - Day 19.
package day19

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
	"sync"
)

// Constant indices for each resource type.
const (
	indexOre = iota
	indexClay
	indexObsidian
	indexGeode
)

// state holds information on the current state
type state struct {
	remainingTurns    int
	robots, resources [4]int
	bp                blueprint
}

// canBuild returns true if the provided robot type can be built with the available resources.
func (s *state) canBuild(robotIndex int) bool {
	for i, cost := range s.bp.resourceRequirements[robotIndex] {
		if cost > s.resources[i] {
			return false
		}
	}
	return true
}

// buildRobot consumes resources to build a robot of the specified type.
// It returns a callback function to complete building the robot.
func (s *state) buildRobot(robotIndex int) func() {
	for i, cost := range s.bp.resourceRequirements[robotIndex] {
		s.resources[i] -= cost
	}
	return func() {
		s.robots[robotIndex] += 1
	}
}

// collectResources increases the resources by the number of robots for each resource.
func (s *state) collectResources() {
	for i, n := range s.robots {
		s.resources[i] += n
	}
}

// geodeUpperBound returns an upper bound for the total number of geodes that could be produced.
// It assumes that one new geode-collecting robot is produced every turn for all remaining turns.
// This is an arithmetic sum and is added to the current number of geodes collected.
//
// SUM = n/2 * (2a +(n-1)d)
func (s *state) geodeUpperBound() int {
	return s.resources[indexGeode] + s.remainingTurns*(2*s.robots[indexGeode]+s.remainingTurns-1)/2
}

// newState returns a new state.
func newState(turns int, bp blueprint) state {
	return state{
		remainingTurns: turns,
		robots:         [4]int{1, 0, 0, 0},
		resources:      [4]int{0, 0, 0, 0},
		bp:             bp,
	}
}

// blueprint contains the resource requirements for each robot type.
type blueprint struct {
	resourceRequirements [4][3]int
	maxRates             [4]int
}

// maxResourceRates returns the maximum cost for each material for any robot.
func (bp *blueprint) maxResourceRates() [4]int {
	if bp.maxRates == [4]int{} {
		maxRates := [4]int{}
		for _, requiredResources := range bp.resourceRequirements {
			for i, cost := range requiredResources {
				if cost > maxRates[i] {
					maxRates[i] = cost
				}
			}
		}
		maxRates[3] = 50 // hack
		bp.maxRates = maxRates
	}
	return bp.maxRates
}

func parseInput(data []byte) []blueprint {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	blueprints := make([]blueprint, len(lines))
	pattern := "Blueprint %d: Each ore robot costs %d ore. Each clay robot costs %d ore. Each obsidian robot costs %d ore and %d clay. Each geode robot costs %d ore and %d obsidian."
	for i, line := range lines {
		var j, oreRobotOreCost, clayRobotOreCost, obsidianRobotOreCost, obsidianRobotClayCost, geodeRobotOreCost, geodeRobotObsideanCost int
		n, err := fmt.Sscanf(line, pattern, &j, &oreRobotOreCost, &clayRobotOreCost, &obsidianRobotOreCost, &obsidianRobotClayCost, &geodeRobotOreCost, &geodeRobotObsideanCost)
		if err != nil {
			log.Fatal(err)
		}
		if n != 7 {
			log.Fatal("Not enough matches found")
		}
		bp := blueprint{
			resourceRequirements: [4][3]int{
				[3]int{oreRobotOreCost, 0, 0},                          // Ore robot.
				[3]int{clayRobotOreCost, 0, 0},                         // Clay robot.
				[3]int{obsidianRobotOreCost, obsidianRobotClayCost, 0}, // Obsidian robot.
				[3]int{geodeRobotOreCost, 0, geodeRobotObsideanCost},   // Geode robot.
			},
		}
		blueprints[i] = bp
	}
	return blueprints
}

// tryBuild attempts to build a specified robot on a local copy of the state.
func tryBuild(s state, robotIndex int, maxGeodes *int) {
	if s.robots[robotIndex] >= s.bp.maxResourceRates()[robotIndex] {
		return
	}
	if !s.canBuild(robotIndex) {
		return
	}
	cb := s.buildRobot(robotIndex)
	s.collectResources()
	cb()
	dfs(s, maxGeodes)
}

func dfs(s state, maxGeodes *int) {
	if s.remainingTurns == 0 {
		if s.resources[indexGeode] > *maxGeodes {
			*maxGeodes = s.resources[indexGeode]
		}
		return
	}

	if s.geodeUpperBound() <= *maxGeodes {
		return
	}

	s.remainingTurns -= 1
	for robotIndex := len(s.robots) - 1; robotIndex >= 0; robotIndex-- {
		tryBuild(s, robotIndex, maxGeodes)
	}
	s.collectResources()

	dfs(s, maxGeodes)
}

// runBlueprint runs the provided blueprint for the given number of iterations.
// It returns the maximum number of geodes the blueprint can produce.
func runBlueprint(bp blueprint, minutes int) int {
	maxGeodes := 0
	s := newState(minutes, bp)
	dfs(s, &maxGeodes)
	fmt.Println("DONE:", maxGeodes)
	return maxGeodes
}

// runBlueprints concurrently runs the provided blueprints for the given number of iterations.
// It returns a slice containing the maximum number of geodes each blueprint can produce.
func runBlueprints(blueprints []blueprint, minutes int) []int {
	results := make([]int, len(blueprints))
	var wg sync.WaitGroup
	for i, bp := range blueprints {
		wg.Add(1)
		i := i
		bp := bp
		go func() {
			results[i] = runBlueprint(bp, minutes)
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(results)
	return results
}

func part1(blueprints []blueprint) int {
	results := runBlueprints(blueprints, 24)
	ql := 0
	for i, geodes := range results {
		ql += (i + 1) * geodes
	}
	return ql
}

func part2(blueprints []blueprint) int {
	results := runBlueprints(blueprints[:3], 32)
	out := 1
	for _, geodes := range results {
		out *= geodes
	}
	return out
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	blueprints := parseInput(data)

	p1 := part1(blueprints)
	fmt.Println("Part 1:", p1)

	p2 := part2(blueprints)
	fmt.Println("Part 2:", p2)
}
