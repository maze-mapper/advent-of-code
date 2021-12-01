// Advent of Code 2019 - Day 7
package day7

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"adventofcode/2019/intcode"
)

// generatePermutations uses Heap's algorithm to generate all permutations of a slice
func generatePermutations(k int, s []int) [][]int {
	output := [][]int{}

	if k == 1 {
		// Append a copy of the slice to avoid altering this permutation
		sl := make([]int, len(s))
		copy(sl, s)
		output = append(output, sl)
	} else {
		// Generate permutations with the kth element and beyond unaltered
		output = append(output, generatePermutations(k-1, s)...)

		// Generate permutations for kth element swapped with each k - 1
		for i := 0; i < k-1; i++ {
			// Swap choice is dependent on parity of k (even or odd)
			if k%2 == 0 {
				tmp := s[i]
				s[i] = s[k-1]
				s[k-1] = tmp
			} else {
				tmp := s[0]
				s[0] = s[k-1]
				s[k-1] = tmp
			}
			output = append(output, generatePermutations(k-1, s)...)
		}
	}

	return output
}

// runAmplifierPrograms runs the given program on a set of amplifiers with the given phase settings
func runAmplifierPrograms(program, phaseSettings []int) int {
	numAmplifiers := len(phaseSettings)

	// Set up input channels
	inputChannels := make([]chan int, numAmplifiers)
	for i, phase := range phaseSettings {
		// Each input first takes the phase setting and then the signal
		// So create a buffered channel with length 2
		inputChannels[i] = make(chan int, 2)
		inputChannels[i] <- phase
	}
	// Pass in input signal to first amplifier
	inputChannels[0] <- 0

	// Set up and run amplifiers with connected channels
	outputChannel := make(chan int)
	for i := 0; i < len(inputChannels); i++ {
		computer := intcode.New(program)
		computer.SetChanIn(inputChannels[i])
		if i < len(phaseSettings)-1 {
			computer.SetChanOut(inputChannels[i+1])
		} else {
			computer.SetChanOut(outputChannel)
		}
		go computer.Run()
	}

	// Pass output signals back in to the first input channel and return the last signal received
	var output int
	for out := range outputChannel {
		inputChannels[0] <- out
		output = out
	}

	return output
}

// findMaxSignal finds the maximum output signal for a program and slice of phase settings
func findMaxSignal(program, phaseSettings []int) int {
	permutations := generatePermutations(len(phaseSettings), phaseSettings)
	maxSignal := 0
	signals := make(chan int)

	// Process output signals and update maxSignal
	var chanWg sync.WaitGroup
	chanWg.Add(1)
	go func() {
		defer chanWg.Done()
		for signal := range signals {
			if signal > maxSignal {
				maxSignal = signal
			}
		}
	}()

	// Run each permutation of phase settings
	var wg sync.WaitGroup
	for _, permutation := range permutations {
		p := permutation
		wg.Add(1)
		go func() {
			defer wg.Done()
			signals <- runAmplifierPrograms(program, p)
		}()
	}
	wg.Wait()
	close(signals)
	chanWg.Wait()

	return maxSignal
}

func part1(program []int) int {
	return findMaxSignal(program, []int{0, 1, 2, 3, 4})
}

func part2(program []int) int {
	return findMaxSignal(program, []int{5, 6, 7, 8, 9})
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	program := intcode.ReadProgram(data)

	p1 := part1(program)
	fmt.Println("Part 1:", p1)

	p2 := part2(program)
	fmt.Println("Part 2:", p2)
}
