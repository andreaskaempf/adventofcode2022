// Advent of Code 2022, Day 16
//
// Given a network (graph) of closed "valves", each with a certain flow rate,
// find the sequence of opening the valves (takes one minute, plus one minute
// per step to get there) that yields the highest possible total flow during a
// 30-minute period. Used brute force for Part 1, but need to revisit this to
// formulate a true optimization and complete Part 2.
//
// AK, 16 Dec 2022

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
	fname = "input.txt"
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
	//optimize()

	// Brute force: works fine and quicly on the problem sample
	// (720 permutations), but takes about 20 hours on the full
	// problem input
	tryAll()
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
