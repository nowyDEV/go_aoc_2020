package main

import (
	"aoc-2020/utils"
	"fmt"
	"log"
	"regexp"
	"strings"
)

type bag struct {
	name string
	children []string
}

type bagTree struct {
	root *bag
}

func main() {
	formData := utils.ReadTextFile("./data.txt")
	tree := createTree(formData)


	count := 0
	for _, bagData := range tree {
		if isOutermostBag("shiny gold bag", bagData, tree) {
			count = count + 1
		}
	}

	fmt.Println(count)
}

func createTree(formData []string) []*bag {
	result := make([]*bag, 0)

	for _, bagData := range formData {
		name := strings.Split(bagData, "s contain ")[0]
		
		children := getBags(strings.Split(bagData, "s contain ")[1])
		fmt.Println("getBags", children)

		result = append(result, &bag{ name, children })
	}

	return result
}

const emptyBag string = "no other bag"

func isOutermostBag(childName string, bagToCheck *bag, bags []*bag) bool {
  for _, child := range bagToCheck.children {
		if (child == emptyBag) {
			continue
		}
		if (child == childName) {
      return true
		}
		return isOutermostBag(childName, getBag(child, bags), bags)
	}
	return false
}

func getBag(name string, bags []*bag) *bag {
	for _, bag := range bags {
		if bag.name == name {
      return bag
		}
	}

	log.Fatalf("bag not found")
	return bags[0]
}

func getBags(input string) []string {
	r := regexp.MustCompile("(bags?).")
	d := regexp.MustCompile("[0-9,.] ")

	formatted := r.ReplaceAllString(input, "bag")
	formatted = d.ReplaceAllString(formatted, "")

	arr := strings.Split(formatted, " bag")
	var newArr []string
	
	for _, item := range arr {
		trimmed := strings.Trim(item, "\t \n")
		
		if (trimmed != "") {
			newArr = append(newArr, trimmed + " bag")
		}
	}

	return newArr
}
