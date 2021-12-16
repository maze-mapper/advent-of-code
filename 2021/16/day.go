// Advent of Code 2021 - Day 16
package day16

import (
	"bytes"
	"encoding/hex"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strings"
)

// ========================

type BitReader struct {
	reader io.ByteReader
	byte   byte
	offset byte
}

func New(r io.ByteReader) *BitReader {
	return &BitReader{r, 0, 0}
}

func (r *BitReader) ReadBit() (bool, error) {
	if r.offset == 8 {
		r.offset = 0
	}
	if r.offset == 0 {
		var err error
		if r.byte, err = r.reader.ReadByte(); err != nil {
			return false, err
		}
	}
	bit := (r.byte & (0x80 >> r.offset)) != 0
	r.offset++
	return bit, nil
}

func (r *BitReader) ReadUint(nbits int) (uint, error) {
	var result uint
	for i := nbits - 1; i >= 0; i-- {
		bit, err := r.ReadBit()
		if err != nil {
			return 0, err
		}
		if bit {
			result |= 1 << uint(i)
		}
	}
	return result, nil
}

func (r *BitReader) ReadByteSlice(nbits int) ([]byte, error) {
	var result []byte
	for i := nbits - 1; i >= 0; i-- {
		bit, err := r.ReadBit()
		if err != nil {
			return []byte{}, err
		}
		if bit {
			result = append(result, byte(1))
		} else {
			result = append(result, byte(0))
		}
	}
	return result, nil
}

// ==========================

func hexToBytes(s string) []byte {
	data, err := hex.DecodeString(s)
	if err != nil {
		log.Fatal(err)
	}
	return data
}

func parseData(data []byte) []byte {
	s := strings.TrimSuffix(string(data), "\n")
	b := hexToBytes(s)
	return b
}

func readHeader(r *BitReader) (uint, uint) {
	version, _ := r.ReadUint(3)
	typeID, _ := r.ReadUint(3)
	return version, typeID
}

func readPacket(r *BitReader, versionSum *uint) (int, uint) {
	bitsRead := 0
	var value uint

	version, err := r.ReadUint(3)
	bitsRead += 3
	if err != nil {
		return bitsRead, value
	}
	typeID, _ := r.ReadUint(3)
	bitsRead += 3
	*versionSum += version

	switch typeID {

	// Literal
	case 4:
		var prefixBit uint = 1
		var literal uint
		for prefixBit != 0 {
			prefixBit, _ = r.ReadUint(1)
			n, _ := r.ReadUint(4)
			bitsRead += 5

			literal = literal << 4
			literal += n
		}
		value = literal

	// Operator
	default:
		lengthTypeID, _ := r.ReadUint(1)
		bitsRead += 1

		subPacketValues := []uint{}

		switch lengthTypeID {

		case 0:
			lengthInBits, _ := r.ReadUint(15)
			bitsRead += 15

			newBitsRead := 0
			for newBitsRead < int(lengthInBits) {
				read, p := readPacket(r, versionSum)
				newBitsRead += read
				subPacketValues = append(subPacketValues, p)
			}
			bitsRead += newBitsRead

		case 1:
			lengthInPackets, _ := r.ReadUint(11)
			bitsRead += 11
			for packet := uint(0); packet < lengthInPackets; packet++ {
				read, p := readPacket(r, versionSum)
				bitsRead += read
				subPacketValues = append(subPacketValues, p)
			}
		}

		// Do something with subpackets
		switch typeID {

		// Sum
		case 0:
			for _, p := range subPacketValues {
				value += p
			}

		// Product
		case 1:
			value = 1
			for _, p := range subPacketValues {
				value *= p
			}

		// Minimum
		case 2:
			min := ^uint(0)
			for _, p := range subPacketValues {
				if p < min {
					min = p
				}
			}
			value = min

		// Maximum
		case 3:
			max := uint(0)
			for _, p := range subPacketValues {
				if p > max {
					max = p
				}
			}
			value = max

		// Greater than
		case 5:
			if subPacketValues[0] > subPacketValues[1] {
				value = 1
			}

		// Less than
		case 6:
			if subPacketValues[0] < subPacketValues[1] {
				value = 1
			}

		// Equal
		case 7:
			if subPacketValues[0] == subPacketValues[1] {
				value = 1
			}
		}

	}
	return bitsRead, value
}

func decode(b []byte) (uint, uint) {
	r := New(bytes.NewBuffer(b))
	var versionSum uint
	_, val := readPacket(r, &versionSum)
	return versionSum, val
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	b := parseData(data)

	versionSum, val := decode(b)

	fmt.Println("Part 1:", versionSum)

	fmt.Println("Part 2:", val)
}
