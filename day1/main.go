package day1

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
)

func readAndStoreAsMap(fileName string) (mp map[int64]string, err error) {
	mp = make(map[int64]string)

	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		return mp, err
	}

	// Make sure to close at the end of the function
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	// Create buffered reader and read line by line
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		i, err := strconv.ParseInt(scanner.Text(), 10, 64)
		if err != nil {
			return mp, err
		}
		mp[i] = scanner.Text()
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		return mp, err
	}

	return mp, err
}

// Find the pair in the map such that when added together they sum to 2020.
func find2020Pair(mp map[int64]string) (v1 int64, v2 int64, err error) {
	for key := range mp {
		if _, exists := mp[2020-key]; exists {
			return key, 2020 - key, nil
		}
	}

	return 0, 0, errors.New("no matching pair")
}

// Find the two entries that sum to 2020
func Problem1() string {
	// Create hash map from entry
	mp, err := readAndStoreAsMap("day1/input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	// Find the 2020 pair
	v1, v2, err := find2020Pair(mp)
	if err != nil {
		return fmt.Sprintf("Error finding 2020 pair: %s", err)
	}

	return fmt.Sprintf("The pair that add up to 2020 is (%v, %v). Multiplied together they are %v", v1, v2, v1*v2)
}

// Find the pair in the map such that when added together they sum to 2020.
func find2020Triple(mp map[int64]string) (v1 int64, v2 int64, v3 int64, err error) {
	for key := range mp {
		for key2 := range mp {
			_, exists := mp[2020-key-key2]
			if key != key2 && exists  {
				return key, key2, 2020-key-key2, nil
			}
		}
	}

	return 0, 0, 0, errors.New("no matching triple")
}

// Find the three entries that sum to 2020
func Problem2() string {
	// Create hash map from entry
	mp, err := readAndStoreAsMap("day1/input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	// Find the 2020 pair
	v1, v2, v3, err := find2020Triple(mp)
	if err != nil {
		return fmt.Sprintf("Error finding 2020 triple: %s", err)
	}

	return fmt.Sprintf("The triple that add up to 2020 is (%v, %v, %v). Multiplied together they are %v", v1, v2, v3, v1*v2*v3)
}
