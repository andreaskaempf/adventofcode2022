// Advent of Code 2022, Day 08
//
// Description:
//
// AK, 8 Dec 2022

package main

import (
	"fmt"
	//"strings"
)

func main() {

	// Read the input file
	//lines := readLines("sample.txt")
	lines := readLines("input.txt")
	fmt.Println(len(lines), "lines read")

	// Part 1: count how many trees are "visible"
	nr := len(lines)
	nc := len(lines[0])
	nvis := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			if isVisible(r, c, lines) {
				nvis++
			}
		}
	}
	fmt.Println("Part 1 (s/b 21):", nvis)

	// Part 2: maximum scenic score
	// 560 too low, also 14400
	score := 0
	for r := 0; r < nr; r++ {
		for c := 0; c < nc; c++ {
			ss := scenicScore(r, c, lines)
			if ss > score {
				//fmt.Printf("Max %d found at %d,%d\n", ss, r, c)
				score = ss
			}
		}
	}
	fmt.Println("Part 2 (s/b 8):", score)
	fmt.Println("Score 1,2 =", scenicScore(1, 2, lines), " (s/b 4)")
	fmt.Println("Score 3,2 =", scenicScore(3, 2, lines), " (s/b 8)")
}

// For Part 1: Is tree in the given position "visible"
func isVisible(r, c int, forest []string) bool {

	// Trees on the edges are always visible
	nr := len(forest)
	nc := len(forest[0])
	if r == 0 || r == nr-1 || c == 0 || c == nc-1 {
		return true
	}

	// Otherwise, look up & down left & right, to see
	// if there is an "opening" in any direction, i.e.,
	// all trees in that direction are lower than this one
	x := forest[r][c] // this "tree"

	visLeft := true
	for i := 0; i < c; i++ { // left
		if forest[r][i] >= x {
			visLeft = false
		}
	}

	visRight := true
	for i := c + 1; i < nc; i++ { // right
		if forest[r][i] >= x {
			visRight = false
		}
	}

	visAbove := true
	for i := 0; i < r; i++ {
		if forest[i][c] >= x { // above
			visAbove = false
		}
	}

	visBelow := true
	for i := r + 1; i < nr; i++ {
		if forest[i][c] >= x { // below
			visBelow = false
		}
	}

	// A tree is visible if not blocked from any direction
	return visLeft || visRight || visAbove || visBelow
}

// For Part 1: Is tree in the given position "visible"
func scenicScore(r, c int, forest []string) int {

	// Count how many trees are visible in each direction
	nr := len(forest)
	nc := len(forest[0])
	/*if r == 0 || r == nr-1 || c == 0 || c == nc-1 {
		return 0
	}*/

	seq := []byte{}
	for i := c; i >= 0; i-- { // left
		seq = append(seq, forest[r][i])
	}
	visLeft := nVisible(seq)

	seq = []byte{}
	for i := c; i < nc; i++ { // right
		seq = append(seq, forest[r][i])
	}
	visRight := nVisible(seq)

	seq = []byte{}
	for i := r; i >= 0; i-- { // up
		seq = append(seq, forest[i][c])
	}
	visUp := nVisible(seq)

	seq = []byte{}
	for i := r; i < nr; i++ { // down
		seq = append(seq, forest[i][c])
	}
	visDown := nVisible(seq)

	// Return product to get score
	score := visLeft * visRight * visUp * visDown
	fmt.Printf("%d,%d: left = %d, right = %d, up = %d, down = %d => %d\n",
		r, c, visLeft, visRight, visUp, visDown, score)
	return score
}

// Count the number of trees visible in this sequence,
// i.e., each >= to the last
// TODO: can you see trees that are higher, but a few beyond the last one?
func nVisible(trees []byte) int {
	n := 0
	if len(trees) < 2 {
		return 0
	}
	highest := trees[1]
	for i := 1; i < len(trees); i++ {
		if trees[i] > trees[0] {
			n++
			break
		}
		if i == 1 || trees[i] >= highest {
			n++
		}
		if trees[i] > highest {
			highest = trees[i]
		}
	}
	return n
}
