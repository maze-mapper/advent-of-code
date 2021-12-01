// Advent of Code 2015 - Day 155
package day15

import (
	"fmt"
	"io/ioutil"
	"log"
	"regexp"
	"strconv"
	"strings"
)

var lineRegExp = regexp.MustCompile(`([a-zA-Z]+): capacity (\-?[0-9]+), durability (\-?[0-9]+), flavor (\-?[0-9]+), texture (\-?[0-9]+), calories (\-?[0-9]+)`)

// Ingredient holds the attributes of an ingredient
type Ingredient struct {
	name                                             string
	capacity, durability, flavour, texture, calories int
}

// parseInput converts the input data in to a slice of Ingredients
func parseInput(input []byte) []Ingredient {
	lines := strings.Split(
		strings.TrimSuffix(string(input), "\n"), "\n",
	)

	ingredients := make([]Ingredient, len(lines))
	for i, line := range lines {
		matches := lineRegExp.FindStringSubmatch(line)
		properties := make([]int, 5) // There are five properties: capacity, durability, flavor, texture and calories
		for j, match := range matches[2:] {
			val, err := strconv.Atoi(match)
			if err != nil {
				log.Fatal(err)
			}
			properties[j] = val
		}

		ingredients[i] = Ingredient{
			name:       matches[1],
			capacity:   properties[0],
			durability: properties[1],
			flavour:    properties[2],
			texture:    properties[3],
			calories:   properties[4],
		}
	}

	return ingredients
}

// makeCombinations finds all combinations of ingredients that add up to the total number of teaspoons
func makeCombinations(numIngredients, currentTsp, maxTsp int, current []int, results *[][]int) {
	remainingTsp := maxTsp - currentTsp
	if numIngredients == 1 {
		newCurrent := append(current, remainingTsp)
		*results = append(*results, newCurrent)
	} else {
		for i := remainingTsp; i >= 0; i-- {
			newCurrent := append(current, i)
			makeCombinations(numIngredients-1, currentTsp+i, maxTsp, newCurrent, results)
		}
	}
}

// calculateScore returns the score for a particular combination of ingredients
func calculateScore(ingredients []Ingredient, recipe []int) int {
	var capacity, durability, flavour, texture int

	for i, r := range recipe {
		capacity += r * ingredients[i].capacity
		durability += r * ingredients[i].durability
		flavour += r * ingredients[i].flavour
		texture += r * ingredients[i].texture
	}

	if capacity <= 0 || durability <= 0 || flavour <= 0 || texture <= 0 {
		return 0
	} else {
		return capacity * durability * flavour * texture
	}
}

// sumCalories returns the number of calories for a particular combination of ingredients
func sumCalories(ingredients []Ingredient, recipe []int) int {
	var calories int

	for i, r := range recipe {
		calories += r * ingredients[i].calories
	}

	return calories
}

func part1(ingredients []Ingredient, recipes [][]int) {
	bestScore := 0
	bestRecipe := []int{}
	for _, recipe := range recipes {
		if score := calculateScore(ingredients, recipe); score > bestScore {
			bestScore = score
			bestRecipe = recipe
		}
	}

	fmt.Println("Part 1: best score is", bestScore)
	fmt.Println("Recipe is:")
	for i, r := range bestRecipe {
		fmt.Println("    ", r, "tsp of", ingredients[i].name)
	}
}

func part2(ingredients []Ingredient, recipes [][]int) {
	bestScore := 0
	bestRecipe := []int{}

	for _, recipe := range recipes {
		if calories := sumCalories(ingredients, recipe); calories == 500 {
			if score := calculateScore(ingredients, recipe); score > bestScore {
				bestScore = score
				bestRecipe = recipe
			}
		}
	}

	fmt.Println("Part 2: best score is", bestScore)
	fmt.Println("Recipe is:")
	for i, r := range bestRecipe {
		fmt.Println("    ", r, "tsp of", ingredients[i].name)
	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	ingredients := parseInput(data)
	combinations := [][]int{}
	makeCombinations(len(ingredients), 0, 100, []int{}, &combinations)

	part1(ingredients, combinations)
	part2(ingredients, combinations)
}
