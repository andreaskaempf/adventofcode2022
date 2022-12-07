// Advent of Code 2022, Day 07
//
// Given a list of Unix shell commands and output (just ls and cd), parse these
// and find the sum of space (calculated recursively, i.e., including
// subdirectories) used by all directories <= 100k (Part 1), and the size of
// the smallest directory that would free up a required amount of space.
//
// AK, 7 Dec 2022

package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

// Directory information is stored in a structure. There is
// initially one of these for the root directory, and it contains
// lists of files and subdirectories within.
type Directory struct {
	name    string      // name of the directory, just the last part
	parent  *Directory  // pointer to parent directory
	files   []File      // list of files
	subdirs []Directory // list of subdirectories
}

// Information about one file
type File struct {
	name string
	size int
}

// Global variables for Parts 1 and 2
var Part1Total, Part2Size int

func main() {

	// Read the input file
	//filename := "sample.txt"
	filename := "input.txt" // uncomment appropriate file name
	blob, _ := ioutil.ReadFile(filename)
	lines := strings.Split(string(blob), "\n")
	fmt.Println(len(lines), "lines read")

	// Process each line: change directories, and build up recursive lists of
	// files and subdirectories
	root := Directory{name: "/"} // Create one root directory to start
	var curdir *Directory        // Pointer to the current directory, set by "cd" command
	for _, l := range lines {    // Go through each line of input

		// Commands start with $: cd to change dir (ignore ls)
		words := strings.Split(l, " ")
		if words[0] == "$" {
			if words[1] == "cd" {
				if words[2] == "/" { // change to root
					curdir = &root
				} else if words[2] == ".." { // up one dir
					curdir = curdir.parent
				} else { // change into directory below, create if necessary
					found := false
					for i := 0; i < len(curdir.subdirs); i++ {
						if curdir.subdirs[i].name == words[2] {
							curdir = &curdir.subdirs[i]
							found = true
							break
						}
					}
					if !found { // New directory does not exist, create it
						newDir := Directory{name: words[2], parent: curdir}
						curdir.subdirs = append(curdir.subdirs, newDir)
						curdir = &curdir.subdirs[len(curdir.subdirs)-1] // can't use &newDir
					}
				}
			}

			// Otherwise add file with its size to current directory (ignore
			// subdirectory name output)
		} else if words[0] != "dir" { // ignore subdir name output
			f := File{name: words[1], size: atoi(words[0])} // create a File object
			curdir.files = append(curdir.files, f)          // add it to list for current subdir
		}
	}

	// Part 1: get the size of each directory, report sum of those <= 100k
	used := root.totSize() // calculates sizes of all directories, incl Part 1
	fmt.Println("Part 1:", Part1Total)

	// Part 2: given capacity of 70000000 and used space, find the smallest
	// directory that is big enough to free up the required 30000000 space
	unused := 70000000 - used   // current free space
	freeUp := 30000000 - unused // amount we need to free up
	part2(&root, freeUp)        // finds smallest directory that exceeds freeUp
	fmt.Println("Part 2:", Part2Size)
}

// For Part 2, walk through directories, find smallest directory that
// has size >= freeUp
func part2(d *Directory, freeUp int) {

	// Check this directory: does it exceed the space required, and is it
	// smaller than any solution found so far? Note we don't need to check
	// to exclude the root directory, since its size will be too big anyway.
	if d.totSize() >= freeUp && (Part2Size == 0 || d.totSize() < Part2Size) {
		Part2Size = d.totSize()
	}

	// Check all subdirectories
	for _, subdir := range d.subdirs {
		part2(&subdir, freeUp)
	}
}

// Get size of a directory, including subdirs
func (d *Directory) totSize() int {

	// Sum up files in this directory
	tot := 0
	for _, f := range d.files {
		tot += f.size
	}

	// Sum up subdirectories
	for _, subdir := range d.subdirs {
		tot += subdir.totSize()
	}

	// For part 1, add up sizes <= 100k
	if tot <= 100000 {
		Part1Total += tot
	}

	// Return size of this directory, including its subdirs
	return tot
}

// Parse an integer, show message and return -1 if error
func atoi(s string) int {
	n, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		fmt.Println("Could not parse integer:", s)
		n = -1
	}
	return int(n)
}
