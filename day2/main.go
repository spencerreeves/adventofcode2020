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
	rule1 int64
	rule2 int64
	required string
	value string
	isValid bool
}

// Checks to see if this is a valid password based on the Sled rental corporate password policy
// This policy states that the required character must occur more than or equal to rule1, but less than or equal to rule2
func sledPasswordPolicy(pw *password) (bool, error) {
	// Determine if it is a valid password
	count := int64(0)
	for _, c := range pw.value {
		if string(c) == pw.required {
			count += 1
		}

		// Exit if we have surpassed the upper bound
		if count > pw.rule2 {
			break
		}
	}

	return count >= pw.rule1 && count <= pw.rule2, nil
}

// Checks to see if this is a valid password based on the Tobbaggan rental corporate password policy
// This policy states that the required character must occur at exactly one position - rule1 or rule2
func tobogganPasswordPolicy(pw *password) (bool, error) {
	// Make sure rule2 is within bounds of the password
	if pw.rule2 > int64(len(pw.value)) {
		return false, nil
	}

	// Check if required is at position of rule 1 or rule 2
	rule1, rule2 := pw.rule1 - 1, pw.rule2 - 1
	i1, i2 := string(pw.value[rule1]), string(pw.value[rule2])
	return (i1 == pw.required || i2 == pw.required) && i1 != i2, nil
}


/// Expects input in the format [lower]-[upper] [required]: [password]
func NewPassword(input string, validator func(*password)(bool, error)) (pw *password, err error) {
	pw = new(password)
	segments := strings.Split(input, " ")

	// Get the bounds for the required character
	rules := strings.Split(segments[0], "-")
	rule1, rule1Err := strconv.ParseInt(rules[0], 10, 64)
	rule2, rule2Err := strconv.ParseInt(rules[1], 10, 64)
	if rule1Err != nil || rule2Err != nil {
		log.Printf("Invalid rules. Rule 1 error: %s. Rule 2 error: %s\n", rule1Err, rule2Err)
		return nil, errors.New("invalid rules")
	}
	pw.rule1 = rule1
	pw.rule2 = rule2

	// Get the required character
	if last := len(segments[1]) - 1; last >= 0 && segments[1][last] == ':' {
		pw.required = segments[1][:last]
	} else {
		return nil, errors.New("invalid rule")
	}

	// Get the password
	pw.value = segments[2]

	pw.isValid, err = validator(pw)
	if err != nil {
		return nil, errors.New("failed validation")
	}

	return pw, nil
}

func fetchPasswords(fileName string, validator func(*password)(bool, error)) (passwords []*password, valid int, err error) {
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
		pw, err := NewPassword(scanner.Text(), validator)
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

// Find the number of passwords that satisfy the sled password policy
func Problem1() string {
	// Create, format, and count how many valid passwords there are
	_, valid, err := fetchPasswords("day2/input.txt", sledPasswordPolicy)
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	return fmt.Sprintf("The number of valid passwords is %v.", valid)
}

// Find the number of passwords that satisfy the tobboggan password policy
func Problem2() string {
	// Create, format, and count how many valid passwords there are
	_, valid, err := fetchPasswords("day2/input.txt", tobogganPasswordPolicy)
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	return fmt.Sprintf("The number of valid passwords is %v.", valid)
}