package main

import (
	"os"
)

func readFile(filename string) (string, error) {
	data, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}

	return string(data), nil
}

func functionDefault() string {
	return "Please select a problem name to run"
}

func remove(slice []int, s int) []int {
	newSlice := make([]int, 0, len(slice)-1)
	newSlice = append(newSlice, slice...)
	return append(newSlice[:s], newSlice[s+1:]...)
}

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}

func deduplicate(slice []int) []int {
	lookup := make(map[int]bool)

	deduplicatedList := []int{}

	for _, element := range slice {
		_, exists := lookup[element]

		if !exists {
			deduplicatedList = append(deduplicatedList, element)
		}

		lookup[element] = true
	}

	return deduplicatedList
}

type coordinates struct {
	x int
	y int
}

type directedCoordinates struct {
	x         int
	y         int
	direction string
}
