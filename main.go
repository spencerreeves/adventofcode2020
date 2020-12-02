package main

import (
	"./day1"
	"./day2"
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
}