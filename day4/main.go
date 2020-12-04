package day4

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Passport struct {
	credentials map[string]string
	flags uint8
}


func splitKeyValue(s string, sep string) (string, string) {
	x := strings.Split(s, sep)
	return x[0], x[1]
}

func setFlag(s string, flags uint8) uint8 {
	switch s {
	case "byr": return flags | 0x80 // 10000000
	case "iyr": return flags | 0x40 // 01000000
	case "eyr": return flags | 0x20 // 00100000
	case "hgt": return flags | 0x10 // 00010000
	case "hcl": return flags | 0x08 // 00001000
	case "ecl": return flags | 0x04 // 00000100
	case "pid": return flags | 0x02 // 00000010
	case "cid": return flags | 0x01 // 00000001
	default: return flags
	}
}

func isValidNorthPoleCredentials(flags uint8) bool {
	return flags == 0xFF || flags == 0xFE
}

func getPassports(fileName string, validator func(uint8) bool) (passports []Passport, count int, err error) {
	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		return passports, count, err
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
	passport := Passport{
		credentials: make(map[string]string),
		flags: 0x00,
	}
	for scanner.Scan() {
		keyValues := strings.Fields(scanner.Text())
		if len(keyValues) >= 1 {
			for entry := range keyValues {
				key, value := splitKeyValue(keyValues[entry], ":")
				passport.credentials[key] = value
				passport.flags = setFlag(key, passport.flags)
			}
		} else {
			passports = append(passports, passport)
			if validator(passport.flags) {
				count += 1
			}

			passport = Passport{
				credentials: make(map[string]string),
				flags: 0x00,
			}
		}
	}

	// Check for errors
	if err := scanner.Err(); err != nil {
		return passports, count, err
	}

	return passports, count, err
}

// Expedite the passport checking line and hack the system to allow me to enter
func Problem1() string {

	_, count, err := getPassports("day4/input.txt", isValidNorthPoleCredentials)
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	return fmt.Sprintf("The number of valid passports (excluding Country ID) is %v.", count)
}