// Advent of Code 2022, Day 11
//
// Description:
//
// AK, 11 Dec 2022

package main

import (
	"fmt"
	"sort"
	"strings"
)

// State of a monkey
type Monkey struct {
	id              int      // number of this monkey (0...)
	items           []int64  // worry levels of items it holds
	operation       []string // old*old, old+n, old*n
	test            int64    // divisible by this number
	ifTrue, ifFalse int      // monkey to throw to
	inspections     int      // number of inspections made
}

func main() {

	// Set this to true for Part 1, false for Part 2
	part1 := false

	// Read and parse "monkeys" from input file
	monkeys := readMonkeys("sample.txt")
	//monkeys := readMonkeys("input.txt")
	for _, m := range monkeys {
		fmt.Println(m)
	}

	// Do the simulation for 20 or 10k rounds
	niters := ifElse(part1, 20, 10000)
	for round := 1; round <= niters; round++ {

		fmt.Println("\n*** Round", round)

		// Do each monkey
		for mi := 0; mi < len(monkeys); mi++ {

			// Do each item the monkey holds
			m := &monkeys[mi]
			for len(m.items) > 0 {

				// Apply operation to the worry level
				m.inspections++
				i := 0
				wl := m.items[i] // current worry level
				wl = applyOperation(wl, m.operation)

				// Now integer-divide by 3 for Part 1
				if part1 {
					wl /= 3
				} else {
					if wl > 0 {
						wl /= 3
					} else {
						wl /= -100000000
					}
				}
				fmt.Println("    w/l now", wl)

				// Apply test to determine who to throw to
				dest := m.ifTrue // assume divisible by test
				if wl%m.test != 0 {
					dest = m.ifFalse
				}

				// Remove the worry level from this monkey's list, and
				// add it to the destination monkey's list
				m.items = m.items[1:] //m.items[:len(m.items)-1]
				monkeys[dest].items = append(monkeys[dest].items, wl)
			}
		}
	}

	fmt.Println("After rounds:")

	// Get the two highest inspections, answer is the product
	ii := []int{}
	for _, m := range monkeys {
		fmt.Println(m)
		ii = append(ii, m.inspections)
	}
	sort.Ints(ii)
	fmt.Println("Final wls (s/b 1938, 47830, 52013, 52166):", ii)
	fmt.Println("Answer:", ii[len(ii)-1]*ii[len(ii)-2])

}

// Apply an operation to a number: old*old, old+n, old*n
func applyOperation(wl int64, op []string) int64 {
	assert(op[0] == "old" && (op[1] == "+" || op[1] == "*"), "Invalid operation")
	if op[1] == "+" {
		return wl + atoi64(op[2])
	} else if op[2] == "old" {
		return wl * wl
	} else {
		return wl * atoi64(op[2])
	}
}

func assert(cond bool, msg string) {
	if !cond {
		panic(msg)
	}
}

// Read and parse "monkeys" from input file
func readMonkeys(filename string) []Monkey {

	// Create empty list of monkeys, and initial monkey
	monkeys := []Monkey{}
	m := Monkey{id: 0}

	// Process each line of input file
	lines := readLines(filename)
	for _, l := range lines {

		// Blank line starts a new "monkey"
		l = strings.TrimSpace(l)
		if len(l) == 0 {
			monkeys = append(monkeys, m)
			m = Monkey{id: len(monkeys)} // add current monkey to list
			continue
		}

		// Otherwise fill in fields about the current monkey
		words := strings.Split(l, " ")
		if words[0] == "Starting" { // list of numbers
			for i := 2; i < len(words); i++ {
				n := words[i]
				if n[len(n)-1] == ',' {
					n = n[:len(n)-1]
				}
				m.items = append(m.items, atoi64(n))
			}
		} else if words[0] == "Operation:" {
			m.operation = words[3:]
		} else if words[0] == "Test:" { // e.g., "divisible by 13"
			m.test = atoi64(words[3])
		} else if words[1] == "true:" { // e.g., "throw to monkey 3"
			m.ifTrue = atoi(words[5])
		} else if words[1] == "false:" {
			m.ifFalse = atoi(words[5])
		} else if !strings.HasPrefix(l, "Monkey ") {
			panic("Invalid line: " + l)
		}
	}

	// Add the last monkey and return list
	monkeys = append(monkeys, m)
	return monkeys
}
