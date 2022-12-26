// Advent of Code 2022, Day 16
//
// Given a network (graph) of closed "valves", each with a certain flow rate,
// find the sequence of opening the valves (takes one minute, plus one minute
// per step to get there) that yields the highest possible total flow during a
// 30-minute period. Used brute force for Part 1, but need to revisit this to
// formulate a true optimization and complete Part 2.
//
// AK, 16 and 26 Dec 2022

package main

import (
	"fmt"
	"strings"

	"github.com/yourbasic/graph"
)

// A node in the graph
type Node struct {
	id     string
	index  int // zero-based index
	flow   int
	connTo []string
	opened bool
}

// Global list of nodes
var nodes []Node

// Index of node AA
var nodeAA int

// Global pointer to the graph
var g *graph.Mutable

// For memoization of graph shortest distances
type Pair struct {
	a, b int
}

var distances map[Pair]int

func main() {

	// Initialize for memoization of distances
	distances = map[Pair]int{}

	// Read the input file into a graph
	fname := "sample.txt"
	//fname = "input.txt"
	readInput(fname)

	// Try simulating some sample sequences
	//sampleSeq := []int{3, 1, 9, 7, 4, 2} // from problem
	//fmt.Println("Simulate:", sampleSeq, simulate2(sampleSeq))
	//s2 := []int{11, 14, 59, 6, 38, 19, 44, 53, 39, 60, 54, 24, 57, 12}
	//fmt.Println("Simulate:", simulate(s2), simulate2(s2))

	// Optimization is not working yet, only looks one-head, produces
	// decent solution but not the optimum
	// TODO: Get this working, by implementing some look-ahead, but
	// not all permutations
	fmt.Println("Part 1 (s/b 1651, 1647):", optimize3())

	// Brute force: works fine and quicly on the problem sample
	// (720 permutations), but takes about 20 hours on the full
	// problem input
	//tryAll()
}

func optimize3() int {

	// Get list of (indices of) all valves that have flow, since these
	// are the only ones we care about
	valves := []int{}
	for i := 0; i < len(nodes); i++ {
		if nodes[i].flow > 0 {
			valves = append(valves, i)
			fmt.Printf("Valve %s: flow %d\n", nodes[i].id, nodes[i].flow)
		}
	}

	// Optimize from starting node with all valves that have flow as
	// candidates for the next move, and return the best result
	return optimize3a(nodeAA, valves, 1)
}

// One recursive iteration of the optimization: given that you are
// at node "here" at time "t", try to open valves in list "candidates"
// and find the highest cumulative flow.
func optimize3a(here int, candidates []int, t int) int {

	// Out of time!
	if t > 30 {
		return 0
	}

	// Remove this node we're at from the list of possible
	// candidates for future moves
	candidates = remove(candidates, here) // makes a copy

	// Try each candidate, pruning the ones that would take too
	// long to reach, and re-using memoized solutions from previous
	// iterations
	best := 0
	for _, vi := range candidates {

		// Don't bother if not enough time to get there, i.e.,
		// by the time you got there, you could not open the valve
		// in time to get any flow
		dist := shortest(here, vi) // time to get to the next valve
		if (30-t)-dist < 1 {
			continue
		}

		// Get the value of opening this candidate valve now, until the
		// end of the simulation (takes one time step to open)
		thisValveFlow := nodes[vi].flow * (30 - t - dist)
		//fmt.Printf("Opening valve %s (flow %d) at t=%d, creates %d of flow\n",
		//		nodes[vi].id, nodes[vi].flow, t, thisValveFlow)

		// Simulate moving from this node to the alternative,
		// and optimizing from there this node
		newCand := remove(candidates, vi) // makes a copy
		o := optimize3a(vi, newCand, t+dist+1)

		// Is this the best found?
		if thisValveFlow+o > best {
			best = thisValveFlow + o
		}
	}

	return best
}

// Shortest distance between two nodes, with memoization
func shortest(a, b int) int {
	pair := Pair{a, b}
	dist, ok := distances[pair]
	if !ok {
		//fmt.Println("Calculating", a, b)
		_, d := graph.ShortestPath(g, a, b)
		dist = int(d)
		distances[pair] = dist
	}
	return dist
}

// Parse input, create graph
func readInput(filename string) {

	// Read the file, one line per node
	lines := readLines(filename)
	fmt.Println(len(lines), "lines read")

	// Process each line, build list of nodes
	// Valve CC has flow rate=2; tunnels lead to valves DD, BB
	nodes = []Node{} // initialize global variable
	nodeIndex := map[string]int{}
	for _, l := range lines {
		words := strings.Split(l, " ")
		rate := atoi(words[4][5:(len(words[4]) - 1)])
		n := Node{id: words[1], flow: rate}
		nodeIndex[n.id] = len(nodes)  // index of this node
		for _, c := range words[9:] { // list of connected nodes
			if c[len(c)-1] == ',' { // remove comma
				c = c[:len(c)-1]
			}
			n.connTo = append(n.connTo, strings.TrimSpace(c))
		}
		nodes = append(nodes, n)
	}

	// Find the first node AA (global variable)
	nodeAA = -1 // index of node we're at, start at AA
	for i := 0; i < len(nodes); i++ {
		if nodes[i].id == "AA" {
			nodeAA = i
		}
	}
	if nodeAA < 0 {
		panic("Node AA not found!")
	}

	// Great a graph for the nodes
	g = graph.New(len(nodes))
	for ni := 0; ni < len(nodes); ni++ {
		n := nodes[ni]
		for _, c := range n.connTo {
			ci := nodeIndex[c]
			g.AddBothCost(ni, ci, 1) // add bidirectional connection, weight 1
		}
	}
}

// Remove element from a list, returning new copy
func remove(elems []int, e int) []int {
	res := []int{}
	for _, x := range elems {
		if x != e {
			res = append(res, x)
		}
	}
	return res
}
