// Advent of Code 2022, Day 0X
//
// Description:
//
// AK, 1 Dec 2022

package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {

	// Read the input file
	filename := "sample.txt"
	//filename := "input.txt"
	data, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(data), "\n")
	fmt.Println(len(lines), "lines read")

	// Process each line
	for _, l := range lines {
		fmt.Println(l) // TODO
	}
}
