package main

import (
	"regexp"
	"strconv"
	"strings"
)

func day3problem1() string {
	inputString, err := readFile("daythree.txt")
	if err != nil {
		return "Could not read file"
	}

	r, _ := regexp.Compile(`mul\(\d+,\d+\)`)

	multiplicationOperations := r.FindAllString(inputString, -1)

	multiplicationSum := 0

	for _, operation := range multiplicationOperations {
		operation = strings.ReplaceAll(operation, "mul(", "")
		operation = strings.ReplaceAll(operation, ")", "")
		operands := strings.Split(operation, ",")
		leftOperand, leftOperandErr := strconv.Atoi(operands[0])

		if leftOperandErr != nil {
			return "Invalid input"
		}

		rightOperand, rightOperandErr := strconv.Atoi(operands[1])

		if rightOperandErr != nil {
			return "Invalid input"
		}

		multiplicationSum += leftOperand * rightOperand
	}

	return strconv.Itoa(multiplicationSum)
}

func day3problem2() string {
	inputString, err := readFile("daythree.txt")
	if err != nil {
		return "Could not read file"
	}

	r, _ := regexp.Compile(`mul\(\d+,\d+\)|don\'t\(\)|do\(\)`)

	multiplicationOperations := r.FindAllString(inputString, -1)

	multiplicationSum := 0

	operationsEnabled := true

	for _, operation := range multiplicationOperations {
		if operation == "do()" {
			operationsEnabled = true
		} else if operation == "don't()" {
			operationsEnabled = false
		} else {
			if operationsEnabled {
				operation = strings.ReplaceAll(operation, "mul(", "")
				operation = strings.ReplaceAll(operation, ")", "")
				operands := strings.Split(operation, ",")
				leftOperand, leftOperandErr := strconv.Atoi(operands[0])

				if leftOperandErr != nil {
					return "Invalid input"
				}

				rightOperand, rightOperandErr := strconv.Atoi(operands[1])

				if rightOperandErr != nil {
					return "Invalid input"
				}

				multiplicationSum += leftOperand * rightOperand
			}
		}
	}

	return strconv.Itoa(multiplicationSum)
}
