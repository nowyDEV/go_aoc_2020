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
	_, winningCardsRecursive := playRecursiveCombat(players[0], players[1])
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

var gameLoop []string

func playRecursiveCombat(playerOneCards []int, playerTwoCards []int) (winner string, cards []int) {
	var deckHistory []string

	for true {
		if isInfiniteGame(playerOneCards, playerTwoCards, deckHistory) {
			fmt.Println("INFINITE GAME")
			fmt.Println("------------------------------------------------")

			return "playerOne", playerOneCards
		}

		if len(playerOneCards) == 0 {
			return "playerTwo", playerTwoCards
		}

		if len(playerTwoCards) == 0 {
			return "playerOne", playerOneCards
		}

		playerOnePick := playerOneCards[0]
		playerTwoPick := playerTwoCards[0]

		deckHistory = append(deckHistory, convertCardsToString(playerOneCards)+"x"+convertCardsToString(playerTwoCards))

		if len(playerOneCards) > playerOnePick && len(playerTwoCards) > playerTwoPick {
			newCardsPlayerOne := make([]int, playerOnePick)
			newCardsPlayerTwo := make([]int, playerTwoPick)

			copy(newCardsPlayerOne, playerOneCards[1:playerOnePick+1])
			copy(newCardsPlayerTwo, playerTwoCards[1:playerTwoPick+1])

			// playerOne does not have highest card
			if utils.ContainsInt(newCardsPlayerOne, utils.GetHighestNumber(append(newCardsPlayerOne, newCardsPlayerTwo...))) {
				winner = "playerOne"
			} else {
				subGameWinner, _ := playRecursiveCombat(newCardsPlayerOne, newCardsPlayerTwo)
				winner = subGameWinner
			}
		} else {
			if playerOnePick > playerTwoPick {
				winner = "playerOne"
			}

			if playerOnePick < playerTwoPick {
				winner = "playerTwo"
			}
		}

		if winner == "playerOne" {
			playerOneCards = append(playerOneCards[1:], playerOnePick, playerTwoPick)
			playerTwoCards = playerTwoCards[1:]
		}
		if winner == "playerTwo" {
			playerOneCards = playerOneCards[1:]
			playerTwoCards = append(playerTwoCards[1:], playerTwoPick, playerOnePick)
		}
	}

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
