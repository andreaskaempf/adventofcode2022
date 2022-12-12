# Advent of Code 2022

My solutions for the Advent of Code 2022, 
see https://adventofcode.com/2022

* **Day 1** (Go): Find the maximum sum of newline-separated 
  groups of numbers. For part 2, find the sum of the top 3 (*easy*).

* **Day 2** (Go): Rock/Paper/Scissors simulation  (*easy*)

* **Day 3** (Go, Python): Sum up characters that are common in two halves of
  some strings (Part 1), and common to groups of 3 strings (Part 2) (*easy*)

* **Day 4** (Go): Given pairs of numeric ranges, in how many pairs is one range
  entirely contained within another (Part 1), and how many pairs overlap at 
  all (Part 2) (*easy*)

* **Day 5** (Go): Simulate execution of instructions to move crates from 
  one tower to another, one at a time (Part 1), or in groups (Part 2) (*easy*)

* **Day 6** (Go): Look for first block of 4 (Part 1) or 14 (Part 2) 
  non-repeating characters in a string.  (*easy*)

* **Day 7** (Go): Given a list of Unix shell commands and output (just 
  ls and cd), parse these and find the sum of space (calculated recursively,
  i.e., including subdirectories) used by all directories <= 100k (Part 1), 
  and the size of the smallest directory that would free 
  up a required amount of space. (*medium*)

* **Day 8** (Go): Given a topographical map (matrix) of tree heights, count 
  the number of trees that have visibility all the way to the edge (Part 1), 
  and the highest "visibility" score of any tree, where that score is the 
  product of the numbers of trees less than the current tree in each 
  direction. (*medium*)

* **Day 9** (Go): Simulate movement of "knots" along a rope, in response to 
  the first knot being moved. For part 1, there are only two knots (head and 
  tail), for part 2 there are 10. After the simulation, report the number 
  of positions the tail has covered. (*easy*)

* **Day 10** (Go): Simulate accumulator register during a series of given 
  ADD or NOOP instructions, and report the accumulator values during 
  selected clock cycles (Part 1). For part 2, simulate drawing of pixels 
  on a screen, using the sequence of acculator values, and report the eight 
  letters that appear. (*medium*)

* **Day 11** (Go): Simulate transfer of objects between a bunch of monkeys, 
  with "worry levels" assigned to each object. Each monkey modifies the 
  worry level according to some rules, then passes it to one of two monkeys, 
  depending on whether the worry level is divisible by that monkey's "test" 
  number.  Count up the number of inspections each monkey makes during the 
  simulation. The answer is the product of the two highest inspection counts. 
  Trivial (if tedious) for 20 iterations in Part 1, but integer values 
  overflow for 10,000 iterations in Part 2, unless you apply an adjustment 
  that preserves the decision outcomes while keeping the numbers fom getting 
  too large. (*hard*)

* **Day 12** (Go): Find the lowest cost path (i.e., shortest number of steps)
  through a terrain of letters, from point S to E, allowing 'increase' (e.g.,
  next letter) of maximum 1. For Part 2, find the shortest path from any 'a'
  cell to 'E'. (*medium*, using yourbasic/graph library)

To compile and run a **Go** program
* Change into the directory with the program
* go build day01.go
* ./day01

To run a **Julia** program
* Change into the directory with the program
* julia day02.jl

To run a **Python** program
* Change into the directory with the program
* python day06.py

AK, Dec 2022
