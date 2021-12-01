// Advent of Code 2015 - Day 199
// Note: Part 2 may be better done with grammar, for the input an equation can give the number of steps
package day19

import (
	"container/heap"
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

type Replacement struct {
	before, after string
}

// parseInput parses the input data to extract the allowed replacements and the target molecule
func parseInput(data []byte) ([]Replacement, string) {
	parts := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n\n",
	)
	lines := strings.Split(parts[0], "\n")
	replacements := []Replacement{}
	for _, replacement := range lines {
		molecules := strings.Split(replacement, " => ")
		replacements = append(
			replacements,
			Replacement{before: molecules[0], after: molecules[1]},
		)
	}
	return replacements, parts[1]
}

// ReplaceNth will return a copy of the string s where the n-th occurrence of old is replaced by new
func replaceNth(s, old, new string, n int) string {
	i := 0
	for m := 1; m <= n; m++ {
		idx := strings.Index(s[i:], old)
		if idx < 0 {
			break
		}
		i += idx
		if m == n {
			return s[:i] + new + s[i+len(old):]
		}
		i += len(old)
	}
	return s
}

func part1(replacements []Replacement, targetMolecule string) {
	distinctMolecules := make(map[string]int)
	for _, replacement := range replacements {
		count := strings.Count(targetMolecule, replacement.before)
		for i := 1; i <= count; i++ {
			newMolecule := replaceNth(targetMolecule, replacement.before, replacement.after, i)
			//			newMolecule := strings.Replace(targetMolecule, replacement.before, replacement.after, i) // NO, INCORRECT
			distinctMolecules[newMolecule] = 1
		}
	}
	fmt.Println("Part 1:", len(distinctMolecules), "distinct molecules can be made from the medicine molecule")
}

// previousDistinctMolecules finds all the distinct molecules that could have produced the current molecule
func previousDistinctMolecules(replacements []Replacement, currentMolecule string) []string {
	distinctMolecules := make(map[string]int)
	for _, replacement := range replacements {
		count := strings.Count(currentMolecule, replacement.after)
		for i := 1; i <= count; i++ {
			previousMolecule := replaceNth(currentMolecule, replacement.after, replacement.before, i)
			//previousMolecule := strings.Replace(currentMolecule, replacement.after, replacement.before, 1)
			distinctMolecules[previousMolecule] = 1
		}
	}

	molecules := make([]string, 0, len(distinctMolecules))
	for k := range distinctMolecules {
		molecules = append(molecules, k)
	}
	return molecules
}

func countUppercase(s string) int {
	count := 0
	for _, c := range s {
		if c >= 65 && c <= 90 {
			count += 1
		}
	}
	return count
}

type Item struct {
	value    string
	priority int // The priority of the item in the queue.
	steps    int // The number of steps taken
	// The index is needed by update and is maintained by the heap.Interface methods.
	index int // The index of the item in the heap.
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

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
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
/*func (pq *PriorityQueue) update(item *Item, value string, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}*/
/*
func (pq *PriorityQueue) updateOrPush(item *Item) {
	for _, it := range *pq {
		if item.value == it.value && item.priority < it.priority {
			it.priority = item.priority
			heap.Fix(pq, it.index)
			return
		}
	}
	heap.Push(pq, item)
}*/

func part2(replacements []Replacement, medicineMolecule string) {
	// Use an A* approach to work backwards from the medicine molecule to the single electron "e"
	// We work in this direction as the heuristic is simpler: a smaller molecule is better

	visted := map[string]int{medicineMolecule: 0}
	//	steps := 0

	// Initialise priority queue
	pq := make(PriorityQueue, 1)
	pq[0] = &Item{
		value:    medicineMolecule,
		steps:    0,
		priority: 1, // Will be the first item removed regardless of priority
		index:    0,
	}
	heap.Init(&pq)

	for pq.Len() > 0 {
		// Take next item off priority queue
		item := heap.Pop(&pq).(*Item)
		fmt.Println(item)

		// Check if we have arrived at the electron
		if item.value == "e" {
			fmt.Println("Part 2:", item.steps, "steps")
			return
		}

		steps := item.steps + 1

		// Find all previous distinct molecules and add any unvisted ones to the priority queue
		previousMolecules := previousDistinctMolecules(replacements, item.value)
		for _, molucule := range previousMolecules {
			if s, ok := visted[molucule]; !ok || steps < s { // TODO
				it := Item{
					value:    molucule,
					steps:    steps,
					priority: steps + 2*countUppercase(molucule),
				}
				heap.Push(&pq, &it)
				visted[molucule] = steps
			}
		}
	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	replacements, targetMolecule := parseInput(data)

	part1(replacements, targetMolecule)
	part2(replacements, targetMolecule)
}
