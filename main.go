package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	day := flag.String("d", "", "Number of day")
	problem := flag.String("p", "", "Number of problem")

	flag.Parse()

	pn := "d" + string(*day) + "p" + string(*problem)

	functions := map[string]func() string{
		"d1p1": day1problem1,
		"d1p2": day1problem2,
		"d2p1": day2problem1,
		"d2p2": day2problem2,
		"d3p1": day3problem1,
		"d3p2": day3problem2,
		"d4p1": day4problem1,
		"d4p2": day4problem2,
		"d5p1": day5problem1,
		"d5p2": day5problem2,
		"d6p1": day6problem1,
	}

	function, exists := functions[pn]
	if !exists {
		function = functionDefault
	}

	result := function()
	fmt.Println(result)

	os.Exit(0)
}
