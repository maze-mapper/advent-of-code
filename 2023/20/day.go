// Advent of Code 2023 - Day 20.
package day20

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type pulseFreq int

const (
	lowPulse pulseFreq = iota
	highPulse
)

type pulseInfo struct {
	pulse        pulseFreq
	source, dest string
}

func (pi *pulseInfo) String() string {
	var pulseStr string
	switch pi.pulse {
	case lowPulse:
		pulseStr = "low"
	case highPulse:
		pulseStr = "high"
	}
	return fmt.Sprintf("%s -%s-> %s", pi.source, pulseStr, pi.dest)
}

type module interface {
	sendPulse(pulseFreq, string, map[string]module, *[]pulseInfo)
	outputs() []string
}

type flipFlopModule struct {
	name               string
	destinationModules []string
	state              bool
}

func (m *flipFlopModule) sendPulse(pulse pulseFreq, _ string, modules map[string]module, queue *[]pulseInfo) {
	if pulse == highPulse {
		return
	}
	m.state = !m.state
	var outputPulse pulseFreq
	if m.state {
		outputPulse = highPulse
	} else {
		outputPulse = lowPulse
	}
	for _, dest := range m.destinationModules {
		pi := pulseInfo{
			pulse:  outputPulse,
			source: m.name,
			dest:   dest,
		}
		*queue = append(*queue, pi)
	}
}

func (m *flipFlopModule) outputs() []string {
	return m.destinationModules
}

type conjunctionModule struct {
	name               string
	destinationModules []string
	lastPulsesReceived map[string]pulseFreq
}

func (m *conjunctionModule) sendPulse(pulse pulseFreq, source string, modules map[string]module, queue *[]pulseInfo) {
	m.lastPulsesReceived[source] = pulse
	allHigh := true
	for _, v := range m.lastPulsesReceived {
		if v != highPulse {
			allHigh = false
			break
		}
	}
	var outputPulse pulseFreq
	if allHigh {
		outputPulse = lowPulse
	} else {
		outputPulse = highPulse
	}
	for _, dest := range m.destinationModules {
		pi := pulseInfo{
			pulse:  outputPulse,
			source: m.name,
			dest:   dest,
		}
		*queue = append(*queue, pi)
	}
}

func (m *conjunctionModule) outputs() []string {
	return m.destinationModules
}

type broadcastModule struct {
	name               string
	destinationModules []string
}

func (m *broadcastModule) sendPulse(pulse pulseFreq, _ string, modules map[string]module, queue *[]pulseInfo) {
	for _, dest := range m.destinationModules {
		pi := pulseInfo{
			pulse:  pulse,
			source: m.name,
			dest:   dest,
		}
		*queue = append(*queue, pi)
	}
}

func (m *broadcastModule) outputs() []string {
	return m.destinationModules
}

func parseData(data []byte) map[string]module {
	lines := strings.Split(strings.TrimSuffix(string(data), "\n"), "\n")
	modules := map[string]module{}
	for _, line := range lines {
		parts := strings.Split(line, " -> ")
		var m module
		var name string
		destinations := strings.Split(parts[1], ", ")

		switch {
		case parts[0] == "broadcaster":
			name = parts[0]
			m = &broadcastModule{
				name:               name,
				destinationModules: destinations,
			}
		case strings.HasPrefix(parts[0], "%"):
			name = strings.TrimPrefix(parts[0], "%")
			m = &flipFlopModule{
				name:               name,
				destinationModules: destinations,
				state:              false,
			}
		case strings.HasPrefix(parts[0], "&"):
			name = strings.TrimPrefix(parts[0], "&")
			m = &conjunctionModule{
				name:               name,
				destinationModules: destinations,
				lastPulsesReceived: map[string]pulseFreq{},
			}
		default:
			log.Fatalf("Unknown module %s", parts[0])
		}
		modules[name] = m
	}
	// Add inputs to conjunction modules.
	for name, module := range modules {
		for _, dest := range module.outputs() {
			dm, ok := modules[dest]
			if !ok {
				continue
			}
			switch dm.(type) {
			case *conjunctionModule:
				dm.(*conjunctionModule).lastPulsesReceived[name] = lowPulse
			}
		}
	}
	return modules
}

func part1(modules map[string]module) int {
	counts := map[pulseFreq]int{
		lowPulse:  0,
		highPulse: 0,
	}

	for i := 0; i < 1000; i++ {
		queue := []pulseInfo{
			{
				pulse:  lowPulse,
				source: "button",
				dest:   "broadcaster",
			},
		}
		for len(queue) > 0 {
			pi := queue[0]
			queue[0] = pulseInfo{}
			queue = queue[1:]
			counts[pi.pulse] += 1
			// fmt.Println(pi.String())
			if destMod, ok := modules[pi.dest]; ok {
				destMod.sendPulse(pi.pulse, pi.source, modules, &queue)
			}
		}
	}
	return counts[lowPulse] * counts[highPulse]
}

func part2(modules map[string]module) int {
	for i := 0; ; i++ {
		queue := []pulseInfo{
			{
				pulse:  lowPulse,
				source: "button",
				dest:   "broadcaster",
			},
		}
		// TODO: brute force doesn't work.
		// Instead need to find the LCM of when the four inputs to the comparator module send a high pulse.
		for len(queue) > 0 {
			pi := queue[0]
			queue[0] = pulseInfo{}
			queue = queue[1:]
			// fmt.Println(pi.String())
			if pi.dest == "rx" && pi.pulse == lowPulse {
				return i
			}
			if pi.dest == "rs" && pi.pulse == highPulse {
				fmt.Println(i, pi.String())
			}
			if destMod, ok := modules[pi.dest]; ok {
				destMod.sendPulse(pi.pulse, pi.source, modules, &queue)
			}
		}
	}
}

func Run(inputFile string) {
	data, err := os.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}
	modules := parseData(data)

	p1 := part1(modules)
	fmt.Println("Part 1:", p1)

	p2 := part2(modules)
	fmt.Println("Part 2:", p2)
}
