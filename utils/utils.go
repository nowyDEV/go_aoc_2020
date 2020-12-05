package utils

import (
	"bufio"
	"log"
	"os"
	"encoding/hex"
	"strconv"
	"regexp"
	"strings"
)

// ReadTextFile is a helper function for reading txt files
func ReadTextFile(filePath string) []string {
	file, err := os.Open(filePath)

	if err != nil {
		log.Fatalf("failed to open")
	}

	scanner := bufio.NewScanner(file)

	scanner.Split(bufio.ScanLines)
	var listOfLines []string

	for scanner.Scan() {
		listOfLines = append(listOfLines, scanner.Text())
	}

	file.Close()

	return listOfLines
}

// IsInRange checks whether input number is between start and end (including them)
func IsInRange(input int, start int, end int) bool {
	return input >= start && input <= end
}

// IsHexValue checks whether input string is 7-char hex value (e.g. #ffffff)
func IsHexValue(input string) bool {
	_, err := hex.DecodeString(strings.Replace(input, "#", "", 1))
	if (err != nil) {
		return false
	}
	return len(input) == 7
}

// GetIntFromString extracts numbers from string and converts them to integer
func GetIntFromString(input string) int {
	r, _ := regexp.Compile("([0-9]+)")
	value, err := strconv.Atoi(r.FindString(input))

	if (err != nil) {
		log.Fatalf("failed to convert")
	}

	return value
}

// ContainsString checks whether provided array contains provided string
func ContainsString(s []string, e string) bool {
	for _, a := range s {
			if a == e {
					return true
			}
	}
	return false
}