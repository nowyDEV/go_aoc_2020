package main

import (
	"aoc-2020/utils"
	"fmt"
	"regexp"
	"strings"
)

const imageSize = 10

type tile struct {
	id      int
	borders []string
}

func main() {
	data := utils.ReadTextFile("./data.txt")

	tiles := parseTiles(data)

	result := 1

	for i := 0; i < len(tiles); i++ {
		if isEdgeTile(tiles[i], filterTiles(tiles, tiles[i].id)) {
			result = result * tiles[i].id
		}
	}

	fmt.Println(result)
}

func filterTiles(tiles []tile, filterID int) []tile {
	var result []tile

	for _, item := range tiles {
		if item.id != filterID {
			result = append(result, item)
		}
	}

	return result
}

func parseTiles(input []string) []tile {
	r := regexp.MustCompile("[0-9]")
	var result []tile

	for i, item := range input {
		if r.MatchString(item) {
			tile := tile{utils.GetIntFromString(item), getBorders(input[i+1 : i+11])}
			result = append(result, tile)
		}
	}

	return result
}

func getBorders(input []string) []string {
	result := append([]string{}, input[0], utils.ReverseString(input[0]), input[len(input)-1], utils.ReverseString(input[len(input)-1]))

	var left []string
	var right []string

	for _, item := range input {
		left = append(left, item[0:1])
		right = append(right, item[len(item)-1:])
	}

	leftStr := strings.Join(left, "")
	rightStr := strings.Join(right, "")

	return append(result, leftStr, utils.ReverseString(leftStr), rightStr, utils.ReverseString(rightStr))
}

func isEdgeTile(tile tile, tiles []tile) bool {
	matchedBorders := 0

	for _, border := range tile.borders {
		for _, tileItem := range tiles {
			if utils.ContainsString(tileItem.borders, border) {
				matchedBorders++
			}
		}
	}

	return matchedBorders <= 4
}
