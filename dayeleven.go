package main

import (
	"strconv"
	"strings"
)

func blink(stones []string) []string {
	result := []string{}
	lookup := make(map[string][]string)
	lookup["0"] = []string{"1"}
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
	for i := 0; i < 25; i++ {
		stones = blink(stones)
	}

	return strconv.Itoa(len(stones))
}
