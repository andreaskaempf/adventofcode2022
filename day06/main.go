// Advent of Code 2022, Day 06
//
// Look for first block of 4 (Part 1) or 14 (Part 2) non-repeating
// characters in a string.
//
// AK, 6 Dec 2022

package main

import (
	"fmt"
	"io/ioutil"
)

func main() {

	// Sample values, with expected values
	samples := []string{"mjqjpqmgbljsphdztnvjfqwrcgsmlb", // 4
		"bvwbjplbgvbhsrlpgdmjqwftvncz",      // 5
		"nppdvjthqldpwncqszvftbrmjlhg",      // 6
		"nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg", // 10
		"zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw"}  // 11

	// Read input, append to list of samples
	data, _ := ioutil.ReadFile("input.txt")
	samples = append(samples, string(data))

	// Part 1: Process each line, look for markers of length 4
	fmt.Println("Part 1")
	for _, msg := range samples {
		fmt.Println(marker(msg, 4))
	}

	// Part 2: Process each line, look for markers of length 14
	fmt.Println("\nPart 2")
	for _, msg := range samples {
		fmt.Println(marker(msg, 14))
	}
}

// Find the position of the "marker", i.e., a block where n chars
// are all different
func marker(s string, n int) int {
	for i := n - 1; i < len(s); i++ {
		substr := s[i-n+1 : i+1]
		if !duplicates(substr) {
			return i + 1
		}
	}
	return -1 // no marker found (should never happen)
}

// Check if a string contains any duplicate characters
func duplicates(s string) bool {
	chars := map[byte]int{}
	for i := 0; i < len(s); i++ {
		c := s[i]
		chars[c] += 1     // remember we found this character
		if chars[c] > 1 { // duplicate found
			return true
		}
	}
	return false
}
