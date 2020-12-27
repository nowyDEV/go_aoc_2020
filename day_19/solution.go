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
	base       string
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

	result := 0

	for _, answer := range answers {
		if validateAnswer(answer, rules) {
			result++
		}
	}

	fmt.Println(result)
}

func validateAnswer(answer string, rules []rule) bool {
	isValid, _ := validateRule(answer, getRule(0, rules), rules)
	return isValid
}

func validateRule(answer string, rule rule, rules []rule) (bool, string) {
	fmt.Println("validateRule", answer, rule)
	if len(answer) <= 0 {
		return false, ""
	}

	if len(rule.base) > 0 {
		index := utils.FindIndex(strings.Split(answer, ""), rule.base)
		if index < 0 {
			return false, ""
		}
		return true, answer[index+1:]
	}

	if len(rule.conditions) == 2 {
		isValid, restAnswer := validateCondition(answer, rule.conditions[0], rules)

		if isValid {
			return true, restAnswer
		}
		return validateCondition(answer, rule.conditions[1], rules)
	}

	return validateCondition(answer, rule.conditions[0], rules)
}

func validateCondition(answer string, condition []string, rules []rule) (bool, string) {
	fmt.Println("validateCondition", answer, condition)
	answr := answer[0:]

	for i := 0; i < len(condition); i++ {
		rule := getRule(utils.GetIntFromString(condition[i]), rules)
		isValid, restAnswer := validateRule(answr, rule, rules)

		if isValid {
			answr = restAnswer
		} else {
			return false, ""
		}
	}

	fmt.Println("isValid", true, answr)

	return true, answr
}

func createRule(input string) rule {
	r := regexp.MustCompile((`\d:.`))
	rb := regexp.MustCompile(`[a-b]`)

	index := utils.GetIntFromString(strings.Replace(r.FindString(input), ":", "", 1))

	return rule{
		index,
		parseConditions(r.ReplaceAllString(input, "")),
		rb.FindString(input),
	}
}

func parseConditions(input string) [][]string {
	r := regexp.MustCompile(`\d[^(|)]+\d`)

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
