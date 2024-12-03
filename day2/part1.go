package main

import (
	"bufio"
	"fmt"
	"strconv"

	. "github.com/jbhicks/advent-of-code-2024/util"
)

func Part1() int {
	scanner := ReadInputFile("input-example")
	reports := ProcessInput(scanner)

	for i, v := range reports {
		if isLineSafe(v) {
			fmt.Println("Line", i, "is safe")
		} else {
			fmt.Println("Line", i, "is not safe")
		}
	}

	result := 0
	return result
}

func isLineSafe(line string) bool {
	fmt.Println(line)
	// split the line into a slice of strings
	var readings []int

	for _, v := range line {
		if v != 32 {
			num, err := strconv.Atoi(string(v))
			Check(err)
			readings = append(readings, num)
		}
	}
	fmt.Println(readings)

	if readings[0] == readings[1] {
		return false
	}

	if readings[0] < readings[1] {
		for i := 1; i < len(readings)-1; i++ {
			if readings[i] >= readings[i+1] {
				fmt.Println("ASCENDING: ", readings[i], readings[i+1])
				return false
			}
		}
	}

	if readings[0] > readings[1] {
		for i := 1; i < len(readings)-1; i++ {
			if readings[i] <= readings[i+1] {
				fmt.Println("DESCENDING: ", readings[i], readings[i+1])
				return false
			}
		}
	}
	return true
}

func ProcessInput(scanner *bufio.Scanner) []string {
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
