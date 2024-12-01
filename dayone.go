package main

import (
	"slices"
	"strconv"
	"strings"
)

func day1problem1() string {
	inputString, err := readFile("dayone.txt")
	if err != nil {
		return "Could not read file"
	}

	listPairs := strings.Split(inputString, "\n")
	listOne := []string{}
	listTwo := []string{}
	for _, pair := range listPairs {
		splitPair := strings.Split(pair, "   ")
		listOne = append(listOne, splitPair[0])
		listTwo = append(listTwo, splitPair[1])
	}

	slices.Sort(listOne)
	slices.Sort(listTwo)

	sumOfDifferences := 0

	for i := 0; i < len(listOne); i++ {
		rightInt, rightIntErr := strconv.Atoi(listTwo[i])

		if rightIntErr != nil {
			return "Input improperly formatted"
		}

		leftInt, leftIntErr := strconv.Atoi(listOne[i])

		if leftIntErr != nil {
			return "Input improperly formatted"
		}

		sumOfDifferences += abs(leftInt - rightInt)
	}
	return strconv.Itoa(sumOfDifferences)
}

func day1problem2() string {
	inputString, err := readFile("dayone.txt")
	if err != nil {
		return "Could not read file"
	}

	listPairs := strings.Split(inputString, "\n")
	listOne := []string{}
	listTwo := []string{}
	multiplierMap := make(map[string]int)
	for _, pair := range listPairs {
		splitPair := strings.Split(pair, "   ")
		listOne = append(listOne, splitPair[0])
		listTwo = append(listTwo, splitPair[1])

		multiplierMap[splitPair[1]]++
	}

	sumOfSimilarityScores := 0

	for _, number := range listOne {
		numberValue, err := strconv.Atoi(number)

		if err != nil {
			return "Input improperly formatted"
		}

		sumOfSimilarityScores += numberValue * multiplierMap[number]
	}

	return strconv.Itoa(sumOfSimilarityScores)

}
