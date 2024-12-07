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
