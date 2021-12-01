// Advent of Code 2015 - Day 4
package day4

import (
	"crypto/md5"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"strconv"
)

func solve(prefix, target string) {
	hashPrefix := ""
	targetLen := len(target)

	n := 0
	for hashPrefix != target {
		n++
		h := md5.New()
		io.WriteString(h, prefix)
		io.WriteString(h, strconv.Itoa(n))
		hex := fmt.Sprintf("%x", h.Sum(nil))
		hashPrefix = hex[:targetLen]
	}

	fmt.Println(n)
}

func Run(inputFile string) {
	data, err := ioutil.ReadFile(inputFile)
	if err != nil {
		log.Fatal(err)
	}

	prefix := string(data)
	solve(prefix, "00000")
	solve(prefix, "000000")
}
