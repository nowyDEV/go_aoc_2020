package main

import (
	"aoc-2020/utils"
	"fmt"
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

	for _, rule := range rules {
		fmt.Println(rule)
	}

	fmt.Println(answers)
}

func createRule(input string) rule {
	r := regexp.MustCompile((`\d:.`))
	rb := regexp.MustCompile(`[a-b]`)

	index := utils.GetIntFromString(strings.Replace(r.FindString(input), ":", "", 1) )

	return rule{
		index,
		parseConditions(r.ReplaceAllString(input, "")),
		rb.FindString(input),
	}
}

func parseConditions(input string) [][]string {
	r := regexp.MustCompile(`\d[^(|)]+\d`)

	return r.FindAllStringSubmatch(input, 2)
}
