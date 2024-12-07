package main

import (
	"bufio"
	"fmt"
)

func main() {
	part1, err := Part1()
	if err != nil {
		panic(err)
	}
	fmt.Println("Part1: ", part1)
	fmt.Println("Part2: ", Part2())
}

func ProcessInput(scanner *bufio.Scanner) []string {
	lines := []string{}

	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	return lines
}
