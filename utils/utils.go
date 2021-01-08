package utils

import (
	"bufio"
	"encoding/hex"
	"log"
	"os"
	"regexp"
	"strconv"
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
	if err != nil {
		return false
	}
	return len(input) == 7
}

// GetIntFromString extracts numbers from string and converts them to integer
func GetIntFromString(input string) int {
	r, _ := regexp.Compile("([0-9]+)")
	value, err := strconv.Atoi(r.FindString(input))

	if err != nil {
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

// ContainsInt checks whether provided array contains provided number
func ContainsInt(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

// Sum return the sum of array/slice numbers
func Sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

// FindIndex returns index of a string item inside slice
func FindIndex(list []string, item string) int {
	for index, word := range list {
		if word == item {
			return index
		}
	}
	return -1
}

// FindIndexNums returns index of a number item inside slice
func FindIndexNums(list []int, item int) int {
	for index, num := range list {
		if num == item {
			return index
		}
	}
	return -1
}

// FilterOut filters out provided string value from the list of strings
func FilterOut(list []string, item string) []string {
	var result []string
	for _, listItem := range list {
		if listItem != item {
			result = append(result, listItem)
		}
	}

	return result
}

// ReverseString returns a new reversed string
func ReverseString(input string) string {
	var result []string

	for i := len(input) - 1; i >= 0; i-- {
		result = append(result, input[i:i+1])
	}

	return strings.Join(result, "")
}

// ReverseNumbers reverses a slice of integers
func ReverseNumbers(numbers []int) []int {
	for i := 0; i < len(numbers)/2; i++ {
		j := len(numbers) - i - 1
		numbers[i], numbers[j] = numbers[j], numbers[i]
	}

	return numbers
}

// GetHighestNumber returns highest int from the slice
func GetHighestNumber(numbers []int) int {
	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > result {
			result = numbers[i]
		}
	}

	return result
}

// GetHighestNumber returns lowest int from the slice
func GetLowestNumber(numbers []int) int {
	result := numbers[0]

	for i := 1; i < len(numbers); i++ {
		if numbers[i] < result {
			result = numbers[i]
		}
	}

	return result
}

// SubtractSlice removes elements provided in the second argument list from the slice
func SubtractSlice(target []int, values []int) (result []int) {
	for _, item := range target {
		if !ContainsInt(values, item) {
			result = append(result, item)
		}
	}

	return result
}
