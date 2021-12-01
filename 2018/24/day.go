// Advent of Code 2018 - Day 24
package day24

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// group represents a group of units
type group struct {
	units, hitPoints, attackDamage, initiative int
	attackType                                 string
	weaknesses, immunities                     []string
}

// effectivePower calcualtes the effective power of a group of units
func (g *group) effectivePower() int {
	return g.units * g.attackDamage
}

// applyDamage reduces the unit count for a group after taking damage
func (g *group) applyDamage(damage int) bool {
	casualties := damage / g.hitPoints
	if casualties == 0 {
		return false
	} else if casualties >= g.units {
		g.units = 0
	} else {
		g.units -= casualties
	}
	return true
}

// contains tests if a slice "sl" contains a string "s"
func contains(sl []string, s string) bool {
	for _, ss := range sl {
		if ss == s {
			return true
		}
	}
	return false
}

// calculateDamage returns the damage an attacking group would deal to a defending group
func calculateDamage(attacker, defender *group) int {
	switch {
	case contains(defender.immunities, attacker.attackType):
		return 0

	case contains(defender.weaknesses, attacker.attackType):
		return 2 * attacker.effectivePower()

	default:
		return attacker.effectivePower()
	}
}

var pattern = regexp.MustCompile(`([0-9]+) units each with ([0-9]+) hit points (\([a-z,; ]+\) )?with an attack that does ([0-9]+) ([a-z]+) damage at initiative ([0-9]+)`)

// parseGroup converts a textual representation in to a group object
func parseGroup(line string) *group {
	if match := pattern.FindStringSubmatch(line); match != nil {
		// Due to regexp match there should be no errors converting to int
		units, _ := strconv.Atoi(match[1])
		hitPoints, _ := strconv.Atoi(match[2])
		attackDamage, _ := strconv.Atoi(match[4])
		initiative, _ := strconv.Atoi(match[6])

		// Find weaknesses and immunities
		parts := strings.Split(
			strings.TrimSuffix(
				strings.TrimPrefix(match[3], "("), ") ",
			), "; ",
		)
		weaknesses := []string{}
		immunities := []string{}
		for _, p := range parts {

			p = strings.TrimSuffix(p, ")")
			switch {
			case strings.HasPrefix(p, "weak to "):
				subparts := strings.Split(
					strings.TrimPrefix(p, "weak to "), ", ",
				)
				weaknesses = append(weaknesses, subparts...)

			case strings.HasPrefix(p, "immune to "):
				subparts := strings.Split(
					strings.TrimPrefix(p, "immune to "), ", ",
				)
				immunities = append(immunities, subparts...)
			}
		}

		g := &group{
			units:        units,
			hitPoints:    hitPoints,
			weaknesses:   weaknesses,
			immunities:   immunities,
			attackDamage: attackDamage,
			attackType:   match[5],
			initiative:   initiative,
		}
		return g
	} else {
		log.Fatal("Cannot parse line for group of units")
	}
	return &group{}
}

// parseArmy converts a single army in to a slice of groups
func parseArmy(lines []string) []*group {
	army := make([]*group, len(lines))
	for i, line := range lines {
		army[i] = parseGroup(line)
	}
	return army
}

// parseData reads the input text file and returns groups for each army
func parseData(file string) ([]*group, []*group) {
	data, err := ioutil.ReadFile(file)
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n\n",
	)
	// Lines for each army, removing line with army name
	immuneSystemLines := strings.Split(lines[0], "\n")[1:]
	infectionLines := strings.Split(lines[1], "\n")[1:]

	immuneSystem := parseArmy(immuneSystemLines)
	infection := parseArmy(infectionLines)

	return immuneSystem, infection
}

// sortByEffectivePower sorts groups by descending effective power. Initiative breaks ties.
func sortByEffectivePower(army []*group) {
	sort.Slice(army, func(i, j int) bool {
		if army[i].effectivePower() == army[j].effectivePower() {
			return army[i].initiative > army[j].initiative
		} else {
			return army[i].effectivePower() > army[j].effectivePower()
		}
	})
}

// sortByInitiative sorts groups by descending initiative
func sortByInitiative(army []*group) {
	sort.Slice(army, func(i, j int) bool {
		return army[i].initiative > army[j].initiative
	})
}

func targetSelection(attackers, defenders []*group) map[*group]*group {
	sortByEffectivePower(attackers)

	// Make a copy of the defenders so we don't alter the original
	newDefenders := make([]*group, len(defenders))
	copy(newDefenders, defenders)
	defenders = newDefenders

	targets := map[*group]*group{}
	for _, attacker := range attackers {
		maxDamage := 0
		target := -1
		for i, defender := range defenders {
			damage := calculateDamage(attacker, defender)
			switch {

			case damage == 0:
				continue

			case damage == maxDamage:
				update := false
				if defender.effectivePower() == defenders[target].effectivePower() {
					update = defender.initiative > defenders[target].initiative
				} else {
					update = defender.effectivePower() > defenders[target].effectivePower()
				}
				if update {
					target = i
				}

			case damage > maxDamage:
				target = i
				maxDamage = damage

			}
		}

		if target != -1 {
			targets[attacker] = defenders[target]
			// Remove chosen target from defenders
			defenders[target] = defenders[len(defenders)-1]
			defenders = defenders[:len(defenders)-1]
		}

	}

	return targets
}

func fight(army1, army2 []*group) bool {
	stalemate := true

	// Target selection phase
	army1Targets := targetSelection(army1, army2)
	army2Targets := targetSelection(army2, army1)

	// Combine both army targets and find attackers
	targets := map[*group]*group{}
	attackers := []*group{}
	for _, army := range []map[*group]*group{army1Targets, army2Targets} {
		for k, v := range army {
			attackers = append(attackers, k)
			targets[k] = v
		}
	}

	// Attacking phase
	// Sort attackers in initiative order
	sortByInitiative(attackers)

	// Resolve each combat
	for _, attacker := range attackers {
		if attacker.units > 0 {
			defender := targets[attacker]
			damage := calculateDamage(attacker, defender)
			casualties := defender.applyDamage(damage)
			if casualties {
				stalemate = false
			}
		}
	}

	return stalemate
}

func combat(army1, army2 []*group) (int, bool) {
	for len(army1) > 0 && len(army2) > 0 {
		stalemate := fight(army1, army2)
		if stalemate {
			return 0, false
		}

		// Remove groups with no units
		newArmy1 := []*group{}
		for _, g := range army1 {
			if g.units > 0 {
				newArmy1 = append(newArmy1, g)
			}
		}
		army1 = newArmy1

		newArmy2 := []*group{}
		for _, g := range army2 {
			if g.units > 0 {
				newArmy2 = append(newArmy2, g)
			}
		}
		army2 = newArmy2
	}

	// Check who won and determine the number of remaining units
	var army1Win bool
	units := 0

	if len(army1) > 0 {
		for _, g := range army1 {
			units += g.units
		}
		army1Win = true
	} else {
		for _, g := range army2 {
			units += g.units
		}
		army1Win = false
	}

	return units, army1Win
}

// copyArmy returns a deep copy of an army
func copyArmy(army []*group) []*group {
	newArmy := make([]*group, len(army))
	for i, g := range army {
		newArmy[i] = &group{
			units:        g.units,
			hitPoints:    g.hitPoints,
			attackDamage: g.attackDamage,
			initiative:   g.initiative,
			attackType:   g.attackType,
			//			weaknesses: copy([]string{}, g.weaknesses),
			//			immunities: copy([]string{}, g.immunities),
		}
		newArmy[i].weaknesses = append(newArmy[i].weaknesses, g.weaknesses...)
		newArmy[i].immunities = append(newArmy[i].immunities, g.immunities...)
	}
	return newArmy
}

// boostDamage increases the attack damage for every group in an army
func boostDamage(army []*group, boost int) {
	for _, g := range army {
		g.attackDamage += boost
	}
}

func Run(inputFile string) {
	originalImmuneSystem, originalInfection := parseData(inputFile)

	// Part 1
	immuneSystem := copyArmy(originalImmuneSystem)
	infection := copyArmy(originalInfection)
	units, _ := combat(immuneSystem, infection)
	fmt.Println("Part 1:", units)

	// Part 2
	immuneSystemWon := false
	for boost := 1; !immuneSystemWon; boost++ {
		immuneSystem := copyArmy(originalImmuneSystem)
		infection := copyArmy(originalInfection)
		boostDamage(immuneSystem, boost)

		units, immuneSystemWon = combat(immuneSystem, infection)
	}
	fmt.Println("Part 2:", units)
}
