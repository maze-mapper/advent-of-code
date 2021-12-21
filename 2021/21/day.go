// Advent of Code 2021 - Day 21
package day21

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// boardSize is the number of positions on the board
var boardSize int = 10

// dieSides is the number of sides on the deterministic die
var dieSides int = 100

// winningScore is the score required to win the game with a deterministic die
var winningScore int = 1000

// diracWinningScore is the score required to win the game with a Dirac die
var diracWinningScore int = 21

// DeterministicDie holds information on the state of the die
type DeterministicDie struct {
	rolls, value int
}

// Roll rolls the die and returns the roll value
func (d *DeterministicDie) Roll() int {
	d.rolls += 1
	d.value += 1
	if d.value > dieSides {
		d.value = d.value % dieSides
	}
	return d.value
}

// GetRolls returns the number of times the die has been rolled
func (d *DeterministicDie) GetRolls() int {
	return d.rolls
}

func parseData(data []byte) (int, int) {
	lines := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n",
	)
	var playerOne, playerTwo int
	fmt.Sscanf(lines[0], "Player 1 starting position: %d", &playerOne)
	fmt.Sscanf(lines[1], "Player 2 starting position: %d", &playerTwo)
	return playerOne, playerTwo
}

// getMove returns the sum of the next three deterministic die rolls
func getMove(die *DeterministicDie) int {
	move := 0
	for roll := 0; roll < 3; roll++ {
		move += die.Roll()
	}
	return move
}

// play plays the game with a deterministic die
func play(playerOne, playerTwo int) int {
	die := DeterministicDie{}
	playerPositions := [2]int{playerOne, playerTwo}
	playerScores := [2]int{0, 0}
	currentPlayer := 0
	for playerScores[0] < winningScore && playerScores[1] < winningScore {
		move := getMove(&die)
		move = move % boardSize

		playerPositions[currentPlayer] += move
		if playerPositions[currentPlayer] > boardSize {
			playerPositions[currentPlayer] = playerPositions[currentPlayer] % boardSize
		}

		playerScores[currentPlayer] += playerPositions[currentPlayer]
		currentPlayer = (currentPlayer + 1) % 2
	}

	// currentPlayer is now the losing player
	return playerScores[currentPlayer] * die.GetRolls()
}

// diracRolls returns the sum of three Dirac die rolls and the number of ways that roll can be achieved
func diracRolls() map[int]int {
	return map[int]int{
		3: 1,
		4: 3,
		5: 6,
		6: 7,
		7: 6,
		8: 3,
		9: 1,
	}
}

// playDirac plays the game with a Dirac die
func playDirac(playerPositions, playerScores [2]int, currentPlayer int, wins map[int]int, universes int) {
	if playerScores[0] >= diracWinningScore {
		wins[1] += universes
		return
	} else if playerScores[1] >= diracWinningScore {
		wins[2] += universes
		return
	}

	moves := diracRolls()
	for move, number := range moves {
		newPlayerPositions := playerPositions
		newPlayerScores := playerScores

		newPlayerPositions[currentPlayer] += move
		if newPlayerPositions[currentPlayer] > boardSize {
			newPlayerPositions[currentPlayer] = newPlayerPositions[currentPlayer] % boardSize
		}
		newPlayerScores[currentPlayer] += newPlayerPositions[currentPlayer]
		newCurrentPlayer := (currentPlayer + 1) % 2

		playDirac(newPlayerPositions, newPlayerScores, newCurrentPlayer, wins, universes*number)
	}
}

func part1(playerOne, playerTwo int) int {
	return play(playerOne, playerTwo)
}

func part2(playerOne, playerTwo int) int {
	playerPositions := [2]int{playerOne, playerTwo}
	playerScores := [2]int{0, 0}
	wins := map[int]int{}
	playDirac(playerPositions, playerScores, 0, wins, 1)
	if wins[0] > wins[1] {
		return wins[0]
	} else {
		return wins[1]
	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	playerOne, playerTwo := parseData(data)

	p1 := part1(playerOne, playerTwo)
	fmt.Println("Part 1:", p1)

	p2 := part2(playerOne, playerTwo)
	fmt.Println("Part 2:", p2)
}
