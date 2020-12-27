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
	data := utils.ReadTextFile("./data.txt")

	fmt.Println(solveWithSimpleMath(data))
}

func solveWithSimpleMath(input []string) int {
	result := 0

	for _, row := range input {
		result += solveExpression(row)
	}

	return result
}

func solveExpression(input string) int {
	if hasParentheses(input) {
		subs := getParenthesesSubstr(input)

		result := strconv.FormatInt(int64(executeOperation(splitBySpace(subs[1]))), 10)

		return solveExpression(strings.Replace(input, subs[0], result, 1))
	}
	return executeOperation(splitBySpace(input))
}

func executeOperation(input []string) int {
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

	return 0
}

var parenthesesRgx = regexp.MustCompile(`\(([^()][^\()]+)\)`)

func getParenthesesSubstr(input string) []string {
	return parenthesesRgx.FindStringSubmatch(input)
}

func hasParentheses(input string) bool {
	return parenthesesRgx.MatchString(input)
}
