package main

import (
	"aoc-2020/utils"
	"fmt"
	"regexp"
)

func main() {
	data := utils.ReadTextFile("./data.txt")

	players := parseData(data)
	winningCards := playRegularCombat(players[0], players[1])
	fmt.Println(getScore(winningCards))
}

func getScore(cards []int) (score int) {
	reversedCards := utils.ReverseNumbers(cards)

	for i, card := range reversedCards {
		score = score + (i+1)*card
	}

	return score
}

func playRegularCombat(playerOne []int, playerTwo []int) (winningCards []int) {
	if len(playerOne) == 0 {
		return playerTwo
	}

	if len(playerTwo) == 0 {
		return playerOne
	}

	if playerOne[0] > playerTwo[0] {
		newPlayerOne := append(playerOne[1:], playerOne[0], playerTwo[0])
		newPlayerTwo := playerTwo[1:]
		return playRegularCombat(newPlayerOne, newPlayerTwo)
	}

	if playerOne[0] < playerTwo[0] {
		newPlayerOne := playerOne[1:]
		newPlayerTwo := append(playerTwo[1:], playerTwo[0], playerOne[0])
		return playRegularCombat(newPlayerOne, newPlayerTwo)
	}

	return
}

func parseData(input []string) (playersCards [][]int) {
	r := regexp.MustCompile("Player")
	index := -1

	for _, row := range input {
		if row == "" {
			continue
		}

		if r.MatchString(row) {
			index++
			playersCards = append(playersCards, []int{})
			continue
		}

		playersCards[index] = append(playersCards[index], utils.GetIntFromString(row))
	}

	return playersCards
}
