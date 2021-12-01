// Advent of Code 2015 - Day 122
package day12

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
)

// part1 counts all numbers in a JSON document
func part1(i interface{}, c *float64) {
	switch i.(type) {

	case float64:
		*c += i.(float64)

	case string:
		// Nothing to count but stops checking other cases

	// JSON array
	case []interface{}:
		for _, v := range i.([]interface{}) {
			part1(v, c)
		}

	// JSON object
	case map[string]interface{}:
		for _, v := range i.(map[string]interface{}) {
			part1(v, c)
		}

	}
}

// part2 counts all numbers in a JSON document but skips any JSON object with a value of "red"
func part2(i interface{}, c *float64) bool {
	switch i.(type) {

	case float64:
		*c += i.(float64)

	case string:
		if i.(string) == "red" {
			return true
		}

	// JSON array
	case []interface{}:
		for _, v := range i.([]interface{}) {
			part2(v, c)
		}

	// JSON object
	case map[string]interface{}:
		var tempCount float64
		for _, v := range i.(map[string]interface{}) {
			if foundRed := part2(v, &tempCount); foundRed {
				// Skip this object
				return false
			}
		}
		*c += tempCount

	}

	return false
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	var decodedJSON interface{}
	err = json.Unmarshal(data, &decodedJSON)

	var count float64
	part1(decodedJSON, &count)
	fmt.Println("Part 1:", count)

	count = 0
	part2(decodedJSON, &count)
	fmt.Println("Part 2:", count)
}
