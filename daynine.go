package main

import (
	"strconv"
)

func day9problem1() string {
	inputString, err := readFile("daynine.txt")
	if err != nil {
		return "Could not read file"
	}

	disk := []int{}

	id := 0

	for index, character := range inputString {
		numberOfInsertions := int(character - '0')
		if index%2 == 0 { // files
			for i := 0; i < numberOfInsertions; i++ {
				disk = append(disk, id)
			}
			id++
		} else { // freespace
			for i := 0; i < numberOfInsertions; i++ {
				disk = append(disk, -1)
			}
		}
	}

	i := 0
	j := len(disk) - 1

	for i < j {
		if disk[i] != -1 {
			i++
		}

		if disk[j] == -1 {
			j--
		}

		if disk[i] == -1 && disk[j] != -1 {
			disk[i], disk[j] = disk[j], disk[i]
		}
	}

	count := 0

	for mulitiplier, diskValue := range disk {
		if diskValue > -1 {
			count += mulitiplier * diskValue
		}
	}

	return strconv.Itoa(count)
}
