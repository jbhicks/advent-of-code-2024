package main

import (
	"errors"
	"fmt"
)

func Day1() {
	scanner := ReadInputFile("input")
	list1, list2, err := ProcessInput(scanner)
	Check(err)
	totalDistance, err := findTotalDistance(list1, list2)
	Check(err)
	fmt.Println("Total distance:", totalDistance)
}

func findTotalDistance(list1 []int, list2 []int) (int, error) {
	if len(list1) != len(list2) {
		return 0, errors.New("Lists must be the same length")
	}
	distances := make([]int, len(list1))

	for i := 0; i < len(list1); i++ {
		if list1[i] < list2[i] {
			distances[i] = list2[i] - list1[i]
		} else {
			distances[i] = list1[i] - list2[i]
		}
	}

	totalDistance := 0
	for _, distance := range distances {
		totalDistance += distance
	}
	return totalDistance, nil
}
