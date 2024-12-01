package main

import "fmt"

func Day2() {
	scanner := ReadInputFile("input")
	list1, list2, err := ProcessInput(scanner)
	fmt.Println(list1, list2)
	Check(err)
	score := getSimilarityScore(list1, list2)
	fmt.Println("Similarity score:", score)
}

func getSimilarityScore(list1 []int, list2 []int) int {
	total := 0

	for _, value := range list1 {
		count := 0
		for _, value2 := range list2 {
			if value2 == value {
				fmt.Println("HIT", value, value2)
				count++
			}
		}
		fmt.Println(value, "occurs", count, "times")
		total += (value * count)
		fmt.Println("Total:", total)
	}

	return total
}
