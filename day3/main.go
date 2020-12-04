package day3

import (
	"bufio"
	"fmt"
	"os"
)

type Path struct {
	xOffset int
	yOffset int
	x int
	y int
	numOfTrees int
}

/// Expects input in the format [lower]-[upper] [required]: [password]
func readTrees(fileName string, paths []Path) (err error) {
	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		return err
	}

	// Make sure to close at the end of the function
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Create buffered reader and read line by line
	yPos := 0
	var trees []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Add this line to the list of trees
		trees = append(trees, scanner.Text())

		for i := range paths {
			// When this iteration produces a valid location on the slope. Useful for slopes with a y offset greater than 1
			if yPos % paths[i].yOffset == 0 {
				if relativeX := paths[i].x % len(trees[yPos]); trees[yPos][relativeX] == '#' {
					paths[i].numOfTrees += 1
				}

				paths[i].x += paths[i].xOffset
				paths[i].y += paths[i].yOffset
			}
		}

		yPos += 1
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		return err
	}

	return err
}

// Get the number of trees in the path
func Problem1() string {
	paths := []Path{Path{3,1,0,0,0}}
	err := readTrees("day3/input.txt", paths)
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	return fmt.Sprintf("The number of trees in the path is %v.", paths[0].numOfTrees)
}

// Get the number of trees in multiple paths
func Problem2() string {
	paths := []Path{
		Path{1,1, 0, 0, 0},
		Path{3,1, 0, 0, 0},
		Path{5,1, 0,0,0},
		Path{7,1,0,0,0},
		Path{1,2,0,0,0}}

	err := readTrees("day3/input.txt", paths)
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	multiplied := 1
	var trees []int
	for i := range paths{
		multiplied = multiplied * paths[i].numOfTrees
		trees = append(trees, paths[i].numOfTrees)
	}

	return fmt.Sprintf("The number of trees in each path is %v. Multiplied together: %v", trees, multiplied)
}
