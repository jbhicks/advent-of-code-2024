package main

import (
	"strings"

	. "github.com/jbhicks/advent-of-code-2024/util"
)

func Part2() (int, error) {
	total := 0
	scanner := ReadInputFile("input")
	// scanner := ReadInputFile("input-example2")
	inputLines := ProcessInput(scanner)

	for _, line := range inputLines {
		// Split on "don't()", first part will be valid to add
		parts := strings.Split(line, "don't()")

		// Parts before a dont() are valid to add, split on second part
		for i, part := range parts {
			if i%2 == 0 {
				lineValue, err := GetLineValue(part)
				if err != nil {
					return -1, err
				}
				total += lineValue
			}
			if i%2 == 1 { // Second part means it comes after a dont()
				doparts := strings.Split(part, "do()")
				for j, dopart := range doparts {
					if j%2 == 1 { // Second part after 'do()' is valid to add
						lineValue, err := GetLineValue(dopart)
						if err != nil {
							return -1, err
						}
						total += lineValue
					}
				}
			}
		}
	}

	return total, nil
}
