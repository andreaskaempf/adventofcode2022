// Advent of Code 2022, Day 17
//
// Simulate simple geometric shapes falling down a shaft, getting moved left
// and right by gusts of "gas", and falling on top of each other. For part 1,
// determine the total height of the shapes after 2022 have fallen. For Part 2,
// do the same for 1 000 000 000 000 shapes (infeasible to simulate, need to
// find a better way, perhaps by finding when the pattern starts to repeat).
//
// AK, 17 Dec 2022

package main

import (
	"fmt"
	"io/ioutil"
	//"time"
)

// A position in 2-d space
type Point struct {
	x, y int64
}

// Rocks in the chamber (sparse matrix)
var chamber map[Point]byte

// For memozation of row values
var rows map[int64]int

var pow2 []int

func main() {

	// Read the input file, just one line!
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	patt := data //string(data) // remove trailing newlines with editor
	fmt.Println("Pattern length is", len(patt))

	// Five rock shapes, expressed as pseudo-matrices
	minus := [][]int{[]int{1, 1, 1, 1}}
	plus := [][]int{[]int{0, 1, 0}, []int{1, 1, 1}, []int{0, 1, 0}}
	L := [][]int{[]int{0, 0, 1}, []int{0, 0, 1}, []int{1, 1, 1}}
	I := [][]int{[]int{1}, []int{1}, []int{1}, []int{1}}
	square := [][]int{[]int{1, 1}, []int{1, 1}}
	shapes := [][][]int{minus, plus, L, I, square}

	// Simulate the falling of rocks to the bottom of a chamber 7-wide
	chamber = map[Point]byte{} // initialize map used as sparse matrix
	rows = map[int64]int{}     // for part 2, memoization of each row
	nextShape := 0             // type of the next rock
	var height int64 = 0       // current height of the hightest rock
	pi := 0                    // start in position 0 of pattern
	var rocks int64 = 2022     // 2022 for part 1
	//var part2rocks int64 = 1000000000000 // for part 2
	rocks = 20000  // uncomment for part 2
	var rock int64 // the current rock
	pow2 = []int{1, 2, 4, 6, 8, 16, 32, 64, 128}
	var prevHeight int64
	deltas := []int64{}                   // height added by each rock during simulation
	for rock = 1; rock <= rocks; rock++ { // number of rocks

		// Get the shape of this rock
		shape := shapes[nextShape]
		nextShape++
		if nextShape >= len(shapes) {
			nextShape = 0
		}

		// Simulate appearance of the new rock: each rock appears so that its
		// left edge is two units away from the left wall and its *bottom* edge
		// is three units above the highest rock in the room (or the floor, if
		// there isn't one)
		var x int64 = 3
		var y int64 = height + int64(3+len(shape))

		// Simulate movement/falling of rock
		for {

			// Get the direction of next gas burst
			gas := patt[pi]
			pi++
			if pi >= len(patt) {
				pi = 0
			}

			// Move left/right according to gas burst, if possible
			assert(gas == '<' || gas == '>', "Bad pattern symbol!")
			if gas == '<' {
				if x > 1 && !occupied(x-1, y, shape) {
					x--
				}
			} else if gas == '>' {
				if x+int64(len(shape[0]))-1 < 7 && !occupied(x+1, y, shape) {
					x++
				}
			}

			// Fall if possible, stop this rock if not
			y-- // adjust y down
			if y-int64(len(shape)) < 0 || occupied(x, y, shape) {
				y++ // move back up
				placeShape(x, y, shape)
				if y > height {
					height = y
				}
				//fmt.Printf("Rock %d landed at %d,%d, height = %d\n", rock+1, x, y, hight)
				break
			}
		}

		// Part 1 is the answer at 2022 rocks, but continue simulation for part 2
		if rock == 2022 {
			fmt.Println("Part 1 (s/b 3068, 3114):", height)
			//break
		}

		// Show difference in height from previous iteration
		deltas = append(deltas, height-prevHeight)
		//fmt.Println(height - prevHeight)
		prevHeight = height

		// Has a pattern been found?
		/*repeatPattLength := patternFound()
		if repeatPattLength > 0 {
			fmt.Printf("Pattern found at rock %d, height = %d\n", rock, height)
			x := part2rocks / repeatPattLength // number of patterns
			modulo := part2rocks % repeatPattLength
			fmt.Println("Part 2 est:", x*height+modulo)
			break
		}*/
	}

	// Part 2: find repeating patterns, use to reduce 100B iterations to manageable size
	//repeating()

	fmt.Println("Part 2 (s/b 1514285714288):", part2(deltas))
}

func part2(deltas []int64) int64 {
	fmt.Println(len(deltas), "deltas")

	// Find repeating pattern in deltas
	l := 1740                                   // length of the pattern to look for
	for from := 0; from < len(deltas); from++ { // pattern starts here
		seq := deltas[from : from+l] // sequence to look for
		prevFind := 0
		finds := 0
		gaps := []int{}
		firstFind := 0
		for i := 0; i < len(deltas)-l; i++ {
			this := deltas[i : i+l]
			if same(seq, this) {
				if firstFind == 0 {
					firstFind = i
				}
				if prevFind > 0 {
					gap := i - prevFind
					gaps = append(gaps, gap)
				}
				prevFind = i
				finds++
			}
		}
		if finds > 2 {
			fmt.Println("Sequence found", finds, "times:")
			fmt.Println("  starts at", from, ", first found at", firstFind)
			fmt.Println("  gaps =", gaps)
			fmt.Printf("  sequence is %d long, sum = %d\n", len(seq), sum(seq))
			fmt.Printf("  prefix is %d long, sum = %d\n", firstFind, sum(deltas[:firstFind]))
			fmt.Println("Sequence:", seq)
		}
	}
	return 0
}

// Check if position is occupied by a rock of given shape,
// i.e., if any point touches
func occupied(x, y int64, shape [][]int) bool {
	for sy := 0; sy < len(shape); sy++ {
		for sx := 0; sx < len(shape[0]); sx++ {
			if shape[sy][sx] == 1 && filled(x+int64(sx), y-int64(sy)) {
				return true
			}
		}
	}
	return false
}

// Place shape at position
func placeShape(x, y int64, shape [][]int) {
	for sy := 0; sy < len(shape); sy++ {
		for sx := 0; sx < len(shape[0]); sx++ {
			if shape[sy][sx] == 1 {
				p := Point{x + int64(sx), y - int64(sy)}
				chamber[p] = 1
			}
		}
	}
}

// Check if position is occupied
func filled(x, y int64) bool {
	_, ok := chamber[Point{int64(x), int64(y)}]
	return ok
}

// Look for repeating sequences of rows in chamber, i.e., from row 1 upward,
// is there a sequence that repeats?
func repeating2() {
	maxY := maxRow() // the highest row number so far
	fmt.Println("Looking for repeating pattern, maximum row =", maxY)
	var y, z, i int64
	for z = 1; z <= maxY; z++ { // from this row upward, do they match rows 1..?
		for y = 1; y <= maxY; y++ { // from this row upward, do they match rows 1..?
			match := 0                    // number of consecutive rows that match
			for i = 0; y+i <= maxY; i++ { // check the next n rows, count matches
				if row(y+i) == row(z+i) {
					match++
				} else {
					break
				}
			}
			if match > 3 {
				fmt.Println("Rows from", y, "match first", match, "rows from row", z)
			}
		}
	}
}

// Check if a pattern has been found from the last row, i.e., do n rows from
// the last row matcht the n rows before that? Returns length of the pattern.
func patternFound() int64 {
	maxY := maxRow() // the highest row number so far
	var p, i int64
	for p = 10; p < 200; p++ { // plausible lengths of patterns
		matches := 0 // number of rows that match
		for i = 0; i < p; i++ {
			if row(maxY-i) == row(maxY-p-i) {
				matches++
			} else {
				break
			}
		}
		if matches > 10 {
			fmt.Println(matches, "rows match pattern length", p)
			return p
		}
	}
	return 0
}

// Maximum row occupied in the chamber
func maxRow() int64 {
	var rows int64 = 0
	for p, _ := range chamber {
		if p.y > rows {
			rows = p.y
		}
	}
	return rows
}

// Return value of one row in the cavern, encoded using binary to
// enable fast comparison (each row of 7 is just 1/0 values)
func row(y int64) int {

	// Return from cache if already computed
	r, ok := rows[y]
	if ok {
		return r
	}

	// Otherwise assemble from data
	r = 0
	for x := 1; x <= 7; x++ {
		if filled(int64(x), y) {
			r += pow2[x-1]
		}
	}

	// Save in cache and return
	rows[y] = r
	return r
}

// Visualize the chamber (used for Part 1 debugging)
func visualize() {
	rows := maxRow()
	var x, y int64
	for y = rows; y >= 0; y-- {
		fmt.Printf("%4d ", y)
		for x = 1; x <= 7; x++ {
			fmt.Print(ifElse(filled(x, y), "X", "."))
		}
		fmt.Print("\n")
	}
}
