// Intcode computer
package intcode

import (
	"log"
	"strconv"
	"strings"
)

// Opcodes known by the Intcode computer
const (
	OpcodeAdd                = 1
	OpcodeMultiply           = 2
	OpcodeInput              = 3
	OpcodeOutput             = 4
	OpcodeJumpIfTrue         = 5
	OpcodeJumpIfFalse        = 6
	OpcodeLessThan           = 7
	OpcodeEquals             = 8
	OpcodeRelativeBaseOffset = 9
	OpcodeHalt               = 99
)

// Parameter modes
const (
	ParameterModePosition  = 0
	ParameterModeImmediate = 1
	ParameterModeRelative  = 2
)

// Computer is an Intcode computer
type Computer struct {
	program                      []int
	instructionPtr, relativeBase int
	chanIn                       <-chan int
	chanOut                      chan<- int
	blocking                     bool
	defaultInput                 int
	idleFor                      int
}

// readOpcode returns the last two digits of an instruction which represent the opcode
func readOpcode(number int) int {
	return number % 100
}

// getDigits extracts the digits from a number
func getDigits(n int, digits *[]int) {
	if n < 10 {
		*digits = append(*digits, n)
	} else {
		getDigits(n/10, digits)
		*digits = append(*digits, n%10)
	}
}

// readParameterMode returns a slice of the parameter modes from a number
func readParameterMode(number int) []int {
	digits := []int{}
	getDigits(number, &digits)

	// Reverse the digits so that the modes are in parameter order
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}

	// Copy to mode to ensure correct length
	modeLength := 3
	modes := make([]int, modeLength)
	copy(modes, digits)

	return modes
}

// copyProgram returns a copy of the given program
func copyProgram(program []int) []int {
	newProgram := make([]int, len(program))
	copy(newProgram, program)
	return newProgram
}

// New creates a computer with the given program
func New(program []int) *Computer {
	return &Computer{
		program:  copyProgram(program),
		chanOut:  make(chan int),
		blocking: true,
	}
}

// SetNonBlocking sets the computer to use a default input when no value is received from the input channel
func (c *Computer) SetNonBlocking(input int) {
	c.blocking = false
	c.defaultInput = input
}

// setChanIn sets the input channel
func (c *Computer) SetChanIn(ch <-chan int) {
	c.chanIn = ch
}

// setChanOut sets the output channel
func (c *Computer) SetChanOut(ch chan<- int) {
	c.chanOut = ch
}

// getParameters returns the next n address locations from the instruction pointer
func (c *Computer) getParameters(n int) []int {
	params := make([]int, n)
	for i := 0; i < n; i++ {
		params[i] = c.program[c.instructionPtr+1+i]
	}
	return params
}

// checkAddress checks that an address is valid and if not grows the memory to accomodate it
func (c *Computer) checkAddress(address int) {
	if address >= len(c.program) {
		newProgram := make([]int, address+1)
		copy(newProgram, c.program)
		c.program = newProgram
	}
}

// getParameterValueAddress returns the appropriate value and address for a parameter given the parameter mode
func (c *Computer) getParameterValue(parameter, mode int) (int, int) {
	var value, address int

	switch mode {

	case ParameterModePosition:
		address = parameter
		c.checkAddress(address)
		value = c.program[address]

	case ParameterModeImmediate:
		address = -1
		value = parameter

	case ParameterModeRelative:
		address = parameter + c.relativeBase
		c.checkAddress(address)
		value = c.program[address]

	default:
		log.Fatalf("Unknown parameter mode %d", mode)
	}

	return value, address
}

// add will execute the actions for OpcodeAdd
func (c *Computer) add(modes []int) {
	params := c.getParameters(3)
	valueA, _ := c.getParameterValue(params[0], modes[0])
	valueB, _ := c.getParameterValue(params[1], modes[1])
	_, destAddress := c.getParameterValue(params[2], modes[2])
	c.program[destAddress] = valueA + valueB
	c.instructionPtr += 4
}

// multiply will execute the actions for OpcodeMultiply
func (c *Computer) multiply(modes []int) {
	params := c.getParameters(3)
	valueA, _ := c.getParameterValue(params[0], modes[0])
	valueB, _ := c.getParameterValue(params[1], modes[1])
	_, destAddress := c.getParameterValue(params[2], modes[2])
	c.program[destAddress] = valueA * valueB
	c.instructionPtr += 4
}

// input will execute the actions for OpcodeInput
func (c *Computer) input(modes []int) {
	params := c.getParameters(1)
	_, destAddress := c.getParameterValue(params[0], modes[0])
	var in int
	if c.blocking {
		in = <-c.chanIn
	} else {
		select {
		case val := <-c.chanIn:
			in = val
			c.idleFor = 0
		default:
			in = c.defaultInput
			c.idleFor += 1
		}
	}
	c.program[destAddress] = in
	c.instructionPtr += 2
}

// output will execute the actions for OpcodeOutput
func (c *Computer) output(modes []int) {
	params := c.getParameters(1)
	value, _ := c.getParameterValue(params[0], modes[0])
	c.chanOut <- value
	c.instructionPtr += 2
}

// jumpIfTrue will execute the actions for OpcodeJumpIfTrue
func (c *Computer) jumpIfTrue(modes []int) {
	params := c.getParameters(2)
	valueA, _ := c.getParameterValue(params[0], modes[0])
	if valueA != 0 {
		valueB, _ := c.getParameterValue(params[1], modes[1])
		c.instructionPtr = valueB
	} else {
		c.instructionPtr += 3
	}
}

// jumpIfFalse will execute the actions for OpcodeJumpIfFalse
func (c *Computer) jumpIfFalse(modes []int) {
	params := c.getParameters(2)
	valueA, _ := c.getParameterValue(params[0], modes[0])
	if valueA == 0 {
		valueB, _ := c.getParameterValue(params[1], modes[1])
		c.instructionPtr = valueB
	} else {
		c.instructionPtr += 3
	}
}

// lessThan will execute the actions for OpcodeLessThan
func (c *Computer) lessThan(modes []int) {
	params := c.getParameters(3)
	valueA, _ := c.getParameterValue(params[0], modes[0])
	valueB, _ := c.getParameterValue(params[1], modes[1])
	_, destAddress := c.getParameterValue(params[2], modes[2])
	if valueA < valueB {
		c.program[destAddress] = 1
	} else {
		c.program[destAddress] = 0
	}
	c.instructionPtr += 4
}

// equals will execute the actions for OpcodeEquals
func (c *Computer) equals(modes []int) {
	params := c.getParameters(3)
	valueA, _ := c.getParameterValue(params[0], modes[0])
	valueB, _ := c.getParameterValue(params[1], modes[1])
	_, destAddress := c.getParameterValue(params[2], modes[2])
	if valueA == valueB {
		c.program[destAddress] = 1
	} else {
		c.program[destAddress] = 0
	}
	c.instructionPtr += 4
}

// relativeBaseOffset will execute the actions for OpcodeRelativeBaseOffset:
func (c *Computer) relativeBaseOffset(modes []int) {
	params := c.getParameters(1)
	value, _ := c.getParameterValue(params[0], modes[0])
	c.relativeBase += value
	c.instructionPtr += 2
}

// Run will run the Intcode computer until it halts
func (c *Computer) Run() {
	for {
		opcode := readOpcode(c.program[c.instructionPtr])
		modes := readParameterMode(c.program[c.instructionPtr] / 100)
		switch opcode {

		case OpcodeAdd:
			c.add(modes)

		case OpcodeMultiply:
			c.multiply(modes)

		case OpcodeInput:
			c.input(modes)

		case OpcodeOutput:
			c.output(modes)

		case OpcodeJumpIfTrue:
			c.jumpIfTrue(modes)

		case OpcodeJumpIfFalse:
			c.jumpIfFalse(modes)

		case OpcodeLessThan:
			c.lessThan(modes)

		case OpcodeEquals:
			c.equals(modes)

		case OpcodeRelativeBaseOffset:
			c.relativeBaseOffset(modes)

		case OpcodeHalt:
			close(c.chanOut)
			return

		default:
			log.Fatalf("Unknown opcode %d", opcode)

		}
	}

}

// Program returns the current state of the Incode program
func (c *Computer) Program() []int {
	return c.program
}

// IdleFor returns whether the computer has not received any input for n read attempts
func (c *Computer) IdleFor(n int) bool {
	return c.idleFor >= n
}

// ReadProgram returns a slice of ints
func ReadProgram(data []byte) []int {
	digits := strings.Split(
		strings.TrimSuffix(string(data), "\n"), ",",
	)
	program := make([]int, len(digits))
	for i, digit := range digits {
		if num, err := strconv.Atoi(digit); err == nil {
			program[i] = num
		} else {
			log.Fatal(err)
		}
	}
	return program
}
