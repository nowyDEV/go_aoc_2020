package main

import (
	"aoc-2020/utils"
	"fmt"
	"regexp"
	"strings"
)

func main() {
	data := utils.ReadTextFile("./data_test.txt")

	foods := parseData(data)
	ingredients := getIngredients(foods, []ingredient{})
	fmt.Println("final ingredients", ingredients)

	count := 0

	for _, food := range foods {
		for _, ingredient := range food.ingredients {
			if isAllergicIngredient(ingredients, ingredient) == false {
				count++
			}
		}
	}

	fmt.Println(count)
}

type food struct {
	id          int
	ingredients []string
	allergens   []string
}

type ingredient struct {
	name     string
	allergen string
}

func isAllergicIngredient(ingredients []ingredient, name string) bool {
	for _, ingredient := range ingredients {
		if ingredient.name == name {
			return true
		}
	}

	return false
}

func getIngredients(input []food, ingredients []ingredient) []ingredient {
	var result = append(ingredients)
	hasChanged := false

	for _, food := range input {
		for _, secFood := range input {
			if food.id == secFood.id {
				continue
			}

			foundIngredientsWithAllergens := findIngredientsWithAllergens(secFood, result)
			if len(foundIngredientsWithAllergens) > 0 {
				result = append(result, foundIngredientsWithAllergens...)
				fmt.Println("foundIngredientsWithAllergens updatedResult", result)
				fmt.Println("----------------------------------------------------")
				hasChanged = true
			}

			fmt.Println("compare", food, secFood)

			commonIngredients := getCommonItems(food.ingredients, secFood.ingredients)
			fmt.Println("commonIngredients", commonIngredients)
			if len(commonIngredients) == 0 {
				continue
			}

			commonAllergens := getCommonItems(food.allergens, secFood.allergens)
			fmt.Println("commonAllergens", commonAllergens)

			clearIngredients := getClearIngredients(commonIngredients, result)
			fmt.Println("clearIngredients", clearIngredients)
			if len(clearIngredients) == 1 && len(commonAllergens) == 1 {
				result = append(result, ingredient{clearIngredients[0], commonAllergens[0]})
				fmt.Println("updatedResult", result)
				fmt.Println("----------------------------------------------------")
				hasChanged = true
			}
		}
	}

	if (hasChanged) {
		return getIngredients(input, result)
	}

	return result
}

func findIngredientsWithAllergens(food food, ingredients []ingredient) []ingredient {
	remainingAllergens := append(food.allergens)
	remainingIngredients := append(food.ingredients)

	for _, foodIngredient := range food.ingredients {
		for _, ingredient := range ingredients {
			if foodIngredient == ingredient.name {
				remainingIngredients = append(utils.FilterOut(remainingIngredients, ingredient.name))
				remainingAllergens = append(utils.FilterOut(remainingAllergens, ingredient.allergen))
			}
		}
	}

	if len(remainingAllergens) == 1 && len(remainingIngredients) == 1 {
		return []ingredient{{remainingIngredients[0], remainingAllergens[0]}}
	}

	return []ingredient{}
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

func getCommonItems(listOne []string, listTwo []string) []string {
	var result []string

	for _, item := range listOne {
		if utils.ContainsString(listTwo, item) {
			result = append(result, item)
		}
	}

	return result
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
