// Advent of Code 2022 - Day 13.
package day13

import (
	"fmt"
	"io/ioutil"
	"log"
	"reflect"
	"strconv"
	"strings"
)

// tokenize converts the given string in to a series of tokens.
// Each token can be one of: "[", "]" or string representation of an integer.
func tokenize(s string) []string {
	s = strings.ReplaceAll(s, "[", "[ ")
	s = strings.ReplaceAll(s, "]", " ]")
	s = strings.ReplaceAll(s, "  ", " ")
	s = strings.ReplaceAll(s, ",", " ")
	return strings.Split(s, " ")
}

// parseTokens is a recursive descent parser for converting a set of tokens in to a data structure for a packet.
func parseTokens(tokens []string, i int) ([]any, int) {
	out := []any{}
	for i < len(tokens) {
		switch tokens[i] {
		case "[":
			subpacket, newI := parseTokens(tokens, i+1)
			out = append(out, subpacket)
			i = newI
		case "]":
			return out, i + 1
		default:
			v, err := strconv.Atoi(tokens[i])
			if err != nil {
				log.Fatal(err)
			}
			out = append(out, v)
			i += 1
		}
	}
	return out, i
}

// parsePacket converts the provided string in to an any type.
// The returned value is a slice where the elements may either be integers or slices.
func parsePacket(s string) any {
	tokens := tokenize(s)
	p, _ := parseTokens(tokens, 0)
	return p[0]
}

func parseInput(data []byte) [][2]any {
	groups := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n\n")
	packetPairs := make([][2]any, len(groups))
	for i, group := range groups {
		pairs := strings.Split(group, "\n")
		packetPairs[i][0] = parsePacket(pairs[0])
		packetPairs[i][1] = parsePacket(pairs[1])
	}
	return packetPairs
}

// Constants for the output of the compare function.
const (
	unknownOrder = iota
	correctOrder
	wrongOrder
)

// compare compares the two given packets.
// If left is smaller then right, then correctOrder is returned.
// If left is bigger then right, then wrongOrder is returned.
// Otherwise, unknownOrder is returned.
func compare(left, right any) int {
	lType := reflect.ValueOf(left).Kind()
	rType := reflect.ValueOf(right).Kind()

	switch {
	case lType == reflect.Int && rType == reflect.Int:
		lInt := left.(int)
		rInt := right.(int)
		if lInt < rInt {
			return correctOrder
		} else if lInt > rInt {
			return wrongOrder
		} else {
			return unknownOrder
		}

	case lType == reflect.Slice && rType == reflect.Slice:
		lSlice := left.([]any)
		rSlice := right.([]any)
		if len(lSlice) == 0 && len(rSlice) > 0 {
			return correctOrder
		}
		if len(rSlice) == 0 && len(lSlice) > 0 {
			return wrongOrder
		}
		if len(lSlice) == 0 && len(rSlice) == 0 {
			return unknownOrder
		}
		if rv := compare(lSlice[0], rSlice[0]); rv == unknownOrder {
			return compare(lSlice[1:], rSlice[1:])
		} else {
			return rv
		}

	case lType == reflect.Int && rType == reflect.Slice:
		return compare([]any{left}, right)

	case lType == reflect.Slice && rType == reflect.Int:
		return compare(left, []any{right})

	}
	return unknownOrder
}

func part1(packetPairs [][2]any) int {
	sum := 0
	for i, p := range packetPairs {
		if r := compare(p[0], p[1]); r == correctOrder {
			sum += i + 1 // Convert from zero to one based indexing.
		}
	}
	return sum
}

func part2(packetPairs [][2]any) int {
	p1 := parsePacket("[[2]]")
	p2 := parsePacket("[[6]]")
	p1GT := 0
	p2GT := 0

	for _, pair := range packetPairs {
		for _, p := range pair {
			if r := compare(p, p1); r == correctOrder {
				p1GT += 1
			}
			if r := compare(p, p2); r == correctOrder {
				p2GT += 1
			}
		}
	}

	p1Idx := p1GT + 1 // Convert from zero to one based indexing.
	p2Idx := p2GT + 2 // Convert from zero to one based indexing and account for p1.

	return p1Idx * p2Idx
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	packets := parseInput(data)

	p1 := part1(packets)
	fmt.Println("Part 1:", p1)

	p2 := part2(packets)
	fmt.Println("Part 2:", p2)
}
