package main

import (
	"strconv"
	"strings"
)

func countAntinodes(frequencyCoordinates []coordinates, maxX int, maxY int, visited map[coordinates]bool) int {
	count := 0
	for _, currentCoordinates := range frequencyCoordinates {
		for _, comparisonCoordinates := range frequencyCoordinates {
			if comparisonCoordinates != currentCoordinates {
				antinode := coordinates{x: currentCoordinates.x - (comparisonCoordinates.x - currentCoordinates.x), y: currentCoordinates.y - (comparisonCoordinates.y - currentCoordinates.y)}
				_, exists := visited[antinode]
				if !exists && antinode.x >= 0 && antinode.x < maxX && antinode.y >= 0 && antinode.y < maxY {
					visited[antinode] = true
					count++
				}
			}
		}
	}
	return count
}

func day8problem1() string {
	inputString, err := readFile("dayeight.txt")
	if err != nil {
		return "Could not read file"
	}

	lines := strings.Split(inputString, "\n")

	inputMatrix := [][]string{}

	visited := make(map[coordinates]bool)

	antennaMap := make(map[string][]coordinates)

	for i, line := range lines {
		inputMatrix = append(inputMatrix, strings.Split(line, ""))
		for j, character := range inputMatrix[i] {
			if character != "." {
				antennaMap[character] = append(antennaMap[character], coordinates{x: i, y: j})
			}
		}
	}

	maxX := len(inputMatrix)
	maxY := len(inputMatrix[0])

	count := 0

	for _, frequencyCoordinates := range antennaMap {
		count += countAntinodes(frequencyCoordinates, maxX, maxY, visited)
	}

	return strconv.Itoa(count)
}
