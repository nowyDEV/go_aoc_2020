package main

import (
	"aoc-2020/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := utils.ReadTextFile("./data_test.txt")

	result := 0
	for _, row := range data {
		result += solveExpression(row)
	}

	fmt.Println("final result", result)
}

func solveExpression(input string) int {
	fmt.Println("solveExpression", input)

	if hasParentheses(input) {
		subs := getParenthesesSubstr(input)

		fmt.Println("substrings", subs)

		result := strconv.FormatInt(int64(executeOperation(splitBySpace(subs[1]))), 10)

		fmt.Println("result", result)

		return solveExpression(strings.Replace(input, subs[0], result, 1))
	}
	return executeOperation(splitBySpace(input))
}

func executeOperation(input []string) int {
	fmt.Println("executeOperation", input)

	if len(input) == 3 {
		return calculate(input)
	}

	partialResult := calculate(input[0:3])
	newInput := append([]string{intToStr(partialResult)}, input[3:]...)

	return executeOperation(newInput)
}

func intToStr(num int) string {
	return strconv.FormatInt(int64(num), 10)
}

func splitBySpace(str string) []string {
	return strings.Split(str, " ")
}

func calculate(input []string) int {
	if len(input) != 3 {
		log.Fatalf("Unable to calculate, wrong input")
	}

	sign := input[1]

	if sign == "+" {
		return utils.GetIntFromString(input[0]) + utils.GetIntFromString(input[2])
	}
	if sign == "*" {
		return utils.GetIntFromString(input[0]) * utils.GetIntFromString(input[2])
	}

	fmt.Println("returning 0", input)
	return 0
}

var singleOperationParentheses = regexp.MustCompile(`\(([0-9].[^\(]*)\)`)
var parenthesesRgx = regexp.MustCompile(`\((.*?)\)`)

func getParenthesesSubstr(input string) []string {
	return singleOperationParentheses.FindStringSubmatch(input)
}

func hasParentheses(input string) bool {
	return parenthesesRgx.MatchString(input)
}
