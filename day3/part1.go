package main

import (
	. "github.com/jbhicks/advent-of-code-2024/util"
)

func Part1() (int, error) {
	// scanner := ReadInputFile("input-example")
	scanner := ReadInputFile("input")
	inputLines := ProcessInput(scanner)
	total := 0

	for _, line := range inputLines {
		lineValue, err := GetLineValue(line)
		if err != nil {
			return -1, err
		}
		total += lineValue
	}
	return total, nil
}
