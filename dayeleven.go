package main

import (
	"fmt"
	"strconv"
	"strings"
)

func blink(stones []string) []string {
	result := []string{}
	lookup := make(map[string][]string)
	for _, stone := range stones {

		values, exists := lookup[stone]

		if exists {
			for _, value := range values {
				result = append(result, value)
			}
		} else {
			if stone == "0" {
				result = append(result, "1")
			} else if len(stone)%2 == 0 {
				firstStoneInteger, firstStoneIntegerError := strconv.Atoi(stone[:len(stone)/2])

				if firstStoneIntegerError != nil {
					result = append(result, "0")
				}

				secondStoneInteger, secondStoneIntegerError := strconv.Atoi(stone[len(stone)/2:])

				if secondStoneIntegerError != nil {
					result = append(result, "0")
				}

				lookup[stone] = []string{strconv.Itoa(firstStoneInteger), strconv.Itoa(secondStoneInteger)}
				result = append(result, strconv.Itoa(firstStoneInteger))
				result = append(result, strconv.Itoa(secondStoneInteger))
			} else {
				stoneInteger, stoneIntegerError := strconv.Atoi(stone)

				if stoneIntegerError != nil {
					result = append(result, "0")
				}

				lookup[stone] = []string{strconv.Itoa(stoneInteger * 2024)}

				result = append(result, strconv.Itoa(stoneInteger*2024))
			}
		}
	}

	return result
}

func day11problem1() string {
	inputString, err := readFile("dayeleven.txt")
	if err != nil {
		return "Could not read file"
	}

	stones := strings.Split(inputString, " ")
	for i := 0; i < 75; i++ {
		stones = blink(stones)
	}

	return strconv.Itoa(len(stones))
}

var lookup = make(map[string]int)

func blink2(stone int, depth int) int {
	if depth == 0 {
		return 1
	}

	lookupKey := fmt.Sprintf("%d:%d", stone, depth)
	if value, found := lookup[lookupKey]; found {
		return value
	}

	var value int
	if stone == 0 {
		value = blink2(1, depth-1)
	} else {
		s := strconv.Itoa(stone)
		x := len(s)
		if x%2 == 0 {
			left, _ := strconv.Atoi(s[:x/2])
			right, _ := strconv.Atoi(s[x/2:])
			value = blink2(left, depth-1) + blink2(right, depth-1)
		} else {
			value = blink2(stone*2024, depth-1)
		}
	}

	lookup[lookupKey] = value
	return value
}

func day11problem2() string {
	inputString, err := readFile("dayeleven.txt")
	if err != nil {
		return "Could not read file"
	}

	stones := strings.Split(inputString, " ")

	stoneCount := 0

	for _, stone := range stones {
		stoneInt, _ := strconv.Atoi(stone)
		stoneCount += blink2(stoneInt, 75)
	}

	return strconv.Itoa(stoneCount)
}
