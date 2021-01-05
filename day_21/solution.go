package main

import (
	"aoc-2020/utils"
	"fmt"
	"log"
	"regexp"
	"sort"
	"strings"
)

type food struct {
	id          int
	ingredients []string
	allergens   []string
}

type ingredient struct {
	name     string
	allergen string
}

func main() {
	data := utils.ReadTextFile("./data.txt")

	foods := parseData(data)

	allergens := getListOfAllergens(foods)
	allergicIngredients := getAllergicIngredients(foods, allergens, []ingredient{})

	fmt.Println("allergens", allergens)
	fmt.Println("allergicIngredients", allergicIngredients)

	count := 0
	for _, food := range foods {
		for _, ingredient := range food.ingredients {
			if isAllergicIngredient(allergicIngredients, ingredient) == false {
				count++
			}
		}
	}

	sort.Sort(alphabetic(allergens))

	fmt.Println("part 1 result", count)
	fmt.Println("part 2 result", generateCanonicalDangerousList(allergens, allergicIngredients))
}

type alphabetic []string

func (list alphabetic) Len() int { return len(list) }

func (list alphabetic) Swap(i, j int) { list[i], list[j] = list[j], list[i] }

func (list alphabetic) Less(i, j int) bool {
	var si string = list[i]
	var sj string = list[j]
	var siLower = strings.ToLower(si)
	var sjLower = strings.ToLower(sj)
	if siLower == sjLower {
		return si < sj
	}
	return siLower < sjLower
}

func generateCanonicalDangerousList(allergens []string, ingredients []ingredient) string {
	result := ""

	for _, allergen := range allergens {
		result = result + getIngredientByAllergen(ingredients, allergen).name + ","
	}

	return result[:len(result)-1]
}

func getIngredientByAllergen(ingredients []ingredient, allergen string) ingredient {
	for _, ingredient := range ingredients {
		if ingredient.allergen == allergen {
			return ingredient
		}
	}

	log.Fatalf("ingredient not found")
	return ingredient{}
}

func getListOfAllergens(foods []food) []string {
	var result []string

	for _, food := range foods {
		for _, allergen := range food.allergens {
			if utils.ContainsString(result, allergen) == false {
				result = append(result, allergen)
			}
		}
	}
	return result
}

func getAllergicIngredients(input []food, allergens []string, ingredients []ingredient) []ingredient {
	var allergicIngredients = append(ingredients)
	hasChanged := false

	for _, allergen := range allergens {
		foodsWithAllergen := getFoodsWithAllergen(input, allergen)
		commonIngredients := getCommonIngredients(foodsWithAllergen)
		clearIngredients := getClearIngredients(commonIngredients, allergicIngredients)

		if len(clearIngredients) == 1 {
			allergicIngredients = append(allergicIngredients, ingredient{clearIngredients[0], allergen})
			hasChanged = true
		}
	}

	if hasChanged {
		return getAllergicIngredients(input, allergens, allergicIngredients)
	}

	return allergicIngredients
}

func getFoodsWithAllergen(foods []food, allergen string) []food {
	var result []food

	for _, food := range foods {
		if utils.ContainsString(food.allergens, allergen) {
			result = append(result, food)
		}
	}

	return result
}

func getCommonIngredients(foods []food) []string {
	if len(foods) == 0 {
		return []string{}
	}

	var result []string

	for index, food := range foods {
		if index == 0 {
			result = append(result, food.ingredients...)
			continue
		}

		for _, ingredient := range result {
			if utils.ContainsString(food.ingredients, ingredient) == false {
				result = append(utils.FilterOut(result, ingredient))
			}
		}
	}

	return result
}

func getClearIngredients(names []string, ingredients []ingredient) []string {
	var result []string

	if len(ingredients) == 0 {
		return names
	}

	for _, name := range names {
		isClear := true

		for _, ingredient := range ingredients {
			if name == ingredient.name {
				isClear = false
			}
		}

		if isClear {
			result = append(result, name)
		}
	}

	return result
}

func isAllergicIngredient(ingredients []ingredient, name string) bool {
	for _, ingredient := range ingredients {
		if ingredient.name == name {
			return true
		}
	}

	return false
}

func parseData(input []string) []food {
	var result []food

	for index, row := range input {
		items := strings.Split(row, " (")

		result = append(result, food{index, strings.Split(items[0], " "), getAllergens(items[1])})
	}

	return result
}

func getAllergens(input string) []string {
	r := regexp.MustCompile("contains |[,)]")

	return strings.Split(r.ReplaceAllString(input, ""), " ")
}
