package main

import (
	"regexp"
	"strconv"

	. "github.com/jbhicks/advent-of-code-2024/util"
)

func Part1() (int, error) {
	// scanner := ReadInputFile("input-example")
	scanner := ReadInputFile("input")
	inputLines := ProcessInput(scanner)
	total := 0

	for _, line := range inputLines {
		// Compile the regular expression pattern
		re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)

		// Find all matches
		matches := re.FindAllString(line, -1)

		// Add all the matches' multiplied values
		for _, match := range matches {
			mult1, mult2, err := GetMultiplicands(match)
			if err != nil {
				return -1, err
			}
			total += mult1 * mult2
		}

	}
	return total, nil
}

func GetMultiplicands(match string) (int, int, error) {
	re := regexp.MustCompile(`\d{1,3}`)
	matches := re.FindAllString(match, -1)
	mult1, err := strconv.Atoi(matches[0])
	if err != nil {
		return -1, -1, err
	}
	mult2, err := strconv.Atoi(matches[1])
	if err != nil {
		return -1, -1, err
	}
	return mult1, mult2, nil
}
