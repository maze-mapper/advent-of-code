// Advent of Code 2019 - Day 23
package day23

import (
	"fmt"
	"io/ioutil"
	"log"
	"sync"

	"adventofcode/2019/intcode"
)

// packet holds the data for a packet sent between computers
type packet struct {
	x, y int
}

// setupComputers creates n Intcode computers along with their input and output channels
func setupComputers(program []int, n int) ([]*intcode.Computer, []chan int, []chan int) {
	inputChannels := make([]chan int, n)
	outputChannels := make([]chan int, n)
	computers := make([]*intcode.Computer, n)

	for i := 0; i < n; i++ {
		chanIn := make(chan int, 256)
		chanOut := make(chan int, 256)
		computer := intcode.New(program)
		computer.SetNonBlocking(-1)
		computer.SetChanIn(chanIn)
		computer.SetChanOut(chanOut)

		inputChannels[i] = chanIn
		outputChannels[i] = chanOut
		computers[i] = computer

		// Send network addres
		chanIn <- i
	}

	return computers, inputChannels, outputChannels
}

func runComputers(computers []*intcode.Computer, inputChannels, outputChannels []chan int, chanDone chan packet) {
	n := len(computers)

	// Set up channels for packets to ensure that we do not interleave packets
	packetInputChannels := make([]chan packet, len(inputChannels))
	for i := 0; i < n; i++ {
		idx := i
		packetInputChannels[i] = make(chan packet, 128)
		go func() {
			for p := range packetInputChannels[idx] {
				inputChannels[idx] <- p.x
				inputChannels[idx] <- p.y
			}
		}()
	}

	// Goroutines for sending and receiving packets
	for i := 0; i < n; i++ {
		chanOut := outputChannels[i]
		go func() {
			var o int
			var addr, x, y int
			for out := range chanOut {
				switch o % 3 {
				case 0:
					addr = out
				case 1:
					x = out
				case 2:
					y = out
					p := packet{x: x, y: y}
					if addr == 255 {
						chanDone <- p
					} else {
						packetInputChannels[addr] <- p
					}
				}
				o += 1
			}
		}()
	}

	// Start computers as close to the same time as possible
	chanStart := make(chan struct{})
	for i := 0; i < n; i++ {
		computer := computers[i]
		go func() {
			<-chanStart
			computer.Run()
		}()
	}
	close(chanStart)
}

func part1(program []int) int {
	computers, inputChannels, outputChannels := setupComputers(program, 50)
	chanDone := make(chan packet, 255)
	runComputers(computers, inputChannels, outputChannels, chanDone)
	p := <-chanDone
	return p.y
}

func part2(program []int) int {
	var mutex = &sync.Mutex{}

	var natPacket packet
	var sendableNatPacket bool
	chanNat := make(chan packet, 255)
	go func() {
		for p := range chanNat {
			//			fmt.Println("Received NAT packet", p)
			natPacket = p
			mutex.Lock()
			sendableNatPacket = true
			mutex.Unlock()
		}
	}()

	computers, inputChannels, outputChannels := setupComputers(program, 50)
	runComputers(computers, inputChannels, outputChannels, chanNat)

	var lastY int
	for {
		// Determine if all computers are idle
		idle := true
		for _, c := range computers {
			// Consider idle if there have been a minimum number of reads with no input
			if !c.IdleFor(1000) {
				idle = false
				break
			}
		}

		if idle && sendableNatPacket {
			sentPacket := natPacket
			//			fmt.Println("Sending packet", sentPacket)
			if sentPacket.y == lastY {
				return sentPacket.y
			} else {
				lastY = sentPacket.y
			}

			mutex.Lock()
			sendableNatPacket = false
			mutex.Unlock()
			inputChannels[0] <- sentPacket.x
			inputChannels[0] <- sentPacket.y
		}
	}
	return 0
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
