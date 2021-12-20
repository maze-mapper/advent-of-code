// Advent of Code 2021 - Day 20
package day20

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"

	"adventofcode/coordinates"
)

func parseData(data []byte) ([]bool, map[coordinates.Coord]bool) {
	parts := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n\n",
	)
	imageEnhancementAlgorithm := make([]bool, len(parts[0]))
	for i, r := range parts[0] {
		switch r {
		case '#':
			imageEnhancementAlgorithm[i] = true
		case '.':
			imageEnhancementAlgorithm[i] = false
		default:
			log.Fatal("Unexpected rune", r)
		}
	}

	pixels := map[coordinates.Coord]bool{}
	lines := strings.Split(parts[1], "\n")
	for i, line := range lines {
		for j, r := range line {
			if r == '#' {
				c := coordinates.Coord{X: j, Y: i}
				pixels[c] = true
			}
		}
	}

	return imageEnhancementAlgorithm, pixels
}

func getBinaryIndex(p coordinates.Coord, pixels map[coordinates.Coord]bool, defaultPixel bool) int {
	number := 0
	shift := 8
	for i := p.Y - 1; i < p.Y+2; i++ {
		for j := p.X - 1; j < p.X+2; j++ {
			c := coordinates.Coord{X: j, Y: i}
			// Absent values will take the default value
			val, ok := pixels[c]
			if (val && ok) || (!ok && defaultPixel) {
				number += 1 << shift
			}
			shift -= 1
		}
	}
	return number
}

func enhanceImage(algorithm []bool, pixels map[coordinates.Coord]bool, defaultPixel bool) map[coordinates.Coord]bool {
	newPixels := map[coordinates.Coord]bool{}

	list := make([]coordinates.Coord, len(pixels))
	i := 0
	for k, _ := range pixels {
		list[i] = k
		i += 1
	}

	min, max := coordinates.Range(list)
	for i := min.Y - 1; i < max.Y+2; i++ {
		for j := min.X - 1; j < max.X+2; j++ {
			c := coordinates.Coord{X: j, Y: i}
			b := getBinaryIndex(c, pixels, defaultPixel)
			newPixels[c] = algorithm[b]
		}
	}
	return newPixels
}

func countLightPixels(pixels map[coordinates.Coord]bool) int {
	count := 0
	for _, v := range pixels {
		if v {
			count += 1
		}
	}
	return count
}

func printPixels(pixels map[coordinates.Coord]bool) {
	list := make([]coordinates.Coord, len(pixels))
	i := 0
	for k, _ := range pixels {
		list[i] = k
		i += 1
	}

	min, max := coordinates.Range(list)
	for i := min.Y; i <= max.Y; i++ {
		for j := min.X; j <= max.X; j++ {
			c := coordinates.Coord{X: j, Y: i}
			if pixels[c] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Print("\n")
	}
}

func part1(algorithm []bool, pixels map[coordinates.Coord]bool) int {
	var defaultPixel bool
	for i := 0; i < 2; i++ {
		if i%2 == 0 {
			defaultPixel = false
		} else {
			defaultPixel = true
		}
		pixels = enhanceImage(algorithm, pixels, defaultPixel)
	}
	return countLightPixels(pixels)
}

func part2(algorithm []bool, pixels map[coordinates.Coord]bool) int {
	var defaultPixel bool
	for i := 0; i < 50; i++ {
		if i%2 == 0 {
			defaultPixel = false
		} else {
			defaultPixel = true
		}
		pixels = enhanceImage(algorithm, pixels, defaultPixel)
	}
	return countLightPixels(pixels)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	algorithm, pixels := parseData(data)

	p1 := part1(algorithm, pixels)
	fmt.Println("Part 1:", p1)

	p2 := part2(algorithm, pixels)
	fmt.Println("Part 2:", p2)
}
