// Advent of Code 2018 - Day 4
package day4

import (
	"fmt"
	"io/ioutil"
	"log"
	"sort"
	"strings"
	"time"
)

// TimeLayout is the timestamp format provided by the puzzle input
const TimeLayout = "2006-01-02 15:04"

// Constants representing each type of event
const (
	BeginShift = iota
	FallAsleep
	WakeUp
)

// event holds information
type event struct {
	timestamp time.Time
	guard, id int
}

// parseEvent converts a string in to an event struct
func parseEvent(s string) event {
	s = strings.TrimPrefix(s, "[")
	parts := strings.Split(s, "] ")

	e := event{}
	if t, err := time.Parse(TimeLayout, parts[0]); err == nil {
		e.timestamp = t
	} else {
		log.Fatal(err)
	}

	switch parts[1] {

	case "falls asleep":
		e.id = FallAsleep

	case "wakes up":
		e.id = WakeUp

	default:
		e.id = BeginShift
		if _, err := fmt.Sscanf(parts[1], "Guard #%d begins shift", &e.guard); err != nil {
			log.Fatal(err)
		}
	}

	return e
}

// parseData converts the input data in to a sorted slice of event structs
func parseData(data []byte) []event {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	// Sort so that lines are in chronological order - string sort is sufficient
	sort.Strings(lines)

	events := make([]event, len(lines))
	for i, line := range lines {
		events[i] = parseEvent(line)
	}
	return events
}

// part1 returns the guard and minute for the guard who sleeps the most and the time spent sleeping at this minute is greatest
func part1(sleeping map[int]map[int]int) (int, int) {
	guard := 0
	mostTimeSleeping := 0

	for k, v := range sleeping {
		timeSleeping := 0
		for _, vv := range v {
			timeSleeping += vv
		}
		if timeSleeping > mostTimeSleeping {
			mostTimeSleeping = timeSleeping
			guard = k
		}
	}

	mostTimeSleeping = 0
	var bestMinute int
	for k, v := range sleeping[guard] {
		if v > mostTimeSleeping {
			mostTimeSleeping = v
			bestMinute = k
		}
	}

	return guard, bestMinute
}

// part2 returns the guard and minute where the time spent sleeping at this minute is greatest
func part2(sleeping map[int]map[int]int) (int, int) {
	guard := 0
	mostTimeSleeping := 0

	var bestMinute int
	for k, v := range sleeping {
		for kk, vv := range v {
			if vv > mostTimeSleeping {
				mostTimeSleeping = vv
				guard = k
				bestMinute = kk
			}
		}
	}

	return guard, bestMinute
}

// prepare returns a nested map of guard to minutes with a count of how often they are asleep at that minute
func prepare(events []event) map[int]map[int]int {
	sleeping := map[int]map[int]int{}
	guard := 0

	minute, _ := time.ParseDuration("1m")

	for i, e := range events {
		fmt.Println(e)
		switch e.id {

		case BeginShift:
			guard = e.guard

		case FallAsleep:

		case WakeUp:
			// Initialise map
			if _, ok := sleeping[guard]; !ok {
				sleeping[guard] = map[int]int{}
			}
			// Enter sleeping times
			fmt.Println("Start", events[i-1].timestamp, "End", e.timestamp)
			for startTime := events[i-1].timestamp; startTime.Before(e.timestamp); startTime = startTime.Add(minute) {
				sleeping[guard][startTime.Minute()] += 1
			}
		}
	}

	return sleeping
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	events := parseData(data)
	sleeping := prepare(events)

	p1Guard, p1Minute := part1(sleeping)
	fmt.Println("Part 1:", p1Guard*p1Minute)

	p2Guard, p2Minute := part2(sleeping)
	fmt.Println("Part 2:", p2Guard*p2Minute)
}
