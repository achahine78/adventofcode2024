package main

import (
	"strconv"
	"strings"
)

func getNeighbours(i, j int, inputMatrix [][]int) [][]int {
	neighbours := [][]int{}
	if i > 0 && inputMatrix[i-1][j]-inputMatrix[i][j] == 1 {
		neighbours = append(neighbours, []int{i - 1, j})
	}
	if j > 0 && inputMatrix[i][j-1]-inputMatrix[i][j] == 1 {
		neighbours = append(neighbours, []int{i, j - 1})
	}
	if i < len(inputMatrix)-1 && inputMatrix[i+1][j]-inputMatrix[i][j] == 1 {
		neighbours = append(neighbours, []int{i + 1, j})
	}
	if j < len(inputMatrix[i])-1 && inputMatrix[i][j+1]-inputMatrix[i][j] == 1 {
		neighbours = append(neighbours, []int{i, j + 1})
	}

	return neighbours
}

func countPathsToPeaks(i, j int, inputMatrix [][]int) int {
	visited := make(map[coordinates]bool)
	peakCount := 0

	neighbourQueue := queue{}
	neighbourQueue.enqueue([]int{i, j})

	for neighbourQueue.length() != 0 {
		current := neighbourQueue.dequeue()
		_, exists := visited[coordinates{x: current[0], y: current[1]}]
		if !exists && inputMatrix[current[0]][current[1]] == 9 {
			peakCount++
		}
		visited[coordinates{x: current[0], y: current[1]}] = true
		neighbours := getNeighbours(current[0], current[1], inputMatrix)
		for _, neighbour := range neighbours {
			neighbourQueue.enqueue(neighbour)
		}
	}

	return peakCount

}

func day10problem1() string {
	inputString, err := readFile("dayten.txt")
	if err != nil {
		return "Could not read file"
	}

	lines := strings.Split(inputString, "\n")
	inputMatrix := [][]int{}

	for _, line := range lines {
		splitLine := strings.Split(line, "")
		splitLineInteger := []int{}
		for _, character := range splitLine {
			characterInteger, characterIntegerError := strconv.Atoi(character)
			if characterIntegerError != nil {
				return "Invalid input"
			}

			splitLineInteger = append(splitLineInteger, characterInteger)
		}
		inputMatrix = append(inputMatrix, splitLineInteger)
	}
	pathCount := 0
	for i := 0; i < len(inputMatrix); i++ {
		for j := 0; j < len(inputMatrix[i]); j++ {
			if inputMatrix[i][j] == 0 {
				pathCount += countPathsToPeaks(i, j, inputMatrix)
			}
		}
	}

	return strconv.Itoa(pathCount)
}
