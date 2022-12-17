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
	"time"
)

// A position in 2-d space
type Point struct {
	x, y int64
}

// Rocks in the chamber (sparse matrix)
var chamber map[Point]byte

func main() {

	// Read the input file, just one line!
	fname := "sample.txt"
	fname = "input.txt"
	data, _ := ioutil.ReadFile(fname)
	patt := string(data) // remove trailing newlines with editor
	fmt.Println("Pattern length is", len(patt))

	// Five rock shapes
	minus := [][]int{[]int{1, 1, 1, 1}}
	plus := [][]int{[]int{0, 1, 0}, []int{1, 1, 1}, []int{0, 1, 0}}
	L := [][]int{[]int{0, 0, 1}, []int{0, 0, 1}, []int{1, 1, 1}}
	I := [][]int{[]int{1}, []int{1}, []int{1}, []int{1}}
	square := [][]int{[]int{1, 1}, []int{1, 1}}
	shapes := [][][]int{minus, plus, L, I, square}

	// For timing
	t0 := time.Now().Unix()

	// Simulate the falling of rocks to the bottom of a chamber 7-wide
	chamber = map[Point]byte{} // initialize map used as sparse matrix
	nextShape := 0             // type of the next rock
	var height int64 = 0       // current height of the hightest rock
	pi := 0                    // start in position 0 of pattern
	var rocks int64 = 2022     // 2022 for part 1
	rocks = 1000000000000      // uncomment for part 2
	var rock int64
	for rock = 0; rock < rocks; rock++ { // number of rocks

		// Show progress
		if rock%1000000 == 0 {
			secs := float64(time.Now().Unix() - t0)
			frac := float64(rock) / float64(rocks)
			totHrs := (secs / frac) / (60 * 60)
			fmt.Printf("%.4f%% done, %.1f hours to go, height = %d\n", frac*100.0, totHrs, height)
		}

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
			//fmt.Printf("Rock %d at %d,%d\n", rock+1, x, y)
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
			//fmt.Printf("After gas burst %c: %d,%d\n", gas, x, y)

			// Adjust y down
			y--

			// Fall if possible, stop this rock if not
			if y-int64(len(shape)) < 0 || occupied(x, y, shape) {
				y++ // move back up
				placeShape(x, y, shape)
				if y > height {
					height = y
				}
				//fmt.Printf("Rock %d landed at %d,%d, height = %d\n", rock+1, x, y, hight)
				//visualize()
				break
			}
		}
		if rock == 2 {
			//break
		}
	}
	fmt.Println("Part 1 (s/b 3068): height =", height)
	//fmt.Println(chamber)
	//visualize()

	// Part 2
	//repeating()
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

// Visualize the chamber
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

// Maximum row occupied in the chamber
func maxRow() int64 { // int64?
	var rows int64 = 0
	for p, _ := range chamber {
		if p.y > rows {
			rows = p.y
		}
	}
	return rows
}

// Look for repeating rows in chamber
func repeating() {
	row1 := row(1)
	maxY := maxRow()
	fmt.Println("Row 1:", row1)
	var y int64
	for y = 2; y <= maxY; y++ {
		if same(row(y), row1) {
			//fmt.Println("Row", y, "matches row 1")
			if same(row(y+1), row(2)) {
				fmt.Println("Next row matches as well!")
			}
		}
	}
}

func row(y int64) []int {
	r := make([]int, 7, 7)
	for x := 1; x <= 7; x++ {
		if filled(int64(x), y) {
			r[x-1] = 1
		}
	}
	return r
}
