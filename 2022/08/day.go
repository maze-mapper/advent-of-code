// Advent of Code 2022 - Day 8.
package day8

import (
        "fmt"
        "io/ioutil"
        "log"
        "strconv"
        "strings"
)

func parseInput(data []byte) [][]int {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	out := make([][]int, len(lines))
	for i, line := range lines {
		out[i] = make([]int, len(line))
		for j, c := range line {
			val, err := strconv.Atoi(string(c))
			if err != nil {
				log.Fatal(err)
			}
			out[i][j] = val
		}
	}
	return out
}

func part1(trees [][]int) int {
	height := len(trees)
	width := len(trees[0])

	visible := make([][]bool, len(trees))
	for i := 0; i < len(trees); i++ {
		visible[i] = make([]bool, len(trees[i]))
	}

	// Horizontal.
	for i, row := range trees {
		left := 0
		right := len(row) - 1
		leftH := row[left]
		rightH := row[right]

		visible[i][left] = true
		visible[i][right] = true

		for ; left < right ; {
			if rightH > leftH {
				left += 1
				if row[left] > leftH {
					visible[i][left] = true
					leftH = row[left]
				}
			} else {
				right -= 1
				if row[right] > rightH {
					visible[i][right] = true
					rightH = row[right]
				}
			}
		}
	}

	// Vertical.
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			top := 0
			bottom := height - 1

			topH := trees[top][j]
			bottomH := trees[bottom][j]

			visible[top][j] = true
			visible[bottom][j] = true

			for ; top < bottom; {
				if bottomH > topH {
					top += 1
					if trees[top][j] > topH {
						visible[top][j] = true
						topH = trees[top][j]
					}
				} else {
					bottom -= 1
					if trees[bottom][j] > bottomH {
						visible[bottom][j] = true
						bottomH = trees[bottom][j]
					}
				}
			}
		}
	}

	count := 0
	for _, row := range visible {
		for _, val := range row {
			if val {
				count += 1
			}
		}
	}
	return count
}

func part2(trees [][]int) int {
	height := len(trees)
	width := len(trees[0])
	maxScore := 0
	for i := 0; i < height; i++ {
		for j := 0; j < width; j++ {
			// Up.
			upSum := 0
			for ii := i - 1; ii >= 0; ii-- {
				upSum += 1
				if trees[ii][j] >= trees[i][j] {
					break
				}
			}
			// Down.
			downSum := 0
                        for ii := i + 1; ii < height; ii++ {
				downSum += 1
                                if trees[ii][j] >= trees[i][j] {
					break
				}
                        }
			// Left.
			leftSum := 0
                        for jj := j - 1; jj >= 0; jj-- {
				leftSum += 1
                                if trees[i][jj] >= trees[i][j] {
					break
				}
                        }
			// Right.
			rightSum := 0
                        for jj := j + 1; jj < width; jj++ {
				rightSum += 1
                                if trees[i][jj] >= trees[i][j] {
                                        break
                                }
                        }

			score := upSum * downSum * leftSum * rightSum
			if score > maxScore {
				maxScore = score
			}
		}
	}
	return maxScore
}

func Run(inputFile string) {
        data, err := ioutil.ReadFile(inputFile)
        if err != nil {
                log.Fatal(err)
        }
        trees := parseInput(data)

        p1 := part1(trees)
        fmt.Println("Part 1:", p1)

        p2 := part2(trees)
        fmt.Println("Part 2:", p2)
}

