package day4

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var /* const */ hclReg = regexp.MustCompile("^#[\\da-f]{6}$")
var /* const */ eclReg = regexp.MustCompile("^(amb|blu|brn|gry|grn|hzl|oth)$")
var /* const */ pidReg = regexp.MustCompile("^\\d{9}$")

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

func alwaysValid(key string, value string) bool {
	return true
}

func withinRange(value string, lower int64, upper int64) bool {
	i, err := strconv.ParseInt(value, 10, 64)
	return err == nil && i <= upper && i >= lower
}

func isValidPassportField(key string, value string) bool {
	switch key {
	case "byr": return withinRange(value, 1920, 2002)
	case "iyr": return withinRange(value, 2010, 2020)
	case "eyr": return withinRange(value, 2020, 2030)
	case "hgt":
		unit := value[len(value) - 2:]
		return unit == "cm" && withinRange(value[:len(value) - 2], 150, 193) ||
			unit == "in" && withinRange(value[:len(value) - 2], 59, 76)
	case "hcl": return hclReg.MatchString(value)
	case "ecl": return eclReg.MatchString(value)
	case "pid": return pidReg.MatchString(value)
	case "cid": return true
	default: return false
	}
}

func getPassports(fileName string, isValidPassport func(uint8) bool, isValidKeyPair func(string, string) bool) (passports []Passport, count int, err error) {
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
				if isValidKeyPair(key, value) {
					passport.flags = setFlag(key, passport.flags)
				}
			}
		} else {
			passports = append(passports, passport)
			if isValidPassport(passport.flags) {
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
	_, count, err := getPassports("day4/input.txt", isValidNorthPoleCredentials, alwaysValid)
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	return fmt.Sprintf("The number of valid passports, excluding Country ID, is %v.", count)
}

// Add validation logic to the expedited passport hacking
func Problem2() string {
	_, count, err := getPassports("day4/input.txt", isValidNorthPoleCredentials, isValidPassportField)
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	return fmt.Sprintf("The number of valid passports, excluding Country ID, is %v.", count)
}
