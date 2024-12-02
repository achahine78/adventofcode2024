package main

import (
	"strconv"
	"strings"
)

func isMonotonic(list []int) bool {
	if len(list) <= 2 {
		return true
	}

	direction := list[1] - list[0]

	for i := 1; i < len(list)-1; i++ {
		currDirection := list[i+1] - list[i]
		if currDirection*direction < 0 {
			return false
		}
	}

	return true
}

func isGapSizeSafe(list []int) bool {
	if len(list) < 2 {
		return true
	}

	for i := 0; i < len(list)-1; i++ {
		gap := abs(list[i+1] - list[i])
		if gap < 1 || gap > 3 {
			return false
		}
	}

	return true
}

func convertToInt(s string) int {
	integerRepresentation, err := strconv.Atoi(s)

	if err != nil {
		return 0
	}

	return integerRepresentation
}

func day2problem1() string {
	inputString, err := readFile("daytwo.txt")
	if err != nil {
		return "Could not read file"
	}

	lists := strings.Split(inputString, "\n")

	safeCount := 0

	for _, list := range lists {
		splitList := strings.Split(list, " ")
		parsedList := []int{}

		for _, element := range splitList {
			parsedList = append(parsedList, convertToInt(element))
		}

		if isMonotonic(parsedList) && isGapSizeSafe(parsedList) {
			safeCount++
		}

	}

	return strconv.Itoa(safeCount)

}
