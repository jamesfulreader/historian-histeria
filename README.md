# Day 1 of Advent of Code 2024
## Historian Histeria
Go implementation for Day 1 of Advent of Code 2024 Historian Histeria [Full Challenge Description](https://adventofcode.com/2024)

## Project Description
My program solves both part 1 and part 2 of the challenge involving location IDs:
1. Calculate the total distance between pairs of location IDs
2. Compute a similarity score between location IDs

## Features
- Efficient parsing of input.txt file containing location ID pairs
- Implement distance calulation algorithm for challenge
- Optimize similarity score computation using hash maps
- Support Large datasets with O(n) time complexity

## Prerequisites
- `slices` package imported Go 1.21 or higher required
- input.txt file can be downloaded from [Advent of Code 2024](https://adventofcode.com/2024/day/1)

## Installation
```bash
# Clone the repository
git clone https://github.com/jamesfulreader/historian-histeria.git
cd historian-histeria

# Build the project
go build
```

## Input.txt file format
```
57643 17620
19062 47340
11105 16109
...
```

## Acknowledgments
- [Advent of Code](https://adventofcode.com/) for the original challenge
