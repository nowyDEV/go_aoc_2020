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
	conditions string
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
	fmt.Println(getFilledRules(rules))
}

func fillRules(rules []rule, possibleFills []int) []rule {
	for _, item := range rules {
		r := regexp.MustCompile("[0-9]")
		conds := strings.Split(item.conditions, "")

		for i := 0; i < len(conds);i++ {
			if (r.MatchString(conds[i]) && utils.ContainsInt(possibleFills, utils.GetIntFromString(conds[i]))) {
				conds[i] = getRule(utils.GetIntFromString(conds[i]), rules).conditions
			}
		}
	}

	return rules
}

func getFilledRules(rules []rule) []int {
	r := regexp.MustCompile("[0-9]")
	var result []int

	for _, item := range rules {
		if (r.MatchString(item.conditions) == false) {
			result = append(result, item.index)
		}
	}

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
	r := regexp.MustCompile((`\d:.|\"`))

	index := utils.GetIntFromString(strings.Replace(r.FindString(input), ":", "", 1))

	return rule{
		index,
		r.ReplaceAllString(input, ""),
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
