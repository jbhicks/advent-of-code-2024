package main

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	part1, err := Part1()
	if err != nil {
		panic(err)
	}

	part2, err := Part2()
	if err != nil {
		panic(err)
	}

	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", part2)
}

func GetLineValue(line string) (int, error) {
	total := 0

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

func ProcessInput(scanner *bufio.Scanner) []string {
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
