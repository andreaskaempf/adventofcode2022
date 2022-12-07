// Advent of Code 2022, Day 0X
//
// Description:
//
// AK, X Dec 2022

package main

import (
	"fmt"
	//"strings"
)

func main() {

	// Read the input file
	lines := readLines("sample.txt")
	// lines := readFile("input.txt")
	fmt.Println(len(lines), "lines read")

	// Process each line
	for _, l := range lines {
		fmt.Println(l) // TODO
	}
}
