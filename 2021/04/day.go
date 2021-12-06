// Advent of Code 2021 - Day 4
package day4

import (
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// boardSize is the dimensions of a Bingo board
var boardSize int = 5

// board holds information on a Bingo board
type board struct {
	numbers  [][]int
	marked   [][]bool
	complete bool
}

// New returns a new Bingo board
func New(numbers [][]int) *board {
	b := board{
		numbers: numbers,
	}
	b.marked = make([][]bool, boardSize)
	for i := 0; i < boardSize; i++ {
		b.marked[i] = make([]bool, boardSize)
	}
	return &b
}

// markNumber marks off a number on the board
func (b *board) markNumber(n int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b.numbers[i][j] == n {
				b.marked[i][j] = true
				return
			}
		}
	}
}

// bingo determines if a board has won
func (b *board) bingo() bool {
	// Check rows for bingo
r:
	for row := 0; row < boardSize; row++ {
		for col := 0; col < boardSize; col++ {
			if b.marked[row][col] == false {
				continue r
			}
		}
		b.complete = true
		return true
	}

	// Check columns for bingo
c:
	for col := 0; col < boardSize; col++ {
		for row := 0; row < boardSize; row++ {
			if b.marked[row][col] == false {
				continue c
			}
		}
		b.complete = true
		return true
	}

	return false
}

// Score returns the score of a board
func (b *board) Score(n int) int {
	sum := 0
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			if b.marked[i][j] == false {
				sum += b.numbers[i][j]
			}
		}
	}
	return n * sum
}

// parseBoard converts a textual representation of a board in to a board data structure
func parseBoard(text string) *board {
	lines := strings.Split(text, "\n")
	numbers := make([][]int, len(lines))
	for i, line := range lines {
		// Remove any double spaces for single digits
		line = strings.ReplaceAll(line, "  ", " ")
		line = strings.TrimPrefix(line, " ")

		cols := strings.Split(line, " ")
		numbers[i] = make([]int, len(cols))
		for j, s := range cols {
			if val, err := strconv.Atoi(s); err == nil {
				numbers[i][j] = val
			} else {
				log.Fatal(err)
			}
		}
	}
	return New(numbers)
}

// parseData extracts the called numbers and boards from the input data
func parseData(data []byte) ([]int, []*board) {
	blocks := strings.Split(
		strings.TrimSuffix(string(data), "\n"), "\n\n",
	)

	// Parse drawn numbers from first line
	numberData := strings.Split(blocks[0], ",")
	numbers := make([]int, len(numberData))
	for i, s := range numberData {
		if val, err := strconv.Atoi(s); err == nil {
			numbers[i] = val
		} else {
			log.Fatal(err)
		}
	}

	// Parse Bingo boards from remainder of file
	boards := make([]*board, len(blocks)-1)
	for i, b := range blocks[1:] {
		boards[i] = parseBoard(b)
	}

	return numbers, boards
}

// play will play Bingo with the given numbers and boards
func play(numbers []int, boards []*board) []int {
	scores := []int{}
	for _, n := range numbers {
		for _, b := range boards {
			if b.complete == false {
				b.markNumber(n)
				if b.bingo() == true {
					scores = append(scores, b.Score(n))
				}
			}
		}
	}
	return scores
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	numbers, boards := parseData(data)

	scores := play(numbers, boards)
	fmt.Println("Part 1:", scores[0])
	fmt.Println("Part 2:", scores[len(scores)-1])
}
