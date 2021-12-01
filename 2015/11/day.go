// Advent of Code 2015 - Day 111
package day11

import (
	"fmt"
	"io/ioutil"
	"log"
	"strings"
)

// nextPassword increments the password
func nextPassword(password []byte) []byte {
	// In ASCII "a" through "z" are represented by 97 to 122
	for i := len(password) - 1; i >= 0; i-- {
		if password[i] == 122 { // "z"
			if i == 0 {
				log.Fatal("No more passwords available at this length")
			}
			password[i] = 97 // "a"
		} else {
			password[i] += 1
			break
		}
	}
	return password
}

// checkStraight checks if there is an increasing straight of at least three letters
func checkStraight(password []byte) bool {
	for i := 0; i < len(password)-2; i++ {
		if password[i] == password[i+1]-1 && password[i] == password[i+2]-2 {
			return true
		}
	}
	return false
}

// checkBannedLetter checks if the letters "i", "l" or "o" appear in the password
func checkBannedLetter(password []byte) bool {
	for _, b := range password {
		if b == 105 || b == 108 || b == 111 { // "i", "l" or "o"
			return false
		}
	}
	return true
}

// checkPairs checks if there are two different non-overlapping pairs of letters
func checkPairs(password []byte) bool {
	for i := 0; i < len(password)-3; i++ {
		if password[i] == password[i+1] {
			for j := i + 2; j < len(password)-1; j++ {
				if password[j] == password[j+1] {
					return true
				}
			}
			return false // No pairs in the remainder of the slice
		}
	}
	return false
}

// findNextValidPassword will return the next valid password
func findNextValidPassword(password []byte) []byte {
	for {
		password = nextPassword(password)
		if checkBannedLetter(password) && checkStraight(password) && checkPairs(password) {
			return password
		}
	}
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	text := strings.TrimSuffix(string(data), "\n")
	password := []byte(text)

	fmt.Println("Initial password:", text)

	password = findNextValidPassword(password)
	fmt.Println("Part 1:", string(password))

	password = findNextValidPassword(password)
	fmt.Println("Part 2:", string(password))
}
