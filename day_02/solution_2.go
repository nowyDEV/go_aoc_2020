package main

import ( 
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func main() {
	text := getFileContents()

	type PasswordConfig struct {
		positions []int
		requiredChar string
		password string
	}
	
	var passwords []PasswordConfig

	for _, item := range text {
		arr := strings.Split(item, " ")

		config := PasswordConfig{convertListOfStrToInt(strings.Split(arr[0], "-")) , strings.Replace(arr[1], ":", "", 1), arr[2]}

		passwords = append(passwords, config)
	}

	var validItems int

	for _, item := range passwords {
		positionsOfRequiredChar := getCharPositions(item.password, item.requiredChar)
		numOfValidPositions := getNumOfOccurencies(positionsOfRequiredChar, item.positions)

		if numOfValidPositions == 1 {
			validItems = validItems + 1
		}
	}

	fmt.Println(validItems)
}

func getFileContents() []string {
	file, err := os.Open("data.txt") 
  
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

func getCharPositions(text string, char string) []int {
	textArr := strings.Split(text, "")
	var positions []int

	for i, item := range textArr {
		if item == char {
			positions = append(positions, i + 1)
		}
	}

	return positions
}

func getNumOfOccurencies(input []int, list []int) int {
	var occurencies []int

	for _, item := range input {
		if item == list[0] || item == list[1] {
			occurencies = append(occurencies, item)
		}
	}

	return len(occurencies)
}

func convertListOfStrToInt(list []string) []int {
	var intList []int

	for _, item := range list {
		integer, err := strconv.Atoi(item)

		if err != nil { 
			log.Fatalf("failed to open") 
		} 
		
		intList = append(intList, integer)
	}

	return intList
}
