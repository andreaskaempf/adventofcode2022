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
