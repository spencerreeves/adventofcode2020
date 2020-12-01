package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readAndStoreAsMap(fileName string) (mp map[int64]string, err error){
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

func find2020Pair(mp map[int64]string) (v1 int64, v2 int64, err error) {
	for key, _ := range mp {
		if _, exists := mp[2020 - key]; exists {
			return key, 2020 - key, nil
		}
	}

	return 0, 0, errors.New("no matching pair")
}

func main() {
	log.Println("Find the two entries that sum to 2020")

	// Create hash map from entry
	mp, err := readAndStoreAsMap("inputs/day_1_problem_1.txt");
	if err != nil {
		log.Printf(fmt.Sprintf("Error reading file: %v", err))
		panic(err)
	}

	// Find the 2020 pair
	v1, v2, err := find2020Pair(mp)
	if err != nil {
		log.Printf(fmt.Sprintf("Error finding 2020 pair: %v", err))
		panic(err)
	}
	log.Printf(fmt.Sprintf("Value 1: %v | Value 2: %v | Multiplied: %v", v1, v2, v1 * v2))
}

