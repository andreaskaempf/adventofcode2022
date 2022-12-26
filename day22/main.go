// Advent of Code 2022, Day 22
//
// Simulate movement on a 2D map, according to a list of instructions, which
// can either be to move n steps, or to turn 90 degrees left or right. There
// are obstacles to avoid, and one wraps around to the other side when walking
// off and edge. For Part 1, the map is in 2D. For Part 3, the map gets folded
// into a cube.
//
// AK, 22 Dec 2022

package main

import (
	"fmt"
	//"strings"
)

type Point struct {
	x, y int
}

// Keep track of what is at each location
var tiles map[Point]byte

// Keep track of the minimum X and Y values for each row/col
var minX, maxX, minY, maxY map[int]int

// List of instructions, left and right encoded as -1 and -2
var instructions []int

const LEFT int = -1
const RIGHT int = -2

func main() {

	// Read the input file: map until blank line, then set of instructions
	fname := "sample.txt"
	fname = "input.txt"
	readMap(fname)

	// Do part 1
	part1()

	// Do part 2
	part2()
}

// Part 2: same as part 1, but wrap around cube instead of 2-d space
func part2() {

	// Start in the first open tile on the first row, facing right
	y := 1
	x := minX[1]
	assert(tiles[Point{x, y}] == '.', "First tile is not open!")
	dir := 90 // start facing right

	// Process instructions, simplify by adding up all moves in each direction?
	for _, inst := range instructions {

		if inst == RIGHT { // Rotate right
			dir = ifElse(dir == 270, 0, dir+90)
		} else if inst == LEFT { // Rotate left
			dir = ifElse(dir == 0, 270, dir-90)
		} else if dir == 90 { // Move right
			x1 := x
			y1 := y
			for k := 0; k < inst; k++ {
				x1++
				if tiles[Point{x1, y1}] == 0 { // no tile, wrap to beginning
					x1 = minX[y1]
				}
				if tiles[Point{x1, y1}] == '.' {
					x = x1
				} else {
					break
				}
			}
		}
	}

}

// Part 1: follow instructions, moving around 2-d space, wrapping as necessary
func part1() {

	// Start in the first open tile on the first row, facing right
	y := 1
	x := minX[1]
	assert(tiles[Point{x, y}] == '.', "First tile is not open!")
	dir := 90 // start facing right

	// Process instructions, simplify by adding up all moves in each direction?
	for i := 0; i < len(instructions); i++ {

		// Process a turn

		// TODO: STOP IF YOU HIT A WALL, EVEN IN MID INSTRUCTION
		inst := instructions[i]
		if inst == RIGHT { // Rotate right
			dir = ifElse(dir == 270, 0, dir+90)
		} else if inst == LEFT { // Rotate left
			dir = ifElse(dir == 0, 270, dir-90)
		} else if dir == 90 { // Move right
			x1 := x
			y1 := y
			for k := 0; k < inst; k++ {
				x1++
				if tiles[Point{x1, y1}] == 0 { // no tile, wrap to beginning
					x1 = minX[y1]
				}
				if tiles[Point{x1, y1}] == '.' {
					x = x1
				} else {
					break
				}
			}

		} else if dir == 270 { // Move left
			x1 := x
			y1 := y
			for k := 0; k < inst; k++ {
				x1--
				if tiles[Point{x1, y1}] == 0 { // no tile, wrap to end
					x1 = maxX[y1]
				}
				if tiles[Point{x1, y1}] == '.' {
					x = x1
				} else {
					break
				}
			}

		} else if dir == 0 { // Move up
			x1 := x
			y1 := y
			for k := 0; k < inst; k++ {
				y1--
				if tiles[Point{x1, y1}] == 0 { // no tile, wrap to bottom
					y1 = maxY[x1]
				}
				if tiles[Point{x1, y1}] == '.' {
					y = y1
				} else {
					break
				}
			}

		} else if dir == 180 { // Move down
			x1 := x
			y1 := y
			for k := 0; k < inst; k++ {
				y1++
				if tiles[Point{x1, y1}] == 0 { // no tile, wrap to top
					y1 = minY[x1]
				}
				if tiles[Point{x1, y1}] == '.' {
					y = y1
				} else {
					break
				}
			}
		}

	}

	//fmt.Printf("Final x = %d, y = %d, dir = %d (s/b 8, 6, dir 90)\n", x, y, dir)
	facing := map[int]int{0: 3, 90: 0, 180: 1, 270: 2}
	score := 1000*y + 4*x + facing[dir]
	fmt.Println("Part 1: score (s/b 6032, 36518) =", score)
}

// Read the input file: map until blank line, then set of instructions,
// also set min/max X/Y
func readMap(fname string) {

	// Read the input file: map until blank line, then set of instructions
	readingMap := true
	tiles = map[Point]byte{}
	var x, y int
	instructions = []int{}
	for _, l := range readLines(fname) {
		if len(l) == 0 {
			readingMap = false
		} else if readingMap {
			for x = 0; x < len(l); x++ {
				if l[x] != ' ' {
					tiles[Point{x + 1, y + 1}] = l[x]
				}
			}
			y++
		} else {
			instructions = parseInstructions(l)
		}
	}

	// Get the min/max col in each row and min/max row in each column
	minX = map[int]int{}
	maxX = map[int]int{}
	minY = map[int]int{}
	maxY = map[int]int{}
	for p, _ := range tiles {
		if minX[p.y] == 0 || p.x < minX[p.y] {
			minX[p.y] = p.x
		}
		if p.x > maxX[p.y] {
			maxX[p.y] = p.x
		}
		if minY[p.x] == 0 || p.y < minY[p.x] {
			minY[p.x] = p.y
		}
		if p.y > maxY[p.x] {
			maxY[p.x] = p.y
		}
	}
}

// Parse string of instructions: list of numbers,
// where -1 means left and -2 means right
func parseInstructions(s string) []int {
	result := []int{}
	n := 0 // the current number
	for i := 0; i < len(s); i++ {
		c := s[i]
		if c >= '0' && c <= '9' {
			n = n*10 + int(c-'0')
		} else if c == 'L' || c == 'R' {
			if n > 0 {
				result = append(result, n)
				n = 0
			}
			result = append(result, ifElse(c == 'L', -1, -2))
		} else {
			panic("Invalid instruction!")
		}
	}
	if n > 0 {
		result = append(result, n)
	}
	return result
}
