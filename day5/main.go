package day5

import (
	"bufio"
	"fmt"
	"os"
)

type PeripheralScan struct {
	passports []*Passport
	deducedPassport *Passport
	maxSeatId int
}

type Passport struct {
	raw string
	binary uint16
	seatId int
	row int
	column int
}

func NewPassport (input string) *Passport {
	var binary uint16
	for index := range input {
		binary = binary << 1 | toBit(input[index])
	}

	return &Passport{
		input,
		binary,
		int(binary >> 3) * 8 + int(binary & 0x7),
		int(binary & 0x7),
		int(binary >> 3),
	}
}

func NewPassportFromBinary(input uint16) *Passport {
	raw := ""
	for index := range []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0} {
		raw = toString(input >> index, index > 2) + raw
	}

	return NewPassport(raw)
}

func toBit(input uint8) uint16 {
	if input == 'F' || input == 'L' {
		return 0
	}

	return 1
}

func toString(bit uint16, top bool) string {
	if top {
		if bit == 0 {
			return "F"
		}
		return "B"
	}

	if bit == 1 {
		return "R"
	}

	return "L"
}

func scanNearByPassports(fileName string) (peripheralScan PeripheralScan, err error){
	// Open file
	file, err := os.Open(fileName)
	if err != nil {
		return peripheralScan, err
	}

	// Make sure to close the file at the end of the function
	defer func() {
		cerr := file.Close()
		if err == nil {
			err = cerr
		}
	}()

	maxSeatId := 0
	var mySeat uint16
	var passports []*Passport
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		passport := NewPassport(scanner.Text())
		passports = append(passports, passport)
		if passport.seatId > maxSeatId {
			maxSeatId = passport.seatId
		}
		mySeat ^= passport.binary
	}

	return PeripheralScan{
		passports: passports,
		deducedPassport: NewPassportFromBinary(mySeat),
		maxSeatId: maxSeatId,
	}, err
}

func Problem1() string {
	scan, err := scanNearByPassports("day5/input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	return fmt.Sprintf("Maximum seat id is %v.", scan.maxSeatId)
}

func Problem2() string {
	scan, err := scanNearByPassports("day5/input.txt")
	if err != nil {
		return fmt.Sprintf("Error reading file: %s", err)
	}

	// Hacky. There should be a way to do this by xor-ing the bits, but I haven't found it yet.
	available := make([]int, 1024)
	for index := range scan.passports {
		available[scan.passports[index].binary] = 1
	}

	seatId := 0
	for index := range available {
		if available[index] != 1 && index > 0 && available[index - 1] == 1 && index < len(available) -2 && available[index + 1] == 1 {
			seatId = index
		}
	}

	return fmt.Sprintf("My seat id is %v.", seatId)
}