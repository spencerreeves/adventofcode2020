package day3

import (
	"bufio"
	"fmt"
	"os"
)

/// Expects input in the format [lower]-[upper] [required]: [password]
func readTrees(fileName string, ) (numOfTrees int, err error) {
	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		return numOfTrees, err
	}

	// Make sure to close at the end of the function
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Create buffered reader and read line by line
	xPos := 0
	yPos := 0
	var trees []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		trees = append(trees, scanner.Text())
		if spot := xPos % len(trees[yPos]); trees[yPos][spot] == '#' {
			numOfTrees += 1
		}

		xPos += 3
		yPos += 1
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		return numOfTrees, err
	}

	return numOfTrees, err
}

// Get the number of trees in the path
func Problem1() string {
	// Create, format, and count how many valid passwords there are
	numOfTrees, err := readTrees("day3/input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	return fmt.Sprintf("The number of trees in the path is %v.", numOfTrees)
}
