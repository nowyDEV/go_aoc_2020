package main

import (
	"aoc-2020/utils"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type rule struct {
	index      int
	conditions [][]string
}

func main() {
	data := utils.ReadTextFile("./data_test.txt")

	var rules []rule
	var answers []string

	for _, line := range data {
		if utils.ContainsString(strings.Split(line, ""), ":") {
			rules = append(rules, createRule(line))
		} else if len(line) > 0 {
			answers = append(answers, line)
		}
	}

	fmt.Println("rules: ", rules)
	fmt.Println("--------------------------------------------")

	fmt.Println(getPossibleAnswers(getRule(1, rules), rules))
}

func getPossibleAnswers(rule rule, rules []rule) [][]string {
	fmt.Println("getPossibleAnswers: ", rule)

	var result [][]string

	if len(rule.conditions) == 2 {
		var resultOR [][]string

		left := getAnswers(rule.conditions[0], rules)
		right := getAnswers(rule.conditions[1], rules)

		fmt.Println("left: ", left)
		fmt.Println("right: ", right)

		var mergeLeft []string
		var mergeRight []string

		for _, itemLeft := range left {
			mergeLeft = append(mergeLeft, itemLeft...)
		}

		for _, itemRight := range right {
			mergeRight = append(mergeRight, itemRight...)
		}

		fmt.Println("mergeLeft: ", mergeLeft)
		fmt.Println("mergeRight: ", mergeRight)

		fmt.Println("resultOR: ", resultOR)

		result = append(result, mergeLeft, mergeRight)

	} else {
		result = append(result, getAnswers(rule.conditions[0], rules)...)
	}

	fmt.Println("getPossibleAnswers return: ", result)

	return result
}

func getAnswers(conditions []string, rules []rule) [][]string {
	fmt.Println("getAnswers: ", conditions)

	if conditions[0] == "a" || conditions[0] == "b" {
		return [][]string{conditions}
	}

	var result [][]string

	for _, item := range conditions {
		result = append(result, getPossibleAnswers(getRule(utils.GetIntFromString(item), rules), rules)...)
	}

	fmt.Println("getAnswers return: ", result)

	return result
}

func updateAnswers(answers []string, newPart string) []string {
	fmt.Println("updateAnswers: ", answers, newPart)

	answrs := answers[:]

	if len(answrs) == 0 {
		return append(answrs, newPart)
	}

	for i := range answrs {
		answrs[i] = answrs[i] + newPart
	}

	return answrs
}

func createRule(input string) rule {
	r := regexp.MustCompile((`\d:.`))

	index := utils.GetIntFromString(strings.Replace(r.FindString(input), ":", "", 1))

	return rule{
		index,
		parseConditions(r.ReplaceAllString(input, "")),
	}
}

func parseConditions(input string) [][]string {
	r := regexp.MustCompile(`[abc]|(\d)[^(|)]+\d`)

	conditions := r.FindAllStringSubmatch(input, 2)

	for i := range conditions {
		conditions[i] = strings.Split(conditions[i][0], " ")
	}

	return conditions
}

func getRule(index int, rules []rule) rule {
	for _, rule := range rules {
		if rule.index == index {
			return rule
		}
	}

	log.Fatalf("found no rules, returning default")
	return rules[0]
}
