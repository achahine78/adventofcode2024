package main

import (
	"strconv"
	"strings"
)

func testCombinations(testCase int, currentValue int, equation []int, index int) bool {
	if currentValue > testCase {
		return false
	}

	if index == len(equation) {
		return testCase == currentValue
	}

	return testCombinations(testCase, currentValue*equation[index], equation, index+1) ||
		testCombinations(testCase, currentValue+equation[index], equation, index+1)
}

func day7problem1() string {
	inputString, err := readFile("dayseven.txt")
	if err != nil {
		return "Could not read file"
	}

	lines := strings.Split(inputString, "\n")

	calibrationCount := 0

	for _, line := range lines {
		splitByColon := strings.Split(line, ":")
		testCase := splitByColon[0]
		equation := strings.Split(strings.TrimLeft(splitByColon[1], " "), " ")

		testCaseInteger, testCaseIntegerErr := strconv.Atoi(testCase)

		if testCaseIntegerErr != nil {
			return "Invalid input"
		}

		equationInteger := make([]int, len(equation))

		for index, operand := range equation {
			operandInteger, operandIntegerErr := strconv.Atoi(operand)

			if operandIntegerErr != nil {
				return "Invalid input"
			}

			equationInteger[index] = operandInteger
		}

		if testCombinations(testCaseInteger, equationInteger[0], equationInteger, 1) {
			calibrationCount += testCaseInteger
		}
	}
	return strconv.Itoa(calibrationCount)
}

func testCombinationsWithConcatenation(testCase int, currentValue int, equation []int, index int) bool {
	if currentValue > testCase {
		return false
	}

	if index == len(equation) {
		return testCase == currentValue
	}

	if index > len(equation) {
		return false
	}

	concatenation := strconv.Itoa(currentValue) + strconv.Itoa(equation[index])

	concatenationInteger, concatenationIntegerErr := strconv.Atoi(concatenation)

	if concatenationIntegerErr != nil {
		return false
	}

	return testCombinationsWithConcatenation(testCase, currentValue*equation[index], equation, index+1) ||
		testCombinationsWithConcatenation(testCase, currentValue+equation[index], equation, index+1) ||
		testCombinationsWithConcatenation(testCase, concatenationInteger, equation, index+1)
}

func day7problem2() string {
	inputString, err := readFile("dayseven.txt")
	if err != nil {
		return "Could not read file"
	}

	lines := strings.Split(inputString, "\n")

	calibrationCount := 0

	for _, line := range lines {
		splitByColon := strings.Split(line, ":")
		testCase := splitByColon[0]
		equation := strings.Split(strings.TrimLeft(splitByColon[1], " "), " ")

		testCaseInteger, testCaseIntegerErr := strconv.Atoi(testCase)

		if testCaseIntegerErr != nil {
			return "Invalid input"
		}

		equationInteger := make([]int, len(equation))

		for index, operand := range equation {
			operandInteger, operandIntegerErr := strconv.Atoi(operand)

			if operandIntegerErr != nil {
				return "Invalid input"
			}

			equationInteger[index] = operandInteger
		}

		if testCombinationsWithConcatenation(testCaseInteger, equationInteger[0], equationInteger, 1) {
			calibrationCount += testCaseInteger
		}
	}
	return strconv.Itoa(calibrationCount)
}
