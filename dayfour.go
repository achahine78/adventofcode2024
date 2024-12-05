package main

import (
	"strconv"
	"strings"
)

func getChristmasCount(inputArray [][]string, i int, j int) int {
	count := 0

	// north
	if i >= 3 {
		if inputArray[i][j]+inputArray[i-1][j]+inputArray[i-2][j]+inputArray[i-3][j] == "XMAS" {
			count++
		}
	}

	// northeast
	if i >= 3 && j < len(inputArray[i])-3 {
		if inputArray[i][j]+inputArray[i-1][j+1]+inputArray[i-2][j+2]+inputArray[i-3][j+3] == "XMAS" {
			count++
		}
	}

	// east
	if j < len(inputArray[i])-3 {
		if inputArray[i][j]+inputArray[i][j+1]+inputArray[i][j+2]+inputArray[i][j+3] == "XMAS" {
			count++
		}
	}

	// southeast
	if i < len(inputArray)-3 && j < len(inputArray[i])-3 {
		if inputArray[i][j]+inputArray[i+1][j+1]+inputArray[i+2][j+2]+inputArray[i+3][j+3] == "XMAS" {
			count++
		}
	}

	// south
	if i < len(inputArray)-3 {
		if inputArray[i][j]+inputArray[i+1][j]+inputArray[i+2][j]+inputArray[i+3][j] == "XMAS" {
			count++
		}
	}

	// southwest
	if i < len(inputArray)-3 && j >= 3 {
		if inputArray[i][j]+inputArray[i+1][j-1]+inputArray[i+2][j-2]+inputArray[i+3][j-3] == "XMAS" {
			count++
		}
	}

	// west
	if j >= 3 {
		if inputArray[i][j]+inputArray[i][j-1]+inputArray[i][j-2]+inputArray[i][j-3] == "XMAS" {
			count++
		}
	}

	// northwest
	if i >= 3 && j >= 3 {
		if inputArray[i][j]+inputArray[i-1][j-1]+inputArray[i-2][j-2]+inputArray[i-3][j-3] == "XMAS" {
			count++
		}
	}

	return count
}

func day4problem1() string {
	inputString, err := readFile("dayfour.txt")
	if err != nil {
		return "Could not read file"
	}

	inputArray := [][]string{}

	lines := strings.Split(inputString, "\n")
	for _, line := range lines {
		inputArray = append(inputArray, strings.Split(line, ""))
	}

	christmasCount := 0
	for i := 0; i < len(inputArray); i++ {
		for j := 0; j < len(inputArray[i]); j++ {
			christmasCount += getChristmasCount(inputArray, i, j)
		}
	}

	return strconv.Itoa(christmasCount)
}

func isMASX(inputArray [][]string, i int, j int) bool {
	if inputArray[i][j] != "A" {
		return false
	}

	if i >= 1 && j >= 1 && i < len(inputArray)-1 && j < len(inputArray[i])-1 {
		if (inputArray[i-1][j-1]+inputArray[i][j]+inputArray[i+1][j+1] == "SAM" || inputArray[i-1][j-1]+inputArray[i][j]+inputArray[i+1][j+1] == "MAS") && (inputArray[i-1][j+1]+inputArray[i][j]+inputArray[i+1][j-1] == "SAM" || inputArray[i-1][j+1]+inputArray[i][j]+inputArray[i+1][j-1] == "MAS") {
			return true
		}
	}

	return false

}

func day4problem2() string {
	inputString, err := readFile("dayfour.txt")
	if err != nil {
		return "Could not read file"
	}

	inputArray := [][]string{}

	lines := strings.Split(inputString, "\n")
	for _, line := range lines {
		inputArray = append(inputArray, strings.Split(line, ""))
	}

	xmasCount := 0
	for i := 0; i < len(inputArray); i++ {
		for j := 0; j < len(inputArray[i]); j++ {
			if isMASX(inputArray, i, j) {
				xmasCount++
			}
		}
	}

	return strconv.Itoa(xmasCount)
}
