package main

import (
	"aoc-2020/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type bag struct {
	name     string
	value    int
	children []bag
}

const emptyBag string = "no other bag"

func main() {
	formData := utils.ReadTextFile("./data.txt")
	tree := createTree(formData)

	fmt.Println(findNumOfOuterBags("shiny gold bag", tree))
	fmt.Println(getNumOfContainedBags("shiny gold bag", tree))

}

func getNumOfContainedBags(bagName string, tree []*bag) int {
	result := 0
	bag := getBag(bagName, tree)

	for _, child := range bag.children {
		if child.value == 0 {
			continue
		}
		result = result + child.value + (child.value * getNumOfContainedBags(child.name, tree))
	}

	return result
}

func isOutermostBag(searchBag string, bagToCheck *bag, bags []*bag) bool {
	for _, child := range bagToCheck.children {
		if child.name == emptyBag {
			continue
		}
		if child.name == searchBag {
			return true
		}
		if isOutermostBag(searchBag, getBag(child.name, bags), bags) {
			return true
		}
		continue
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

func createTree(formData []string) []*bag {
	result := make([]*bag, 0)

	for _, bagData := range formData {
		name := strings.Split(bagData, "s contain ")[0]

		children := getBags(strings.Split(bagData, "s contain ")[1])

		result = append(result, &bag{name, 1, children})
	}

	return result
}

func getBags(input string) []bag {
	r := regexp.MustCompile("(bags?).")
	d := regexp.MustCompile("[0-9,.] ")
	n := regexp.MustCompile("[0-9]")

	var children []bag

	values := n.FindAllString(input, 10)

	formatted := r.ReplaceAllString(input, "bag")
	formatted = d.ReplaceAllString(formatted, "")

	arr := strings.Split(formatted, " bag")
	var newArr []bag

	for index, item := range arr {
		trimmed := strings.Trim(item, "\t \n")

		if trimmed != "" {
			var bagValue = 0

			if len(values) > 0 {
				num, err := strconv.Atoi(values[index])

				if err != nil {
					log.Fatalf("failed to convert")
				}
				bagValue = num
			}

			newArr = append(newArr, bag{trimmed + " bag", bagValue, children})
		}
	}

	return newArr
}

func findNumOfOuterBags(bag string, tree []*bag) int {
	count := 0

	for _, bagData := range tree {
		if isOutermostBag("shiny gold bag", bagData, tree) {
			count = count + 1
		}
	}

	return count
}
