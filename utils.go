package main

import "os"

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

func abs(a int) int {
	if a < 0 {
		return -a
	}

	return a
}
