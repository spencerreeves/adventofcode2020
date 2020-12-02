package day2

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

type password struct {
	lowerBound int64
	upperBound int64
	required string
	value string
	isValid bool
}


/// Expects input in the format [lower]-[upper] [required]: [password]
func NewPassword(input string) (*password, error) {
	pw := new(password)
	segments := strings.Split(input, " ")

	// Get the bounds for the required character
	bounds := strings.Split(segments[0], "-")
	lower, lowerErr := strconv.ParseInt(bounds[0], 10, 64)
	upper, upperErr := strconv.ParseInt(bounds[1], 10, 64)
	if lowerErr != nil || upperErr != nil {
		log.Printf("Bounds failed to be parsed. Lower error: %s. Upper error: %s", lowerErr, upperErr)
		return nil, errors.New("invalid bound")
	}
	pw.lowerBound = lower
	pw.upperBound = upper

	// Get the required character
	if last := len(segments[1]) - 1; last >= 0 && segments[1][last] == ':' {
		pw.required = segments[1][:last]
	} else {
		return nil, errors.New("invalid rule")
	}

	// Get the password
	pw.value = segments[2]

	// Determine if it is a valid password
	count := int64(0)
	for _, c := range pw.value {
		if string(c) == pw.required {
			count += 1
		}

		// Exit if we have surpassed the upper bound
		if count > pw.upperBound {
			break
		}
	}

	pw.isValid = count >= pw.lowerBound && count <= pw.upperBound

	return pw, nil
}

func fetchPasswords(fileName string) (passwords []*password, valid int, err error) {
	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		return passwords, valid, err
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
		pw, err := NewPassword(scanner.Text())
		if err != nil {
			return passwords, valid, err
		}

		passwords = append(passwords, pw)
		if pw.isValid {
			valid += 1
		}
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		return passwords, valid, err
	}

	return passwords, valid, err
}

// Find the two entries that sum to 2020
func Problem1() string {
	// Create, format, and count how many valid passwords there are
	_, valid, err := fetchPasswords("day2/input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	return fmt.Sprintf("The number of valid passwords is %v.", valid)
}