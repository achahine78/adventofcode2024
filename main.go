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
		"d1p1":  day1problem1,
		"d1p2":  day1problem2,
		"d2p1":  day2problem1,
		"d2p2":  day2problem2,
		"d3p1":  day3problem1,
		"d3p2":  day3problem2,
		"d4p1":  day4problem1,
		"d4p2":  day4problem2,
		"d5p1":  day5problem1,
		"d5p2":  day5problem2,
		"d6p1":  day6problem1,
		"d6p2":  day6problem2,
		"d7p1":  day7problem1,
		"d7p2":  day7problem2,
		"d8p1":  day8problem1,
		"d8p2":  day8problem2,
		"d9p1":  day9problem1,
		"d9p2":  day9problem2,
		"d10p1": day10problem1,
		"d10p2": day10problem2,
		"d11p1": day11problem1,
		"d11p2": day11problem2,
		"d12p1": day12problem1,
		"d12p2": day12problem2,
	}

	function, exists := functions[pn]
	if !exists {
		function = functionDefault
	}

	result := function()
	fmt.Println(result)

	os.Exit(0)
}
