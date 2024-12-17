package main

import (
	"strconv"
	"strings"
)

func getPlot(initialI, initialJ int, inputMatrix [][]string, visited map[coordinates]bool) [][]int {
	plot := [][]int{}
	plotValue := inputMatrix[initialI][initialJ]
	processedMap := make(map[coordinates]bool)

	queue := queue[coordinates]{}

	queue.enqueue(coordinates{x: initialI, y: initialJ})

	for queue.length() != 0 {
		currentCoordinate := queue.dequeue()
		i := currentCoordinate.x
		j := currentCoordinate.y

		_, processed := processedMap[currentCoordinate]

		if !processed {
			plot = append(plot, []int{currentCoordinate.x, currentCoordinate.y})
			processedMap[currentCoordinate] = true

			visited[currentCoordinate] = true

			if i > 0 && inputMatrix[i-1][j] == plotValue {
				_, exists := visited[coordinates{x: i - 1, y: j}]

				if !exists {
					queue.enqueue(coordinates{x: i - 1, y: j})
				}
			}
			if j > 0 && inputMatrix[i][j-1] == plotValue {
				_, exists := visited[coordinates{x: i, y: j - 1}]

				if !exists {
					queue.enqueue(coordinates{x: i, y: j - 1})
				}
			}
			if i < len(inputMatrix)-1 && inputMatrix[i+1][j] == plotValue {
				_, exists := visited[coordinates{x: i + 1, y: j}]

				if !exists {
					queue.enqueue(coordinates{x: i + 1, y: j})
				}
			}
			if j < len(inputMatrix[i])-1 && inputMatrix[i][j+1] == plotValue {
				_, exists := visited[coordinates{x: i, y: j + 1}]

				if !exists {
					queue.enqueue(coordinates{x: i, y: j + 1})
				}
			}
		}
	}

	return plot

}

func computePriceWithPerimeter(plot [][]int) int {
	coordinateMap := make(map[coordinates]bool)
	for _, unit := range plot {
		coordinateMap[coordinates{x: unit[0], y: unit[1]}] = true
	}

	sharedEdges := 0

	for coordinate := range coordinateMap {
		_, bottomExists := coordinateMap[coordinates{coordinate.x + 1, coordinate.y}]
		_, rightExists := coordinateMap[coordinates{coordinate.x, coordinate.y + 1}]

		if bottomExists {
			sharedEdges++
		}

		if rightExists {
			sharedEdges++
		}
	}

	perimeter := 4*len(plot) - 2*sharedEdges

	return perimeter * len(plot)
}

func day12problem1() string {
	inputString, err := readFile("daytwelve.txt")
	if err != nil {
		return "Could not read file"
	}

	lines := strings.Split(inputString, "\n")

	inputMatrix := [][]string{}
	visited := make(map[coordinates]bool)

	for _, line := range lines {
		inputMatrix = append(inputMatrix, strings.Split(line, ""))
	}

	priceSum := 0

	for i := 0; i < len(inputMatrix); i++ {
		for j := 0; j < len(inputMatrix[i]); j++ {

			_, exists := visited[coordinates{x: i, y: j}]

			if !exists {
				plot := getPlot(i, j, inputMatrix, visited)
				priceSum += computePriceWithPerimeter(plot)
			}
		}
	}

	return strconv.Itoa(priceSum)
}

func computePriceWithSides(plot [][]int, inputMatrix [][]string) int {
	coordinateMap := make(map[coordinates]bool)
	for _, unit := range plot {
		coordinateMap[coordinates{x: unit[0], y: unit[1]}] = true
	}

	sides := 0

	for i := 0; i < len(inputMatrix); i++ {
		isBoundary := false
		for j := 0; j < len(inputMatrix[i]); j++ {
			if coordinateMap[coordinates{x: i, y: j}] {
				if !isBoundary {
					sides++
					isBoundary = true
				}
			} else {
				isBoundary = false
			}
		}
	}

	for j := 0; j < len(inputMatrix[0]); j++ {
		isBoundary := false
		for i := 0; i < len(inputMatrix); i++ {
			if coordinateMap[coordinates{x: i, y: j}] {
				if !isBoundary {
					sides++
					isBoundary = true
				}
			} else {
				isBoundary = false
			}
		}
	}

	return sides
}

func day12problem2() string {
	inputString, err := readFile("daytwelve.txt")
	if err != nil {
		return "Could not read file"
	}

	lines := strings.Split(inputString, "\n")

	inputMatrix := [][]string{}
	visited := make(map[coordinates]bool)

	for _, line := range lines {
		inputMatrix = append(inputMatrix, strings.Split(line, ""))
	}

	priceSum := 0

	for i := 0; i < len(inputMatrix); i++ {
		for j := 0; j < len(inputMatrix[i]); j++ {

			_, exists := visited[coordinates{x: i, y: j}]

			if !exists {
				plot := getPlot(i, j, inputMatrix, visited)
				priceSum += computePriceWithSides(plot, inputMatrix)
			}
		}
	}

	return strconv.Itoa(priceSum)
}
