package main

import (
	"fmt"
	"strconv"
	"strings"
)

type coordinates struct {
	x int
	y int
}

type directedCoordinates struct {
	x         int
	y         int
	direction string
}

func printMatrix(matrix [][]string, i, j int, direction string, visited map[coordinates]bool) {
	for a := 0; a < len(matrix); a++ {
		for b := 0; b < len(matrix[a]); b++ {
			if a == i && b == j {
				if direction == "up" {
					fmt.Print("^")
				} else if direction == "right" {
					fmt.Print(">")
				} else if direction == "down" {
					fmt.Print("v")
				} else if direction == "left" {
					fmt.Print("<")
				}
			} else {
				_, exists := visited[coordinates{x: a, y: b}]

				if exists {
					fmt.Print("X")
				} else if matrix[a][b] == "^" {
					fmt.Print(".")
				} else {
					fmt.Print(matrix[a][b])
				}
			}
		}
		fmt.Println()
	}
}

func day6problem1() string {
	inputString, err := readFile("daysix.txt")
	if err != nil {
		return "Could not read file"
	}

	lines := strings.Split(inputString, "\n")

	inputMatrix := [][]string{}

	visited := make(map[coordinates]bool)

	i := 0
	j := 0

	for index, line := range lines {
		inputMatrix = append(inputMatrix, strings.Split(line, ""))
		initialJ := strings.Index(line, "^")
		if initialJ > -1 {
			i = index
			j = initialJ
			visited[coordinates{x: i, y: j}] = true
		}
	}

	locationCount := 1
	direction := "up"

	for i > 0 && j > 0 && i < len(inputMatrix)-1 && j < len(inputMatrix[i])-1 {
		if direction == "up" {
			if inputMatrix[i-1][j] == "#" {
				direction = "right"
			} else {
				i--
			}
		} else if direction == "right" {
			if inputMatrix[i][j+1] == "#" {
				direction = "down"
			} else {
				j++
			}
		} else if direction == "down" {
			if inputMatrix[i+1][j] == "#" {
				direction = "left"
			} else {
				i++
			}
		} else if direction == "left" {
			if inputMatrix[i][j-1] == "#" {
				direction = "up"
			} else {
				j--
			}
		}

		_, exists := visited[coordinates{x: i, y: j}]

		if !exists {
			locationCount++
			visited[coordinates{x: i, y: j}] = true
		}
	}

	printMatrix(inputMatrix, i, j, direction, visited)

	return strconv.Itoa(locationCount)
}

func createsLoop(inputMatrix [][]string, startingI, startingJ, obstacleI, obstacleJ int) bool {
	direction := "up"
	i := startingI
	j := startingJ

	visited := make(map[directedCoordinates]bool)
	visited[directedCoordinates{x: i, y: j, direction: direction}] = true

	for i > 0 && j > 0 && i < len(inputMatrix)-1 && j < len(inputMatrix[i])-1 {
		if direction == "up" {
			if inputMatrix[i-1][j] == "#" || (i-1 == obstacleI && j == obstacleJ) {
				direction = "right"
			} else {
				i--
			}
		} else if direction == "right" {
			if inputMatrix[i][j+1] == "#" || (i == obstacleI && j+1 == obstacleJ) {
				direction = "down"
			} else {
				j++
			}
		} else if direction == "down" {
			if inputMatrix[i+1][j] == "#" || (i+1 == obstacleI && j == obstacleJ) {
				direction = "left"
			} else {
				i++
			}
		} else if direction == "left" {
			if inputMatrix[i][j-1] == "#" || (i == obstacleI && j-1 == obstacleJ) {
				direction = "up"
			} else {
				j--
			}
		}

		_, exists := visited[directedCoordinates{x: i, y: j, direction: direction}]

		if !exists {
			visited[directedCoordinates{x: i, y: j, direction: direction}] = true
		} else {
			return true
		}
	}

	return false
}

func day6problem2() string {
	inputString, err := readFile("daysix.txt")
	if err != nil {
		return "Could not read file"
	}

	lines := strings.Split(inputString, "\n")

	inputMatrix := [][]string{}

	visited := make(map[coordinates]bool)

	i := 0
	j := 0

	startingPosition := coordinates{x: 0, y: 0}

	for index, line := range lines {
		inputMatrix = append(inputMatrix, strings.Split(line, ""))
		initialJ := strings.Index(line, "^")
		if initialJ > -1 {
			i = index
			j = initialJ
			visited[coordinates{x: i, y: j}] = true
			startingPosition.x = i
			startingPosition.y = j
		}
	}

	direction := "up"

	for i > 0 && j > 0 && i < len(inputMatrix)-1 && j < len(inputMatrix[i])-1 {
		if direction == "up" {
			if inputMatrix[i-1][j] == "#" {
				direction = "right"
			} else {
				i--
			}
		} else if direction == "right" {
			if inputMatrix[i][j+1] == "#" {
				direction = "down"
			} else {
				j++
			}
		} else if direction == "down" {
			if inputMatrix[i+1][j] == "#" {
				direction = "left"
			} else {
				i++
			}
		} else if direction == "left" {
			if inputMatrix[i][j-1] == "#" {
				direction = "up"
			} else {
				j--
			}
		}

		_, exists := visited[coordinates{x: i, y: j}]

		if !exists {
			visited[coordinates{x: i, y: j}] = true
		}
	}

	loopCount := 0

	for coordinate := range visited {
		if createsLoop(inputMatrix, startingPosition.x, startingPosition.y, coordinate.x, coordinate.y) {
			loopCount++
		}
	}

	return strconv.Itoa(loopCount)
}
