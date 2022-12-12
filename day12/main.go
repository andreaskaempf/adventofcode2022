// Advent of Code 2022, Day 12
//
// Find the lowest cost path (i.e., shortest number of steps) through a terrain
// of letters, from point S to E, allowing 'increase' (e.g., next letter) of
// maximum 1. For Part 2, find the shortest path from any 'a' cell to 'E'.
//
// AK, 12/12/2022

package main

import (
	"fmt"

	// You need to install this: go get github.com/yourbasic/graph
	"github.com/yourbasic/graph"
)

// Main execution: run parts 1 and 2
func main() {

	// Read matrix from file, into a pseudo-matrix of byte rows
	// fname = "sample.txt"  // uncomment one file name
	fname := "input.txt"
	m := [][]byte{}
	for _, l := range readLines(fname) {
		m = append(m, []byte(l))
	}

	// Find S and E first, adjust their altitudes
	var S, E int
	nr := len(m)
	nc := len(m[0])
	for ri := 0; ri < nr; ri++ {
		for ci := 0; ci < nc; ci++ {
			if m[ri][ci] == 'S' {
				S = ri*nc + ci
				m[ri][ci] = 'a'
			} else if m[ri][ci] == 'E' {
				E = ri*nc + ci
				m[ri][ci] = 'z'
			}
		}
	}
	fmt.Println("S =", S, ", E =", E) // at least one should be non-zero

	// Build the graph: from each cell, add feasible steps to the right
	// and/or down, also in reverse direction if that is also feasible.
	// Here, feasible means "vertical climb" at most 1 (i.e., next cell is at
	// most one letter higher). Always use a cost (e.g., 1), otherwise
	// shortest path will be randomized (because there is no cost).
	g := graph.New(nr * nc) // enough capacity for all nodes
	for ri := 0; ri < nr; ri++ {
		for ci := 0; ci < nc; ci++ {

			// This starting node
			thisNode := ri*nc + ci
			thisLetter := int(m[ri][ci])
			assert(thisLetter >= 'a' && thisLetter <= 'z', "Bad letter")

			// Go right (and left from there) if possible
			maxDiff := 1
			if ci < nc-1 {
				nextNode := ri*nc + (ci + 1)
				nextLetter := int(m[ri][ci+1])
				if nextLetter-thisLetter <= maxDiff { // go right
					g.AddCost(thisNode, nextNode, 1)
				}
				if thisLetter-nextLetter <= maxDiff { // left
					g.AddCost(nextNode, thisNode, 1)
				}
			}

			// Go down/up if possible
			if ri < nr-1 {
				nextNode := (ri+1)*nc + ci
				nextLetter := int(m[ri+1][ci])
				if nextLetter-thisLetter <= maxDiff { // go down
					g.AddCost(thisNode, nextNode, 1)
				}
				if thisLetter-nextLetter <= maxDiff { // up
					g.AddCost(nextNode, thisNode, 1)
				}
			}
		}
	}

	// Part 1: calculate shortest path from S to E
	_, dist := graph.ShortestPath(g, S, E)
	fmt.Println("Part 1 (s/b 31, 490):", dist)

	// Part 2: find the shortest path from any 'a' cell to E
	var shortest int64
	for ri := 0; ri < nr; ri++ {
		for ci := 0; ci < nc; ci++ {
			if m[ri][ci] == 'a' {
				thisNode := ri*nc + ci
				_, dist := graph.ShortestPath(g, thisNode, E)
				if dist > 0 && (shortest == 0 || dist < shortest) {
					shortest = dist
				}
			}
		}
	}
	fmt.Println("Part 2 (s/b 29, 488):", shortest)
}
