package main

import (
	"strconv"
	"strings"
)

func getMiddleOfCorrectOrdering(list []string, precedenceMap map[string]map[string]bool) int {
	for i := 0; i < len(list)-1; i++ {
		earlier := list[i]
		for j := i + 1; j < len(list); j++ {
			later := list[j]
			_, exists := precedenceMap[earlier][later]

			if !exists {
				return 0
			}
		}
	}

	middleOfList, err := strconv.Atoi(list[len(list)/2])

	if err != nil {
		return 0
	}

	return middleOfList
}

func day5problem1() string {
	inputString, err := readFile("dayfive.txt")
	if err != nil {
		return "Could not read file"
	}

	lines := strings.Split(inputString, "\n")

	precedenceMap := make(map[string]map[string]bool)

	count := 0

	for _, line := range lines {
		if strings.ContainsAny(line, "|") {
			precedence := strings.Split(line, "|")
			earlier := precedence[0]
			later := precedence[1]
			_, exists := precedenceMap[earlier]

			if !exists {
				precedenceMap[earlier] = make(map[string]bool)
			}

			precedenceMap[earlier][later] = true
		} else if strings.ContainsAny(line, ",") {
			list := strings.Split(line, ",")
			count += getMiddleOfCorrectOrdering(list, precedenceMap)
		}
	}

	return strconv.Itoa(count)
}
