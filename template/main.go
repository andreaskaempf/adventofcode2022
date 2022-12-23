// Advent of Code 2022, Day 2X
//
// Description:
//
// AK, 2X Dec 2022

package main

import (
	"fmt"
	//"strings"
)

func main() {

	// Read the input file
	fname := "sample.txt"
	//fname = "input.txt"
	for _, l := range readLines(fname) {
		fmt.Println(l)
	}
}
