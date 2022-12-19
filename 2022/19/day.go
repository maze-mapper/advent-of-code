// Advent of Code 2022 - Day 19.
package day19

import (
	//	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
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

// blueprint contains the resource requirements for each robot type.
type blueprint [4][3]int

func (bp blueprint) buildChoices(resources [4]int, robots [4]int) [][4]int {
	buildable := [][4]int{
		[4]int{}, // Build nothing.
	}
	//	maxRates := bp.maxResourceRates()
	//	fmt.Println(maxRates)
	for robotIndex, requiredResources := range bp {
		// Don't build a robot of this type if we have reached the maximum rate at which those resources could be used up.
		//		if maxRates[robotIndex] == robots[robotIndex] {
		//			continue
		//		}

		// Check which robots can be built with the available resources.
		b := true
		for i, cost := range requiredResources {
			if cost > resources[i] {
				b = false
				break
			}
		}
		if b {
			v := [4]int{}
			v[robotIndex] = 1
			buildable = append(buildable, v)
		}
	}
	return buildable
}

// maxResourceRates returns the maximum cost for each material for any robot.
func (bp blueprint) maxResourceRates() [4]int {
	maxRates := [4]int{}
	for _, requiredResources := range bp {
		for i, cost := range requiredResources {
			if cost > maxRates[i] {
				maxRates[i] = cost
			}
		}
	}
	return maxRates
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
			[3]int{oreRobotOreCost, 0, 0},                          // Ore robot.
			[3]int{clayRobotOreCost, 0, 0},                         // Clay robot.
			[3]int{obsidianRobotOreCost, obsidianRobotClayCost, 0}, // Obsidian robot.
			[3]int{geodeRobotOreCost, 0, geodeRobotObsideanCost},   // Geode robot.
		}
		blueprints[i] = bp
	}
	return blueprints
}

type Node struct {
	robots, resources [4]int
	minute            int
	priority          int
	index             int
}

func (n *Node) state() string {
	var b strings.Builder
	for _, r := range n.robots {
		b.WriteString(strconv.Itoa(r))
		b.WriteString(" ")
	}
	for _, r := range n.resources {
		b.WriteString(strconv.Itoa(r))
		b.WriteString(" ")
	}
	return b.String()
}

// A PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the lowest, not highest, priority so we use less than here.
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x interface{}) {
	n := len(*pq)
	node := x.(*Node)
	node.index = n
	*pq = append(*pq, node)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	node := old[n-1]
	old[n-1] = nil  // avoid memory leak
	node.index = -1 // for safety
	*pq = old[0 : n-1]
	return node
}

// geodeUpperBound returns an upper bound for the total number of geodes that could be produced.
// It assumes that one new geode-collecting robot is produced every turn for all remaining turns.
// This is an arithmetic sum and is added to the current number of geodes collected.
func geodeUpperBound(robots [4]int, resources [4]int, remainingTurns int) int {
	return resources[indexGeode] + remainingTurns*(robots[indexGeode]+robots[indexGeode]+remainingTurns)/2
}

func dfs(bp blueprint, robots [4]int, resources [4]int, remainingTurns int, maxGeodes *int) {
	//	fmt.Println(robots, resources, remainingTurns, *maxGeodes)
	if remainingTurns == 0 {
		if resources[indexGeode] > *maxGeodes {
			*maxGeodes = resources[indexGeode]
		}
		//		fmt.Println(robots, resources, remainingTurns, *maxGeodes)
		return
	}

	choices := bp.buildChoices(resources, robots)
	//	fmt.Println(robots, resources, remainingTurns, *maxGeodes, choices)
	for c := len(choices) - 1; c >= 0; c-- {
		//	for _, robotsToBuild := range bp.buildChoices(resources) {
		robotsToBuild := choices[c]
		newRobots := robots
		newResources := resources

		// Spend resources to build a new robot.
		for i, n := range robotsToBuild {
			for j, cost := range bp[i] {
				newResources[j] -= cost * n
			}
		}

		// Collect resources.
		for i, n := range robots {
			newResources[i] += n
		}

		// Finish building new robot.
		for i, n := range robotsToBuild {
			newRobots[i] += n
		}

		if geodeUpperBound(newRobots, newResources, remainingTurns-1) <= *maxGeodes {
			continue
		}

		dfs(bp, newRobots, newResources, remainingTurns-1, maxGeodes)
	}
}

func runBlueprint(bp blueprint, minutes int) int {
	maxGeodes := 0
	dfs(bp, [4]int{1, 0, 0, 0}, [4]int{}, minutes, &maxGeodes)
	fmt.Println("DONE:", maxGeodes)
	return maxGeodes
}

/*
func runBlueprint(bp blueprint, minutes int) int {
	// Initialise one ore-collecting robot.
	robots := [4]int{}
	robots[indexOre] = 1

	pq := make(PriorityQueue, 1)
	start := &Node{
		robots: robots,
		minute: 0,
	}
	pq[0] = start
	heap.Init(&pq)

	visited := map[string]int{
		start.state(): start.resources[indexGeode],
	}

	maxGeodes := 0

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

//		fmt.Println(node, pq.Len())

		if node.minute >= minutes {
			if node.resources[indexGeode] > maxGeodes {
				maxGeodes = node.resources[indexGeode]
			}
			continue
		}

		for _, robotsToBuild := range bp.buildChoices(node.resources) {
			newResources := node.resources

			// Spend resources to build a new robot.
			for i, num := range robotsToBuild {
				for j, cost := range bp[i] {
					newResources[j] -= cost * num
				}
			}

			// Collect resources.
			for j, n := range node.robots {
				newResources[j] += n
			}

			// Finish building new robot.
			newRobots := node.robots
			for i, num := range robotsToBuild {
				newRobots[i] += num
			}

			newMinute := node.minute + 1

			h := (minutes - node.minute - 1) / (1 + newRobots[indexGeode])

			newNode := &Node{
				robots:    newRobots,
				resources: newResources,
				minute:    newMinute,
				priority:  newMinute + h,
			}

			if geodeUpperBound(newNode.robots, newNode.resources, minutes - newMinute) <= maxGeodes {
				continue
			}

			if geodes, ok := visited[newNode.state()]; !ok || newResources[indexGeode] > geodes {
				heap.Push(&pq, newNode)
				visited[newNode.state()] = newResources[indexGeode]
			}
		}

	}

	return maxGeodes
}*/

func part1(blueprints []blueprint) int {
	results := make([]int, len(blueprints))
	var wg sync.WaitGroup
	for i, bp := range blueprints {
		wg.Add(1)
		i := i
		bp := bp
		go func() {
			results[i] = runBlueprint(bp, 24)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(results)

	ql := 0
	for i, geodes := range results {
		ql += (i + 1) * geodes
	}

	return ql
}

func part2(blueprints []blueprint) int {
	bpCount := 3
	results := make([]int, bpCount)
	var wg sync.WaitGroup
	for i, bp := range blueprints[:bpCount] {
		wg.Add(1)
		i := i
		bp := bp
		go func() {
			results[i] = runBlueprint(bp, 32)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(results)

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

	//	p1 := part1(blueprints)
	//	fmt.Println("Part 1:", p1)

	p2 := part2(blueprints)
	fmt.Println("Part 2:", p2)
}
