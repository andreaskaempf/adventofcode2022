// Brute-force solution to Part 1

package main

import "fmt"

// Brute force: try all permutations of open valves
func tryAll() {

	// Get list of all valves that have flow
	valves := []int{}
	for i := 0; i < len(nodes); i++ {
		if nodes[i].flow > 0 {
			valves = append(valves, i)
		}
	}
	fmt.Println("Valves with flow:", valves)

	// Brute force: try simulating every possible permutation, to find the one
	// with the highest cumulative flow
	// simulate([]int{3, 1, 9, 7, 4, 2}) // the problem sample
	best := 0
	tried := 0
	Perm(valves, func(seq []int) { // 1.34e12 for 15 valves!
		y := simulate2(seq)
		if y > best {
			best = y
		}
		tried++
		if tried%100000000 == 0 {
			fmt.Printf("%.3f%% done, best = %d\n", float64(tried)/1.34e12*100.0, best)
		}
	})
	fmt.Println("Part 1:", tried, "tried, best found (s/b 1651, 1647) =", best)
}

// Calls f with each permutation of given list
// https://yourbasic.org/golang/generate-permutation-slice-string/
func Perm(a []int, f func([]int)) {
	perm(a, f, 0)
}

// Permute the values at index i to len(a)-1 (used by Perm above)
func perm(a []int, f func([]int), i int) {
	if i > len(a) {
		f(a)
		return
	}
	perm(a, f, i+1)
	for j := i + 1; j < len(a); j++ {
		a[i], a[j] = a[j], a[i]
		perm(a, f, i+1)
		a[i], a[j] = a[j], a[i]
	}
}

// Simulate one sequence of valves to open (list of indices), return the
// cumulative flow value. This is a shorter version, that does not
// simulate every time step.
func simulate2(seq []int) int {

	here := nodeAA                  // start with node AA
	cumValue := 0                   // cumulative value of the simulation
	t := 30                         // time left
	for i := 0; i < len(seq); i++ { // each valve opened in sequence
		dest := seq[i]               // the next valve's index
		dist := shortest(here, dest) // steps to get there
		if dist < 0 {
			fmt.Println("Infeasible path rejected")
			return 0
		}
		flow := nodes[dest].flow // flow from that node
		if flow == 0 {
			fmt.Println("Warning: zero flow")
		}
		t -= dist + 1 // time left
		if t <= 0 {   // stop if no more time left
			break
		}
		cumValue += flow * t // cumulative value of simulation to end
		//fmt.Println(dest, dist, flow, flow*t, cumValue)
		here = dest // next departure from this node
	}
	if cumValue > 1647 {
		fmt.Println("This sequence yielded", cumValue)
		fmt.Println(seq)
		panic("Simulation aborted")
	}
	return cumValue
}

// Simulate one sequence of valves to open (list of indices), return the
// cumulative flow value. Longer version, simulates every time step.
func simulate(seq []int) int {

	// Close all valves
	for i := 0; i < len(nodes); i++ {
		nodes[i].opened = false
	}

	// Set up for the beginning of the simulation, i.e., start moving
	// to first destination
	here := nodeAA         // start with node AA
	nextValve := 0         // next valve to open (start with first)
	dest := seq[nextValve] // index of the next valve node

	// Calculate distance from start to the first node
	dist := shortest(here, dest) //graph.ShortestPath(g, here, dest)
	moving := dist               // number of steps to next valve

	// Iterate for entire time
	totalFlow := 0             // total flow during the simulation
	curFlow := 0               // the total flow of all open valves
	for t := 1; t <= 30; t++ { // run for 30 mins

		// Add flow for any valves already open
		totalFlow += curFlow
		//fmt.Printf("Flow at t = %d: current %d, cumulative = %d\n", t, curFlow, totalFlow)

		// If in the process of moving to another valve, move one step
		if moving > 0 {
			//fmt.Printf("Moving toward %s at t = %d\n", nodes[dest].id, t)
			moving--
			continue
		}

		// If arrived at valve to open, open the valve,  prepare for next valve
		if moving == 0 {

			// Open this valve
			//fmt.Printf("Opening valve %s at t = %d\n", nodes[dest].id, t)
			nodes[dest].opened = true
			curFlow += nodes[dest].flow
			here = dest // this is now the currrent node

			// Prepare the next valve
			moving = -1
			nextValve++
			if nextValve >= len(seq) { // end of sequence reached
				continue
			}

			dest = seq[nextValve]        // index of the next valve node
			dist := shortest(here, dest) //graph.ShortestPath(g, here, dest)
			moving = dist                // number of steps to next valve
		}
	}

	// Show answer
	//fmt.Println("Total flow (s/b 1651) =", totalFlow)
	return totalFlow
}
