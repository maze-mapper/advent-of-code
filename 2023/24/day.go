// Advent of Code 2023 - Day 24.
package day24

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/maze-mapper/advent-of-code/coordinates"
)

type hailstone struct {
	pos, vel coordinates.Coord
}

func (h hailstone) gradient() float64 {
	return float64(h.vel.Y) / float64(h.vel.X)
}

func (h hailstone) constant() float64 {
	return float64(h.pos.Y) - h.gradient()*float64(h.pos.X)
}

func parseData(data []byte) []hailstone {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	hailstones := make([]hailstone, len(lines))
	for i, line := range lines {
		line = strings.ReplaceAll(line, ",", "")
		h := hailstone{}
		fmt.Sscanf(line, "%d %d %d @ %d %d %d", &h.pos.X, &h.pos.Y, &h.pos.Z, &h.vel.X, &h.vel.Y, &h.vel.Z)
		hailstones[i] = h
	}
	return hailstones
}

func intersect(h1, h2 hailstone) (float64, float64, bool) {
	// ax + by + c = 0
	// xI = (b1*c2 - b2*c1) / (a1*b2 - a2*b1)
	// yI = (a2*c1 - a1*c2) / (a1*b2 - a2*b1)

	denominator := h1.gradient() - h2.gradient()
	if denominator == 0 {
		return 0, 0, false
	}

	x := (h2.constant() - h1.constant()) / denominator
	y := -(h2.gradient()*h1.constant() - h1.gradient()*h2.constant()) / denominator

	return x, y, true
}

func part1(hailstones []hailstone) int {
	testAreaLo := coordinates.Coord{X: 200000000000000, Y: 200000000000000}
	testAreaHi := coordinates.Coord{X: 400000000000000, Y: 400000000000000}

	count := 0
	for i := 0; i < len(hailstones)-1; i++ {
		for j := i + 1; j < len(hailstones); j++ {
			h1 := hailstones[i]
			h2 := hailstones[j]
			x, y, ok := intersect(h1, h2)
			if !ok {
				continue
			}
			t1 := (y - float64(h1.pos.Y)) / float64(h1.vel.Y)
			t2 := (y - float64(h2.pos.Y)) / float64(h2.vel.Y)
			if t1 < 0 || t2 < 0 {
				continue
			}
			if x >= float64(testAreaLo.X) && x <= float64(testAreaHi.X) && y >= float64(testAreaLo.Y) && y <= float64(testAreaHi.Y) {
				count += 1
			}
		}
	}
	return count
}

func z3Commands(hailstones []hailstone) {
	// Variables to solve for.
	fmt.Println("(declare-const x Int)")
	fmt.Println("(declare-const y Int)")
	fmt.Println("(declare-const z Int)")
	fmt.Println("(declare-const vx Int)")
	fmt.Println("(declare-const vy Int)")
	fmt.Println("(declare-const vz Int)")

	// Let p_r and p_i be the position vectors of the rock and a hailstone.
	// Let v_r and v_i be the velocity vectors of the rock and a hailstone.
	// At some time t the rock and hailstone must be at the same position:
	//   p_i + v_i * t = p_r + v_r * t
	//   p_i - p_r + t * (v_i - v_r) = 0
	// This is a translation of hailstone i into the frame of reference of the rock, with the rock
	// at the origin.
	// Since the hailstone must cross the origin, its translated position vector and velocity vectors
	// must be antiparallel (have opposite directions). The cross-product of two parallel or antiparallel
	// vectors is zero.
	//   A × B = 0
	// So, each component of the cross product is also 0:
	//   A_y * B_z - A_z * B_y = 0
	//   A_z * B_x - A_x * B_z = 0
	//   A_x * B_y - A_y * B_x = 0
	// Where in this situation:
	//   A = p_i - p_r
	//   B = v_i - v_r
	for i := 0; i < 3; i++ {
		h := hailstones[i]
		// Components of the translated position and velocity vectors.
		posX := fmt.Sprintf("(- %d x)", h.pos.X)
		posY := fmt.Sprintf("(- %d y)", h.pos.Y)
		posZ := fmt.Sprintf("(- %d z)", h.pos.Z)
		velX := fmt.Sprintf("(- %d vx)", h.vel.X)
		velY := fmt.Sprintf("(- %d vy)", h.vel.Y)
		velZ := fmt.Sprintf("(- %d vz)", h.vel.Z)
		// Cross product components are zero.
		fmt.Printf("(assert (= (- (* %s %s) (* %s %s)) 0))\n", posY, velZ, posZ, velY)
		fmt.Printf("(assert (= (- (* %s %s) (* %s %s)) 0))\n", posZ, velX, posX, velZ)
		fmt.Printf("(assert (= (- (* %s %s) (* %s %s)) 0))\n", posX, velY, posY, velX)
	}
	fmt.Println("(echo \"Checking if satisfiable...\")")
	fmt.Println("(check-sat)")
	fmt.Println("(get-model)")
	fmt.Println("(eval (+ x (+ y z)))")
}

func part2(hailstones []hailstone) int {
	fmt.Println("Use the following as the input to the Z3 Satisfiability Modulo Theories (SMT) solver:")
	z3Commands(hailstones)
	return 0
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	hailstones := parseData(data)

	p1 := part1(hailstones)
	fmt.Println("Part 1:", p1)

	p2 := part2(hailstones)
	fmt.Println("Part 2:", p2)
}
