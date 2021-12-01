// Advent of Code 2015 - Day 222
package day22

import (
	"fmt"
)

type Character struct {
	hp, damage, armour, mana int
	effects                  map[string]Effect
}

type Effect struct {
	duration, damage, armour, manaRestore int
}

type Spell struct {
	cost, damage, heal       int
	userEffect, targetEffect Effect
}

var spellMap = map[string]Spell{
	"Magic Missile": Spell{
		cost:   53,
		damage: 4,
	},
	"Drain": Spell{
		cost:   73,
		damage: 2,
		heal:   2,
	},
	"Shield": Spell{
		cost: 113,
		userEffect: Effect{
			duration: 6,
			armour:   7,
		},
	},
	"Poison": Spell{
		cost: 173,
		targetEffect: Effect{
			duration: 6,
			damage:   3,
		},
	},
	"Recharge": Spell{
		cost: 229,
		userEffect: Effect{
			duration:    5,
			manaRestore: 101,
		},
	},
}

// clone returns a deep copy of a Character
func (char *Character) clone() Character {
	newChar := Character{
		hp:      char.hp,
		damage:  char.damage,
		armour:  char.armour,
		mana:    char.mana,
		effects: make(map[string]Effect),
	}
	for name, effect := range char.effects {
		newChar.effects[name] = effect
	}
	return newChar
}

// applyDamage applies damage to a character
func (char *Character) applyDamage(dmg int) {
	// Always deal at least 1 damage
	if dmg <= 0 {
		dmg = 1
	}
	// HP should not go below zero
	if dmg > char.hp {
		char.hp = 0
	} else {
		char.hp -= dmg
	}
}

// newEffect adds a new effect to a character
func (char *Character) newEffect(name string, effect Effect) {
	char.effects[name] = effect
	if effect.armour != 0 {
		char.armour += effect.armour
	}
}

// applyCurentEffects
func (char *Character) applyCurentEffects() {
	expiredEffects := []string{}
	fmt.Println("Before", char)
	for name := range char.effects {
		effect := char.effects[name]
		effect.duration -= 1
		if effect.duration == 0 {
			expiredEffects = append(expiredEffects, name)
			if effect.armour != 0 {
				char.armour -= effect.armour
			}
		}

		if effect.damage != 0 {
			char.applyDamage(effect.damage)
		}
		if effect.manaRestore != 0 {
			char.mana += effect.manaRestore
		}
		char.effects[name] = effect
	}

	for _, name := range expiredEffects {
		delete(char.effects, name)
	}
	fmt.Println("After", char)
}

// playerTurn executes the actions for the player turn
func playerTurn(spellName string, player, boss *Character) {
	spell := spellMap[spellName]
	player.applyCurentEffects()
	boss.applyCurentEffects()

	player.mana -= spell.cost
	if spell.damage != 0 {
		dmg := spell.damage - boss.armour
		boss.applyDamage(dmg)
	}
	if spell.heal != 0 {
		player.hp += spell.heal // Assuming no max HP
	}
	if spell.userEffect != (Effect{}) {
		player.newEffect(spellName, spell.userEffect)
	}
	if spell.targetEffect != (Effect{}) {
		boss.newEffect(spellName, spell.targetEffect)
	}
}

// bossTurn executes the actions for the boss turn
func bossTurn(player, boss *Character) {
	player.applyCurentEffects()
	boss.applyCurentEffects()

	dmg := boss.damage - player.armour
	player.applyDamage(dmg)
}

// turn first executes the player turn and then the boss turn
func turn(spellName string, player, boss Character) (Character, Character, int) {
	playerTurn(spellName, &player, &boss)
	bossTurn(&player, &boss)

	gameState := gameOngoing
	if boss.hp == 0 {
		gameState = gameWon
	} else if player.hp == 0 {
		gameState = gameLost
	}

	return player, boss, gameState
}

const (
	gameOngoing = iota
	gameWon
	gameLost
)

type Game struct {
	player, boss Character
	manaSpent    int
}

func tree(minManaSpent *int, game Game) {
	for spellName, spell := range spellMap {
		if game.player.mana < spell.cost {
			continue
		}
		manaSpent := game.manaSpent + spell.cost

		// Skip if are already worse than the current minManaSpent
		if manaSpent > *minManaSpent {
			continue
		}

		// Cannot cast spells with active effects
		// Can cast them if the effect is about to expire
		if spell.userEffect != (Effect{}) {
			eff, ok := game.player.effects[spellName]
			if ok == true && eff.duration > 1 {
				continue
			}
		}
		if spell.targetEffect != (Effect{}) {
			eff, ok := game.boss.effects[spellName]
			if ok == true && eff.duration > 1 {
				continue
			}
		}

		fmt.Println("Run(ning turns with", spellName, game, *minManaSpent, manaSpent)

		player, boss, gameState := turn(spellName, game.player.clone(), game.boss.clone())

		if gameState == gameWon {
			fmt.Println("========== WON ==========\n")
			if manaSpent < *minManaSpent {
				*minManaSpent = manaSpent
			}
		} else if gameState == gameLost {
			fmt.Println("========== LOST ==========\n")
		} else if gameState == gameOngoing {
			newGame := Game{
				player:    player,
				boss:      boss,
				manaSpent: manaSpent,
			}
			fmt.Println("Player", game.player, newGame.player)
			tree(minManaSpent, newGame)
		}
	}
}

func Run(inputFile string) {
	minManaSpent := int(^uint(0) >> 1) // Initialise to max value

	game := Game{
		player: Character{
			hp:      50,
			mana:    500,
			effects: map[string]Effect{},
		},
		boss: Character{
			hp:      55,
			damage:  8,
			effects: map[string]Effect{},
		},
		manaSpent: 0,
	}
	tree(&minManaSpent, game)
	fmt.Println(minManaSpent)
}
