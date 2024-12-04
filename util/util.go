package util

import (
	"bufio"
	"os"
	"slices"
	"strconv"
	"strings"
)

// Reads a file and returns 2 sorted lists of integers
func ProcessInput(scanner *bufio.Scanner) ([]int, []int, error) {
	list1 := make([]int, 0)
	list2 := make([]int, 0)

	// Loop over the lines do stuff
	for scanner.Scan() {
		line := scanner.Text()
		nums := strings.Split(line, "   ")

		for i, numStr := range nums {
			num, err := strconv.Atoi(numStr)
			if err != nil {
				return nil, nil, err
			}
			if i == 0 {
				list1 = append(list1, num)
			} else {
				list2 = append(list2, num)
			}
		}
	}
	// Check if there was an error during scanning (e.g., file not found, read errors)
	if err := scanner.Err(); err != nil {
		return nil, nil, err // Return the scan error up the call stack
	} else {
		slices.Sort(list1)
		slices.Sort(list2)
	}
	return list1, list2, nil
}

// Setup reading a file by creating a scanner
func ReadInputFile(fileName string) *bufio.Scanner {
	bytes, err := os.ReadFile(fileName)
	Check(err)
	input := string(bytes)
	scanner := bufio.NewScanner(strings.NewReader(input))
	return scanner
}

func Check(e error) {
	if e != nil {
		panic(e)
	}
}
