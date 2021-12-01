// Advent of Code 2015 - Day 211
package day21

import (
	"fmt"
)

type Item struct {
	name                 string
	cost, damage, armour int
}

type Character struct {
	hp, damage, defence int
}

func (c *Character) Equip(items []Item) int {
	cost := 0
	for _, item := range items {
		cost += item.cost
		c.damage += item.damage
		c.defence += item.armour
	}
	return cost
}

var weapons = [5]Item{
	Item{"Dagger", 8, 4, 0},
	Item{"Shortsword", 10, 5, 0},
	Item{"Warhammer", 25, 6, 0},
	Item{"Longsword", 40, 7, 0},
	Item{"Greataxe", 74, 8, 0},
}

var armours = [6]Item{
	Item{"", 0, 0, 0},
	Item{"Leather", 13, 0, 1},
	Item{"Chainmail", 31, 0, 2},
	Item{"Splintmail", 53, 0, 3},
	Item{"Bandedmail", 75, 0, 4},
	Item{"Platemail", 102, 0, 5},
}

var rings = [8]Item{
	Item{"", 0, 0, 0},
	Item{"", 0, 0, 0},
	Item{"Damage +1", 25, 1, 0},
	Item{"Damage +2", 50, 2, 0},
	Item{"Damage +3", 100, 3, 0},
	Item{"Defence +1", 20, 0, 1},
	Item{"Defence +2", 40, 0, 2},
	Item{"Defence +3", 80, 0, 3},
}

func doTurn(attacker, defender *Character) {
	dmg := attacker.damage - defender.defence
	// Always deal at least 1 damage
	if dmg <= 0 {
		dmg = 1
	}
	// Minimum HP is zero
	if dmg > defender.hp {
		defender.hp = 0
	} else {
		defender.hp -= dmg
	}
}

func play(player, boss Character) int {
	for player.hp > 0 && boss.hp > 0 {
		doTurn(&player, &boss)
		doTurn(&boss, &player)
	}
	if boss.hp == 0 {
		return 1 // Player victory
	} else {
		return 0 // Player defeat
	}
}

func Run(inputFile string) {
	boss := Character{104, 8, 1}
	var maxCost int = 0
	minCost := int(^uint(0) >> 1) // Initialise to max value

	for _, weapon := range weapons {
		for _, armour := range armours {
			for ri, rightHand := range rings {
				for li, leftHand := range rings {
					if ri == li {
						continue
					}
					player := Character{100, 0, 0}
					items := []Item{weapon, armour, rightHand, leftHand}
					cost := player.Equip(items)
					win := play(player, boss)

					switch win {
					case 1: // Player win
						if cost < minCost {
							minCost = cost
						}
					case 0: // Player defeat
						if cost > maxCost {
							maxCost = cost
						}
					}

				}

			}
		}
	}

	fmt.Printf("Player wins with a min spend of %d\n", minCost)
	fmt.Printf("Player loses with a max spend of %d\n", maxCost)

}
