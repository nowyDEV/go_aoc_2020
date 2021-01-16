package main

import (
	"aoc-2020/utils"
	"fmt"
	"math"
	"strings"
)

type memAddress struct {
	address int
	value   int
}

const bitSize = 36

func main() {
	data := utils.ReadTextFile("./data.txt")

	memArr := calculateMemory(data)
	fmt.Println("sum", calculateSum(memArr))
}

func calculateSum(memArr []memAddress) (result int) {
	for _, item := range memArr {
		result = result + item.value
	}

	return result
}

func calculateValue(mask string, value int) int {
	bitmask := getBitmask(value)
	calc := calculateBitmask(mask, bitmask)

	return getValueFromBitmask(calc)
}

func calculateMemory(input []string) (result []memAddress) {
	currMask := ""

	for _, row := range input {
		if isMask(row) {
			currMask = getMask(row)
			continue
		}
		memAddr, value := parseMemoryAssignment(row)
		memIndex := findMemIndex(result, memAddr)
		newValue := calculateValue(currMask, value)

		if (memIndex) != -1 {
			result[memIndex].value = newValue
		} else {
			result = append(result, memAddress{memAddr, newValue})
		}
	}

	return result
}

func findMemIndex(memArr []memAddress, address int) int {
	for index, item := range memArr {
		if item.address == address {
			return index
		}
	}

	return -1
}

func parseMemoryAssignment(input string) (memAddress int, value int) {
	items := strings.Split(input, " = ")
	memAddress = utils.GetIntFromString(items[0])
	value = utils.GetIntFromString(items[1])

	return memAddress, value
}

func getValueFromBitmask(bitmask string) (result int) {
	for i := len(bitmask) - 1; i >= 0; i-- {
		if bitmask[i:i+1] == "1" {
			power := float64((len(bitmask) - 1) - i)
			result = result + int(math.Pow(2, power))
		}
	}

	return result
}

func calculateBitmask(mask string, value string) string {
	bitValue := ""

	for i := 0; i < bitSize; i++ {
		maskBit := mask[i : i+1]
		valueBit := value[i : i+1]

		if maskBit == "X" {
			bitValue = bitValue + valueBit
		} else {
			bitValue = bitValue + maskBit
		}
	}

	return bitValue
}

func getBitmask(value int) string {
	bits := fmt.Sprintf("%b", value)

	zeroesToAdd := bitSize - len(bits)

	for i := 0; i < zeroesToAdd; i++ {
		bits = "0" + bits
	}

	return bits
}

func isMask(input string) bool {
	return strings.Contains(input, "mask")
}

func getMask(input string) string {
	return strings.Replace(input, "mask = ", "", 1)
}
