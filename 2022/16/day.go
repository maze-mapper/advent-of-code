// Advent of Code 2022 - Day 16.
package day16

import (
	"container/heap"
        "fmt"
        "io/ioutil"
        "log"
	"sort"
	"strconv"
        "strings"
)

type valve struct {
	flowRate int
	leadsTo []string
}

func parseInput(data []byte) map[string]valve {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	valves := make(map[string]valve, len(lines))
	for _, line := range lines {
		line = strings.Replace(line, "=", " ", 1)
		line = strings.Replace(line, ";", "", 1)
		line = strings.ReplaceAll(line, ",", "")
		tokens := strings.Split(line, " ")
		
		flowRate, err := strconv.Atoi(tokens[5])
		if err != nil {
			log.Fatal(err)
		}
		valves[tokens[1]] = valve{
			flowRate: flowRate,
			leadsTo: tokens[10:],
		}
	}
	return valves
}

type Node struct {
	name string
	elephant string
	minute int
	openValves map[string]bool
	currentVented int
	priority int
	index int
}

// A PriorityQueue implements heap.Interface and holds Nodes.
type PriorityQueue []*Node

func (pq PriorityQueue) Len() int {
        return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
        // We want Pop to give us the highest, not lowest, priority so we use less than here.
        return pq[i].priority > pq[j].priority
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

func totalFlowRate(openValves map[string]bool, valves map[string]valve) int {
	flowRate := 0
	for v := range openValves {
		flowRate += valves[v].flowRate
	}
	return flowRate
}

func potentialFlowRate(openValves map[string]bool, valves map[string]valve) int {
	flowRate := 0
	for name, v := range valves {
		if _, ok := openValves[name]; !ok {
			flowRate += v.flowRate
		}
	}
	return flowRate
}

func stringState(node *Node) string {
	var b strings.Builder

	positions := []string{node.name, node.elephant}
	sort.Strings(positions)
//	b.WriteString(string(node.minute))
//	b.WriteString(node.name)
//	b.WriteString(node.elephant)
	for _, p := range positions {
		b.WriteString(p)
	}

	valves := []string{}
	for v := range node.openValves {
		valves = append(valves, v)
	}
	sort.Strings(valves)
	for _, v := range valves {
		b.WriteString(v)
	}
	return b.String()
}

func aStar(valves map[string]valve, timeLimit int) int {
	pq := make(PriorityQueue, 1)
	startNode := &Node{
		name: "AA",
		minute: 1,
		openValves: map[string]bool{},
	}
	pq[0] = startNode
	heap.Init(&pq)

	visitedStates := map[string]int{}

	maxFlowRate := 0
	for _, v := range valves {
		maxFlowRate += v.flowRate
	}

	maxVented := 0
	openableValves := 0
        for _, v := range valves {
                if v.flowRate != 0 {
                        openableValves += 1
                }
        }

	for pq.Len() > 0 {
		node := heap.Pop(&pq).(*Node)

		state := stringState(node)
                if vented, ok := visitedStates[state]; ok && vented <= node.currentVented {
                        continue
                }
                visitedStates[state] = node.currentVented

//		fmt.Println(node.minute, node.name, node.priority, node.currentVented, node.openValves)

		if node.minute >= timeLimit {
			if node.currentVented > maxVented {
				fmt.Println("TIME", node.currentVented)
				maxVented = node.currentVented
			}
                        continue
                }

		if len(node.openValves) == openableValves {
			vented := node.currentVented + (timeLimit - node.minute) * totalFlowRate(node.openValves, valves)
			if vented > maxVented {
				maxVented = vented
				fmt.Println("OPEN", node.currentVented)
			}
			continue
		}

		// Open current valve if it is not open and does not have a zero flow rate.
		if _, ok := node.openValves[node.name]; !ok && valves[node.name].flowRate != 0 {
			openValves := map[string]bool{}
			for k, v := range node.openValves {
				openValves[k] = v
			}
			openValves[node.name] = true
			fr := totalFlowRate(openValves, valves)
			m := node.minute + 1
			pot := node.currentVented + fr + (timeLimit - m) * maxFlowRate
//                        pri := node.currentVented + fr
                        if pot < maxVented {
                                continue
                        }
			pri := node.currentVented + fr + (timeLimit - m) * maxFlowRate// (fr + potentialFlowRate(openValves, valves) / 2)
			nextNode := Node{
				name: node.name,
				minute: m,
				openValves: openValves,
				currentVented: node.currentVented + fr,
				priority: pri,
//				priority: fr / m,
//				priority: (node.currentVented + fr) / m + (fr + potentialFlowRate(openValves, valves) * (timeLimit - m)),
			}
			heap.Push(&pq, &nextNode)
		}

		// Move to other valves.
		for _, v := range valves[node.name].leadsTo {
			openValves := map[string]bool{}
                        for k, v := range node.openValves {
                                openValves[k] = v
                        }
			fr := totalFlowRate(openValves, valves)
			m := node.minute + 1
			pot := node.currentVented + fr + (timeLimit - m) * maxFlowRate
//			pri := node.currentVented + fr
                        if pot < maxVented {
                                continue
                        }
			pri := node.currentVented + fr + (timeLimit - m) * maxFlowRate//(fr + potentialFlowRate(openValves, valves) / 2)
			// Increase priority of moving to an unopened valve.
			if _, ok := openValves[v]; !ok {
				pri += valves[v].flowRate
			}
			nextNode := Node{
				name: v,
                                minute: m,
                                openValves: openValves,
				currentVented: node.currentVented + fr,
				priority: pri,
//				priority: (node.currentVented + fr) / m + (fr + potentialFlowRate(openValves, valves) * (timeLimit - m)),
//                                priority: (fr / m) + (potentialFlowRate(openValves, valves) / (timeLimit - m)),
                        }
			heap.Push(&pq, &nextNode)
		}

	}

	return maxVented
}

func part1(valves map[string]valve) int {
	return aStar(valves, 30)
}

func part2(valves map[string]valve) int {
	timeLimit := 26

	pq := make(PriorityQueue, 1)
        startNode := &Node{
                name: "AA",
		elephant: "AA",
                minute: 1,
                openValves: map[string]bool{},
        }
	pq[0] = startNode
        heap.Init(&pq)

        visitedStates := map[string]int{
		stringState(startNode): 0,
	}

        maxFlowRate := 0
        for _, v := range valves {
                maxFlowRate += v.flowRate
        }

        maxVented := 0
	openableValves := 0
	for _, v := range valves {
		if v.flowRate != 0 {
			openableValves += 1
		}
	}

	for pq.Len() > 0 {
                node := heap.Pop(&pq).(*Node)

//                state := stringState(node)
//                if vented, ok := visitedStates[state]; ok && vented <= node.currentVented {
//                        continue
//                }
//                visitedStates[state] = node.currentVented

//	        fmt.Println(node.minute, node.name, node.elephant, node.priority, node.currentVented, node.openValves)

		if node.minute >= timeLimit {
                        if node.currentVented > maxVented {
                                fmt.Println("TIME", node.currentVented)
                                maxVented = node.currentVented
                        }
                        continue
                }

                if len(node.openValves) == openableValves {
                        vented := node.currentVented + (timeLimit - node.minute) * maxFlowRate
                        if vented > maxVented {
                                maxVented = vented
                                fmt.Println("OPEN", node.currentVented)
                        }
                        continue
                }

		youCanOpen := false
		if _, ok := node.openValves[node.name]; !ok && valves[node.name].flowRate != 0 {
			youCanOpen = true
		}
		elephantCanOpen := false
		if _, ok := node.openValves[node.elephant]; !ok && node.name != node.elephant && valves[node.elephant].flowRate != 0 {
                        elephantCanOpen = true
                }

		yourOptions := make([]string, len(valves[node.name].leadsTo))
		copy(yourOptions, valves[node.name].leadsTo)
		if youCanOpen {
			yourOptions = append(yourOptions, "OPEN")
		}
		elephantOptions := make([]string, len(valves[node.elephant].leadsTo))
		copy(elephantOptions, valves[node.elephant].leadsTo)
                if elephantCanOpen {
                        elephantOptions = append(elephantOptions, "OPEN")
                }

		for _, you := range yourOptions {
			for _, el := range elephantOptions {
				openValves := map[string]bool{}
	                        for k, v := range node.openValves {
        	                        openValves[k] = v
                	        }
				youPos := node.name
				elPos := node.elephant

				if you == "OPEN" {
					openValves[node.name] = true
				} else {
					youPos = you
				}
				if el == "OPEN" {
                                        openValves[node.elephant] = true
                                } else {
                                        elPos = el
                                }

				priShift := 0
				if _, ok := openValves[youPos]; !ok {
                                	priShift += valves[youPos].flowRate
                        	}
				if _, ok := openValves[elPos]; !ok && youPos != elPos {
                                        priShift += valves[elPos].flowRate
                                }

				fr := totalFlowRate(openValves, valves)
	                        m := node.minute + 1

				pot := node.currentVented + fr + (timeLimit - m) * maxFlowRate
                	        if pot < maxVented {
        	                        continue
	                        }
                        	pri := node.currentVented + fr + (timeLimit - m) * maxFlowRate

				nextNode := &Node{
                	                name: youPos,
        	                        elephant: elPos,
	                                minute: m,
                                	openValves: openValves,
                        	        currentVented: node.currentVented + fr,
                	                priority: pri + priShift,
        	                }

				state := stringState(nextNode)
		                if vented, ok := visitedStates[state]; ok && vented >= nextNode.currentVented {
	                       		continue
		                } else {
					visitedStates[state] = nextNode.currentVented
		                        heap.Push(&pq, nextNode)
				}
			}
		}

	}


	return maxVented
}

func Run(inputFile string) {
        data, err := ioutil.ReadFile(inputFile)
        if err != nil {
                log.Fatal(err)
        }
        valves := parseInput(data)

        p1 := part1(valves)
        fmt.Println("Part 1:", p1)

        p2 := part2(valves)
        fmt.Println("Part 2:", p2)
}

