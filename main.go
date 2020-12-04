package main

import (
	"./day1"
	"./day2"
	"./day3"
	"fmt"
	"log"
)

func main() {
	log.Printf("\n\n*****     Advent of Code     *****     \n")
	fmt.Printf("**  Day 1 **\n")
	trackFunc("Day 1, Problem 1", day1.Problem1)
	trackFunc("Day 1, Problem 2", day1.Problem2)

	fmt.Printf("**  Day 2  **\n")
	trackFunc("Day 2, Problem 1", day2.Problem1)
	trackFunc("Day 2, Problem 2", day2.Problem2)

	fmt.Printf("**  Day 3 **\n")
	trackFunc("Day 3, Problem 1", day3.Problem1)
}