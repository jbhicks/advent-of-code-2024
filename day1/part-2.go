package main

import (
	. "github.com/jbhicks/advent-of-code-2024/util"
)

func Part2() int {
	scanner := ReadInputFile("input")
	list1, list2, err := ProcessInput(scanner)
	Check(err)
	score := getSimilarityScore(list1, list2)
	return score
}

func getSimilarityScore(list1 []int, list2 []int) int {
	total := 0

	for _, value := range list1 {
		count := 0
		for _, value2 := range list2 {
			if value2 == value {
				count++
			}
		}
		total += (value * count)
	}

	return total
}
