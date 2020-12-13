package main

import (
	"aoc-2020/utils"
	"fmt"
	"log"
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
	data := utils.ReadTextFile("./data.txt")

	fmt.Println(getCountBeforeInfiniteLoop(0, nil, data, 0))
	fmt.Println(getSumOfValidCodeArguments(data))
}

func getItemToSwap(data []string) int {
	for index, command := range data {
		name, value := getCommandValues(command)
		secData := make([]string, len(data))
		copy(secData, data)

		if name == "jmp" {
			secData[index] = fmt.Sprint("nop ", value)

			if isInfinite(0, nil, secData) {
				continue
			} else {
				return index
			}
		}

		if name == "nop" {
			secData[index] = fmt.Sprint("jmp ", value)

			if isInfinite(0, nil, secData) {
				continue
			} else {
				return index
			}
		}
	}

	return -1
}

func getSumOfValidCodeArguments(data []string) int {
	itemToSwap := getItemToSwap(data)

	name, value := getCommandValues(data[itemToSwap])

	if name == "jmp" {
		data[itemToSwap] = fmt.Sprint("nop ", value)
	}
	if name == "nop" {
		data[itemToSwap] = fmt.Sprint("jmp ", value)
	}

	return getCountBeforeSuccessTermination(0, data, 0)
}

func isInfinite(commandIndex int, executedCommands []int, commands []string) bool {
	if commandIndex >= len(commands) {
		return false
	}

	name, value := getCommandValues(commands[commandIndex])

	if name == "jmp" {
		if utils.ContainsInt(executedCommands, commandIndex) {
			return true
		}

		return isInfinite(commandIndex+value, append(executedCommands, commandIndex), commands)
	}

	return isInfinite(commandIndex+1, executedCommands, commands)
}

func getCountBeforeSuccessTermination(commandIndex int, commands []string, count int) int {
	if commandIndex >= len(commands) {
		return count
	}

	name, value := getCommandValues(commands[commandIndex])

	if name == "jmp" {
		return getCountBeforeSuccessTermination(commandIndex+value, commands, count)
	}
	if name == "nop" {
		return getCountBeforeSuccessTermination(commandIndex+1, commands, count)
	}
	return getCountBeforeSuccessTermination(commandIndex+1, commands, count+value)
}

func getCountBeforeInfiniteLoop(commandIndex int, executedCommands []int, commands []string, count int) int {
	if utils.ContainsInt(executedCommands, commandIndex) {
		return count
	}

	name, value := getCommandValues(commands[commandIndex])

	if name == "jmp" {
		return getCountBeforeInfiniteLoop(commandIndex+value, append(executedCommands, commandIndex), commands, count)
	}
	if name == "acc" {
		return getCountBeforeInfiniteLoop(commandIndex+1, append(executedCommands, commandIndex), commands, count+value)
	}
	return getCountBeforeInfiniteLoop(commandIndex+1, append(executedCommands, commandIndex), commands, count)
}

func getCommandValues(command string) (string, int) {
	arr := strings.Split(command, " ")
	value, err := strconv.Atoi(arr[1])

	if err != nil {
		log.Fatalf("failed to convert")
	}

	return arr[0], value
}
