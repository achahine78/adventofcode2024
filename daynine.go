package main

import (
	"fmt"
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

func generateClusters(disk []int) [][]int {
	clusteredDisk := [][]int{}
	i := 0
	j := 0

	for j < len(disk) {
		comparison := disk[i]
		cluster := []int{}
		for j < len(disk) && disk[j] == comparison {
			cluster = append(cluster, disk[j])
			j++
		}
		clusteredDisk = append(clusteredDisk, cluster)
		i = j
	}
	return clusteredDisk
}

func uncluster(clusteredDisk [][]int) []int {
	unclusteredDisk := []int{}

	for _, cluster := range clusteredDisk {
		for _, element := range cluster {
			unclusteredDisk = append(unclusteredDisk, element)
		}
	}

	return unclusteredDisk
}

func printClusteredDisk(clusteredDisk [][]int) {
	for _, file := range uncluster(clusteredDisk) {
		if file == -1 {
			fmt.Print(".")
		} else {
			fmt.Print(file)
		}
	}
	fmt.Println()
}

func day9problem2() string {
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

	clusteredDisk := generateClusters(disk)

	i := 0
	j := len(clusteredDisk) - 1

	for j >= 0 {
		if clusteredDisk[j][0] == -1 {
			j--
		} else {
			for i < j && (len(clusteredDisk[i]) < len(clusteredDisk[j]) || clusteredDisk[i][0] != -1) {
				i++
			}

			if len(clusteredDisk[i]) >= len(clusteredDisk[j]) && clusteredDisk[j][0] != -1 {
				for index := range clusteredDisk[j] {
					clusteredDisk[j][index], clusteredDisk[i][index] = clusteredDisk[i][index], clusteredDisk[j][index]
				}
				clusteredDisk = generateClusters(uncluster(clusteredDisk))
			}
			j--
		}
		i = 0
		// printClusteredDisk(clusteredDisk)
	}

	count := 0

	for mulitiplier, diskValue := range uncluster(clusteredDisk) {
		if diskValue > -1 {
			count += mulitiplier * diskValue
		}
	}

	return strconv.Itoa(count)
}
