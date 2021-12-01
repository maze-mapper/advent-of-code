// Advent of Code 2015 - Day 144
package day14

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var lineRegExp = regexp.MustCompile(`([a-zA-Z]+) can fly ([0-9]+) km/s for ([0-9]+) seconds, but then must rest for ([0-9]+) seconds.`)

type Reindeer struct {
	name                             string
	speed, flyDuration, restDuration int
}

type ScoredReindeer struct {
	name                                                     string
	speed, flyDuration, restDuration, distance, timer, score int
	flying                                                   bool
}

func parseInput(input []byte) []Reindeer {
	lines := strings.Split(
		strings.TrimSuffix(string(input), "\n"), "\n",
	)

	herd := make([]Reindeer, len(lines))
	for i, line := range lines {
		matches := lineRegExp.FindStringSubmatch(line)
		reindeer := Reindeer{name: matches[1]}

		// Map the index of the match to the Reindeer attribute
		pairings := map[int]*int{
			2: &reindeer.speed,
			3: &reindeer.flyDuration,
			4: &reindeer.restDuration,
		}
		for j, attr := range pairings {
			val, err := strconv.Atoi(matches[j])
			if err != nil {
				log.Fatal(err)
			}
			*attr = val
		}
		herd[i] = reindeer
	}
	return herd
}

// distanceTravelled returns the distance a reindeer would travel in a given time t
func (reindeer *Reindeer) distanceTravelled(t int) int {
	cycleDuration := reindeer.flyDuration + reindeer.restDuration
	cyclesCompleted := t / cycleDuration

	distance := cyclesCompleted * reindeer.speed * reindeer.flyDuration

	// Remaining time for non-complete cycles
	t = t % cycleDuration

	if reindeer.flyDuration < t { // Reindeer stops flying before the time expires
		distance += reindeer.speed * reindeer.flyDuration
	} else { // Reindeer flys up until the time expires
		distance += reindeer.speed * t
	}

	return distance
}

// makeScored returns a ScoredReindeer struct
func (reindeer *Reindeer) makeScored() ScoredReindeer {
	r := ScoredReindeer{
		name:         reindeer.name,
		speed:        reindeer.speed,
		flyDuration:  reindeer.flyDuration,
		restDuration: reindeer.restDuration,
		distance:     0,
		timer:        reindeer.flyDuration,
		score:        0,
		flying:       true,
	}
	return r
}

func part1(herd []Reindeer, t int) {
	maxDistance := -int(^uint(0)>>1) - 1 // initialise to min value

	for _, reindeer := range herd {
		if d := reindeer.distanceTravelled(t); d > maxDistance {
			maxDistance = d
		}
	}

	fmt.Println("Part 1: The winning reindeer has travelled", maxDistance, "km")
}

func part2(herd []Reindeer, t int) {
	// Convert Reindeer to ScoredReindeer to store additonal state information
	scoredHerd := make([]ScoredReindeer, len(herd))
	for i, reindeer := range herd {
		scoredHerd[i] = reindeer.makeScored()
	}

	for t > 0 {
		bestDistance := 0
		for i := 0; i < len(scoredHerd); i++ {
			reindeer := &scoredHerd[i]
			// Increase distance travelled
			if reindeer.flying {
				reindeer.distance += reindeer.speed
			}

			// Update the best distance
			if reindeer.distance > bestDistance {
				bestDistance = reindeer.distance
			}

			// Switch between resting and flying
			reindeer.timer -= 1
			if reindeer.timer == 0 {
				if reindeer.flying {
					reindeer.timer = reindeer.restDuration
				} else {
					reindeer.timer = reindeer.flyDuration
				}
				reindeer.flying = !reindeer.flying
			}
		}

		// Update scores
		for i := 0; i < len(scoredHerd); i++ {
			reindeer := &scoredHerd[i]
			if reindeer.distance == bestDistance {
				reindeer.score += 1
			}
		}

		t -= 1
	}

	bestScore := 0
	for _, reindeer := range scoredHerd {
		if reindeer.score > bestScore {
			bestScore = reindeer.score
		}
	}
	fmt.Println("Part 2: The winning reindeer scores", bestScore, "points")
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	reindeer := parseInput(data)

	t := 2503
	part1(reindeer, t)
	part2(reindeer, t)
}
