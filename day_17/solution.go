package main

import (
	"aoc-2020/utils"
	"fmt"
	"strings"
)

const cycles = 6

func main() {
	data := utils.ReadTextFile("./data.txt")

	cubeGrid := make(map[string]bool)
	generateCubeMap(&cubeGrid, data)

	for i := 0; i < cycles; i++ {
		minX, minY, minZ, maxX, maxY, maxZ := 0, 0, 0, 0, 0, 0

		for key := range cubeGrid {
			coords := strings.Split(key, ",")

			x := utils.GetFullIntFromString(coords[0])
			y := utils.GetFullIntFromString(coords[1])
			z := utils.GetFullIntFromString(coords[2])

			if x < minX {
				minX = x
			}
			if x > maxX {
				maxX = x
			}
			if y < minY {
				minY = y
			}
			if y > maxY {
				maxY = y
			}
			if z < minZ {
				minZ = z
			}
			if z > maxZ {
				maxZ = z
			}
		}

		cycleMap := make(map[string]bool)

		for i := minX - 1; i <= maxX+1; i++ {
			for j := minY - 1; j <= maxY+1; j++ {
				for k := minZ - 1; k <= maxZ+1; k++ {
					neighbors := getNeighbors(i, j, k, cubeGrid)
					numOfActiveNeighbors := getNumOfTrueItems(neighbors)
					key := fmt.Sprintf("%d,%d,%d", i, j, k)

					value, ok := cubeGrid[key]
					isActive := ok && value

					if isActive && numOfActiveNeighbors != 2 && numOfActiveNeighbors != 3 {
						cycleMap[key] = false
					} else if !isActive && numOfActiveNeighbors == 3 {
						cycleMap[key] = true
					} else {
						cycleMap[key] = isActive
					}
				}
			}
		}

		cubeGrid = cycleMap
	}

	activeItems := 0

	for _, value := range cubeGrid {
		if value {
			activeItems++
		}
	}

	fmt.Println(activeItems)
}

func generateCubeMap(cubeGrid *map[string]bool, data []string) {
	for index, item := range data {
		row := strings.Split(item, "")
		for rowIndex, rowItem := range row {
			isActive := rowItem == "#"
			key := fmt.Sprintf("%d,%d,%d", rowIndex, index, 0)

			(*cubeGrid)[key] = isActive
		}
	}
}

func getNeighbors(x, y, z int, cubeMap map[string]bool) (result []bool) {
	for i := x - 1; i <= x+1; i++ {
		for j := y - 1; j <= y+1; j++ {
			for k := z - 1; k <= z+1; k++ {
				if x != i || y != j || z != k {
					key := fmt.Sprintf("%d,%d,%d", i, j, k)
					value, ok := cubeMap[key]

					if ok {
						result = append(result, value)
					} else {
						result = append(result, false)
					}
				}
			}
		}
	}

	return result
}

func getNumOfTrueItems(list []bool) (result int) {
	for _, value := range list {
		if value {
			result++
		}
	}

	return result
}
