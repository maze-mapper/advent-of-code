// Advent of Code 2018 - Day 3
package day3

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// fabric represents the area of fabric.
// 0 represents no claims.
// -1 represents where mutliple claims overlap.
// Any other integer is the ID of a claim in that area.
type fabric [1000][1000]int

type void struct{}

// Keys returns the keys of the given map as a slice
func Keys(m map[int]void) []int {
	keys := make([]int, len(m))
	i := 0
	for k := range m {
		keys[i] = k
		i++
	}
	return keys
}

// makeClaim marks the fabric with the given claim.
// It returns the IDs of any claim that overlaps with another.
func (f *fabric) makeClaim(c claim) []int {
	// Store any collsions with other claims
	collisions := make(map[int]void)

	for i := c.y; i <= c.y+c.height-1; i++ {
		for j := c.x; j <= c.x+c.width-1; j++ {
			elem := f[i][j]

			switch elem {

			// No other claims
			case 0:
				f[i][j] = c.id

			// Already been a collsion here
			case -1:
				collisions[c.id] = void{}

			// Resolve collision
			default:
				collisions[c.id] = void{}
				collisions[elem] = void{}
				f[i][j] = -1

			}
		}
	}

	return Keys(collisions)
}

// collisionArea returns the total number of elements where claims overlap
func (f *fabric) collisionArea() int {
	count := 0
	for y := 0; y < len(f); y++ {
		for x := 0; x < len(f[y]); x++ {
			if f[y][x] == -1 {
				count += 1
			}
		}
	}
	return count
}

// claim holds the information for a claim on the fabric
type claim struct {
	id, x, y, width, height int
}

func readData(file string) []byte {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func parseData(data []byte) []claim {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)

	claims := make([]claim, len(lines))
	for i, line := range lines {
		c := claim{}
		if _, err := fmt.Sscanf(line, "#%d @ %d,%d: %dx%d", &c.id, &c.x, &c.y, &c.width, &c.height); err != nil {
			log.Fatal(err)
		} else {
			claims[i] = c
		}
	}

	return claims
}

// solve applies the claims to the fabric.
// It returns the area where mutliple claims overlap and the ID of the one claim that does not overlap.
func solve(claims []claim) (int, int) {
	var f fabric
	uniqueClaims := make(map[int]void)
	for _, c := range claims {
		uniqueClaims[c.id] = void{}
		collisions := f.makeClaim(c)
		for _, id := range collisions {
			delete(uniqueClaims, id)
		}
	}
	uc := Keys(uniqueClaims)
	if len(uc) != 1 {
		log.Fatal("More than one unique claim found")
	}
	return f.collisionArea(), uc[0]
}

func Run(inputFile string) {
	data := readData(inputFile)
	claims := parseData(data)
	part1, part2 := solve(claims)
	fmt.Println("Part 1:", part1)
	fmt.Println("Part 2:", part2)
}
