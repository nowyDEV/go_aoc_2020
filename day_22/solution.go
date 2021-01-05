package main

import (
	"aoc-2020/utils"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	data := utils.ReadTextFile("./data.txt")

	players := parseData(data)
	winningCards := playRegularCombat(players[0], players[1])
	_, winningCardsRecursive := playRecursiveCombat(players[0], players[1], []string{})
	fmt.Println(getScore(winningCards))
	fmt.Println(getScore(winningCardsRecursive))
}

func getScore(cards []int) (score int) {
	reversedCards := utils.ReverseNumbers(cards)

	for i, card := range reversedCards {
		score = score + (i+1)*card
	}

	return score
}

func playRecursiveCombat(playerOne []int, playerTwo []int, deckHistory []string) (winner string, cards []int) {
	if isInfiniteGame(playerOne, playerTwo, deckHistory) {
		fmt.Println("INFINITE GAME")
		fmt.Println("------------------------------------------------")

		return "playerOne", playerOne
	}

	if len(playerOne) == 0 {
		return "playerTwo", playerTwo
	}

	if len(playerTwo) == 0 {
		return "playerOne", playerOne
	}

	playerOnePick := playerOne[0]
	playerTwoPick := playerTwo[0]

	deckHistory = append(deckHistory, convertCardsToString(playerOne) + "x" + convertCardsToString(playerTwo))

	if playerOnePick <= len(playerOne)-1 && playerTwoPick <= len(playerTwo)-1 {
		// winner = "playerOne"

		subGameWinner, _ := playRecursiveCombat(playerOne[1:], playerTwo[1:], []string{})
		winner = subGameWinner
		// playerOne does not have highest card
		// if utils.ContainsInt(playerOne[1:], utils.GetHighestNumber(append(playerOne[1:], playerTwo[1:]...))) == false {
		// 	subGameWinner, _ := playRecursiveCombat(playerOne[1:], playerTwo[1:], deckHistory{[]string{}, []string{}})
		// 	winner = subGameWinner
		// } 
	} else {
		if playerOne[0] > playerTwo[0] {
			winner = "playerOne"
		}

		if playerOne[0] < playerTwo[0] {
			winner = "playerTwo"
		}
	}

	if winner == "playerOne" {
		newPlayerOne := append(playerOne[1:], playerOne[0], playerTwo[0])
		newPlayerTwo := playerTwo[1:]
		return playRecursiveCombat(newPlayerOne, newPlayerTwo, deckHistory)
	}
	if winner == "playerTwo" {
		newPlayerOne := playerOne[1:]
		newPlayerTwo := append(playerTwo[1:], playerTwo[0], playerOne[0])
		return playRecursiveCombat(newPlayerOne, newPlayerTwo, deckHistory)
	}

	fmt.Println("EMPTY RETURN")
	fmt.Println("------------------------------------------------")
	return
}

func isInfiniteGame(playerOneDeck []int, playerTwoDeck []int, deckHistory []string) bool {
	cardStr := convertCardsToString(playerOneDeck) + "x" + convertCardsToString(playerTwoDeck)

	return utils.ContainsString(deckHistory, cardStr)
}

func convertCardsToString(cards []int) string {
	var cardsStr []string

	for _, card := range cards {
		cardStr := strconv.Itoa(card)
		cardsStr = append(cardsStr, cardStr)
	}

	return strings.Join(cardsStr, ",")
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
