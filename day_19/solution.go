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

func getPossibleAnswers(rule rule, rules []rule) []string {
	fmt.Println("getPossibleAnswers: ", rule)

	var result []string

	for _, condition := range rule.conditions {
		var part []string

		for _, check := range condition {
			if check == "a" || check == "b" {
				part = updateAnswers(part, check)
			} else {
				newAnswers := getPossibleAnswers(getRule(utils.GetIntFromString(check), rules), rules)

				fmt.Println("newAnswers", newAnswers)

				if (len(newAnswers) >=  2 && len(part) >= 2) {
					for i, item := range newAnswers {
						part[i] = part[i] + item
					}
				} else {
					part = append(part, newAnswers...)
				}
		
				fmt.Println("newPart", part)

			}
		}

		fmt.Println("currResult", result)
		fmt.Println("partToJoin", part)

		result = append(result, strings.Join(part[:], ""))

		fmt.Println("result after check: ", result)
	}

	fmt.Println("return: ", result, rule)

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
