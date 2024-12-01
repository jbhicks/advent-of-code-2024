# Advent of Code 2024 with Go

Welcome to the Advent of Code 2024 repository using Golang! This repository contains solutions and resources for the annual programming challenge known as "Advent of Code." Each year, a new set of puzzles is released daily from December 1st to December 25th. Participants are encouraged to solve these puzzles using the Go programming language.

## About Advent of Code

Advent of Code is a fun and challenging way to learn programming. Each day, you are presented with two puzzles: one that can be solved relatively quickly (the "star" part), and another that requires more time and effort (the "star" part). The difficulty increases as the challenge progresses, making it an excellent opportunity for both beginners and experienced programmers to improve their skills.

## Project Structure
This repository is structured to help you manage your solutions efficiently:

```
advent-of-code-2024/
├── day1/
│   ├── part1.go
│   └── part2.go
├── day2/
│   ├── part1.go
│   └── part2.go
├── ...
└── utils/
   ├── __init__.go
   └── helpers.go
```

- **`dayX/`**: Directories for each day, where `X` is the day number (e.g., `day1`, `day2`, etc.).
 - **`part1.go`**: Solution script for the first part of the puzzle.
 - **`part2.go`**: Solution script for the second part of the puzzle.
- **`utils/`**: A directory containing utility scripts and modules that might be useful across multiple days.

## Usage
To run a solution, navigate to the specific day's directory and execute the desired script:

```sh
cd day1
go run part1.go  # For the first part of Day 1
go run part2.go  # For the second part of Day 1
```
