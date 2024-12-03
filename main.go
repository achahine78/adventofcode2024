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
	}

	function, exists := functions[pn]
	if !exists {
		function = functionDefault
	}

	result := function()
	fmt.Println(result)

	os.Exit(0)
}
