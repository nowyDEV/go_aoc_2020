package main

import (
	"aoc-2020/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	data := utils.ReadTextFile("./data.txt")

	rules := make(map[int]string)
	var answers []string

	for _, line := range data {
		if utils.ContainsString(strings.Split(line, ""), ":") {
			key, value := parseRule(line)
			rules[key] = value
		} else if len(line) > 0 {
			answers = append(answers, line)
		}
	}

	computedRule := computeRules(rules[0], rules)
	r := regexp.MustCompile("^" + computedRule + "$")

	result := 0
	for _, answer := range answers {
		if r.MatchString(answer) {
			result++
		}
	}
	fmt.Println("part 1 result: ", result)
}

func parseRule(input string) (key int, value string) {
	items := strings.Split(input, ": ")

	return utils.GetIntFromString(items[0]), items[1]
}

var rPipe = regexp.MustCompile(`\|`)
var ruleToRegexp = make(map[string]string)

func computeRules(value string, rules map[int]string) (result string) {
	mapItemValue, ok := ruleToRegexp[value]
	if ok {
		return mapItemValue
	}

	match, _ := regexp.MatchString("^\".*\"$", value)
	if match {
		result = strings.Replace(value, "\"", "", 2)
	} else if rPipe.MatchString(value) {
		options := strings.Split(value, " | ")
		result = fmt.Sprintf("(%s|%s)", computeRules(options[0], rules), computeRules(options[1], rules))
	} else {
		result = computeKeys(value, rules)
	}

	ruleToRegexp[value] = result
	return result
}

func computeKeys(input string, rules map[int]string) string {
	keys := strings.Split(input, " ")
	var res []string

	for _, key := range keys {
		res = append(res, computeRules(rules[utils.GetIntFromString(key)], rules))
	}

	return strings.Join(res, "")
}
