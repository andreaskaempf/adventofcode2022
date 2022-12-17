package main

import (
	"fmt"
	//"github.com/yourbasic/graph"
	"sort"
)

func optimize2() {

	// Get list of all valves that have flow
	valves := []int{}
	for i := 0; i < len(nodes); i++ {
		if nodes[i].flow > 0 {
			valves = append(valves, i)
		}
	}
	sort.Slice(valves, func(i, j int) bool {
		return nodes[valves[i]].flow < nodes[valves[j]].flow
	})
	fmt.Print("Valves with flow:")
	for _, v := range valves {
		n := nodes[v]
		fmt.Printf("%s %d, ", n.id, n.flow)
	}
	fmt.Println("")

	// Try removing each combination of 3
	best := 0
	for i := 0; i < len(valves); i++ {
		for j := 0; j < len(valves); j++ {
			for k := 0; k < len(valves); k++ {
				if i != j && i != k && j != k {
					v1 := valves[i]
					v2 := valves[j]
					v3 := valves[k]
					y := simulate2([]int{v1, v2, v3})
					if y > best {
						best = y
						fmt.Println(nodes[v1].id, nodes[v2].id, nodes[v3].id, y)
					}
				}
			}
		}
	}
}

func optimize() {

	// Iterate until no open valves left
	here := nodeAA             // index of the node we are at, start with AA
	totalFlow := 0             // total flow during the simulation
	curFlow := 0               // current flow from valves already opened
	seq := []string{}          // sequence of valves opened
	moving := -1               // number of steps to next valve
	toOpen := -1               // index of valve to be opened
	for t := 1; t <= 30; t++ { // run for 30 mins

		// Add flow for any valves already open
		totalFlow += curFlow
		fmt.Printf("Flow at t = %d: current %d, cumulative = %d\n", t, curFlow, totalFlow)

		// If in the process of moving to another valve, move one step
		if moving > 0 {
			fmt.Printf("Moving toward %s at t = %d\n", nodes[toOpen].id, t)
			moving--
			continue
		}

		// If arrived at valve to open, open the valve
		if moving == 0 {
			fmt.Printf("Opening valve %s at t = %d\n", nodes[toOpen].id, t)
			nodes[toOpen].opened = true
			curFlow += nodes[toOpen].flow
			here = toOpen // this is now the currrent node
			moving = -1
			toOpen = -1
			continue
		}

		// Find all open valves with flow > 0
		open := []int{}
		for i := 0; i < len(nodes); i++ {
			if nodes[i].flow > 0 && !nodes[i].opened {
				open = append(open, i)
			}
		}

		// Don't try to find any more open valves (but continue the simulation)
		if len(open) == 0 {
			continue
		}

		// Find then next valve to open, by finding the one with the
		// highest expected value to end of the simulation
		var payoff int   // best payoff among remaining valves
		var best int     // index of best valve found
		var bestDist int // length of path of best solution
		for _, dest := range open {
			dist := shortest(here, dest)
			po := nodes[dest].flow * (30 - t - dist)
			if po > payoff {
				payoff = po
				best = dest
				bestDist = int(dist)
			}
		}

		// Mark the next valve to open
		fmt.Printf("Best next destination = %s, payoff %d, steps = %d\n",
			nodes[best].id, payoff, bestDist)
		moving = bestDist //number of steps to get there
		toOpen = best     // index of the valve to open
		seq = append(seq, nodes[best].id)
	}

	// Show Part 1 answer
	fmt.Println("Total flow (s/b 1651, 1647) =", totalFlow)
	fmt.Println("Sequence (s/b DD, BB, JJ, EE, HH, CC):", seq)
}

func optimize1() {

	// start at t = 1, at node 0, with all valves closed (0)
	// pick each possible node, in flow order
	// and try opening that

}

func optimize_(state []int, t int) { // state of each valve: 1 = open, 0 = closed

	// Iterate until no open valves left
	here := 0                  // index of the node we are at
	totalFlow := 0             // total flow during the simulation
	curFlow := 0               // current flow from valves already opened
	seq := []string{}          // sequence of valves opened
	moving := -1               // number of steps to next valve
	toOpen := -1               // index of valve to be opened
	for t := 1; t <= 30; t++ { // run for 30 mins

		// Add flow for any valves already open
		totalFlow += curFlow
		fmt.Printf("Flow at t = %d: current %d, cumulative = %d\n", t, curFlow, totalFlow)

		// If in the process of moving to another valve, move one step
		if moving > 0 {
			fmt.Printf("Moving toward %s at t = %d\n", nodes[toOpen].id, t)
			moving--
			continue
		}

		// If arrived at valve to open, open the valve
		if moving == 0 {
			fmt.Printf("Opening valve %s at t = %d\n", nodes[toOpen].id, t)
			nodes[toOpen].opened = true
			curFlow += nodes[toOpen].flow
			here = toOpen // this is now the currrent node
			moving = -1
			toOpen = -1
			continue
		}

		// Find all open valves with flow > 0
		open := []int{}
		for i := 0; i < len(nodes); i++ {
			if nodes[i].flow > 0 && !nodes[i].opened {
				open = append(open, i)
			}
		}

		// Don't try to find any more open valves (but continue the simulation)
		if len(open) == 0 {
			continue
		}

		// Find then next valve to open, by finding the one with the
		// highest expected value to end of the simulation
		var payoff int   // best payoff among remaining valves
		var best int     // index of best valve found
		var bestDist int // length of path of best solution
		for _, dest := range open {
			dist := shortest(here, dest)
			po := nodes[dest].flow * (30 - t - dist)
			if po > payoff {
				payoff = po
				best = dest
				bestDist = int(dist)
			}
		}

		// Mark the next valve to open
		fmt.Printf("Best next destination = %s, payoff %d, steps = %d\n",
			nodes[best].id, payoff, bestDist)
		moving = bestDist //number of steps to get there
		toOpen = best     // index of the valve to open
		seq = append(seq, nodes[best].id)
	}

	// Show Part 1 answer
	fmt.Println("Total flow (s/b 1651, 1647) =", totalFlow)
	fmt.Println("Sequence (s/b DD, BB, JJ, EE, HH, CC):", seq)
}
