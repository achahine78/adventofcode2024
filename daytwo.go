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

func getMonotonicityFailureIndices(list []int) []int {
	result := []int{}
	if len(list) <= 2 {
		return result
	}

	direction := list[1] - list[0]

	for i := 1; i < len(list)-1; i++ {
		currDirection := list[i+1] - list[i]
		if currDirection*direction < 0 {
			result = append(result, i)
			result = append(result, i+1)
		}
	}

	return result
}

func getGapSizeFailureIndices(list []int) []int {
	result := []int{}
	if len(list) < 2 {
		return result
	}

	for i := 0; i < len(list)-1; i++ {
		gap := abs(list[i+1] - list[i])
		if gap < 1 || gap > 3 {
			result = append(result, i)
			result = append(result, i+1)
		}
	}

	return result
}

func canBeMadeSafeWithFailureIndexRemoved(parsedList, failureIndices []int) bool {
	for _, failureIndex := range failureIndices {
		listWithFailureIndexRemoved := remove(parsedList, failureIndex)

		if isMonotonic(listWithFailureIndexRemoved) && isGapSizeSafe(listWithFailureIndexRemoved) {
			return true
		}
	}

	return false
}

func day2problem2() string {
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

		monotonicityFailureIndices := getMonotonicityFailureIndices(parsedList)
		gapSizeFailureIndices := getGapSizeFailureIndices(parsedList)

		failureIndices := deduplicate(append(monotonicityFailureIndices, gapSizeFailureIndices...))
		failureIndices = append(failureIndices, 0)

		if len(monotonicityFailureIndices) == 0 && len(gapSizeFailureIndices) == 0 {
			safeCount++
		} else if canBeMadeSafeWithFailureIndexRemoved(parsedList, failureIndices) {
			safeCount++
		}
	}

	return strconv.Itoa(safeCount)

}
