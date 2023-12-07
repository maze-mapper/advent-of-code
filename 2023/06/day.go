// Advent of Code 2023 - Day 6.
package day6

import (
	"fmt"
	"log"
	"math"
	"os"
	"strconv"
	"strings"
)

func parseData(data []byte) ([]int, []int) {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	timeParts := strings.Fields(lines[0])[1:]
	distanceParts := strings.Fields(lines[1])[1:]
	if len(timeParts) != len(distanceParts) {
		log.Fatal("The number of time and distance components do not match.")
	}

	var times, distances []int
	for i := 0; i < len(timeParts); i++ {
		n, err := strconv.Atoi(timeParts[i])
		if err != nil {
			log.Fatal(err)
		}
		times = append(times, n)

		m, err := strconv.Atoi(distanceParts[i])
		if err != nil {
			log.Fatal(err)
		}
		distances = append(distances, m)
	}

	return times, distances
}

func solveQuadratic(a, b, c float64) (float64, float64) {
	lo := (-b - math.Sqrt(b*b-4*a*c)) / (2 * a)
	hi := (-b + math.Sqrt(b*b-4*a*c)) / (2 * a)
	return lo, hi
}

// waysToWin returns the number of ways that the record distance can be beaten.
// The speed is equal to the time the button is held.
// The distance travelled is then the speed times the remaining time.
//
//	dist = holdTime * (raceDuration - holdTime)
//
// This be rewritten as a quadratic equation:
//
//	holdTime^2 - raceDuration*holdTime + dist = 0
//
// This has two solutions for the holdTime and the number of integers between these is the number
// of ways to win.
func waysToWin(raceDuration, recordDistance int) int {
	lo, hi := solveQuadratic(1, -float64(raceDuration), float64(recordDistance+1))
	ways := math.Floor(hi) - math.Ceil(lo) + 1
	return int(ways)
}

func part1(times, distances []int) int {
	prod := 1
	for i := 0; i < len(times); i++ {
		duration := times[i]
		recordDistance := distances[i]
		prod *= waysToWin(duration, recordDistance)
	}
	return prod
}

func part2(times, distances []int) int {
	var timeStr, distStr string
	for _, t := range times {
		timeStr += strconv.Itoa(t)
	}
	for _, d := range distances {
		distStr += strconv.Itoa(d)
	}
	duration, err := strconv.Atoi(timeStr)
	if err != nil {
		log.Fatal(err)
	}
	recordDistance, err := strconv.Atoi(distStr)
	if err != nil {
		log.Fatal(err)
	}
	return waysToWin(duration, recordDistance)
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	times, distances := parseData(data)

	p1 := part1(times, distances)
	fmt.Println("Part 1:", p1)

	p2 := part2(times, distances)
	fmt.Println("Part 2:", p2)
}
