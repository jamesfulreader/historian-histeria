package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input, err := os.Open("./input.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer input.Close()

	var leftNums []int
	var rightNums []int

	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		eachLine := strings.Fields(scanner.Text())

		leftNum, err := strconv.Atoi(eachLine[0])
		if err != nil {
			fmt.Println(err)
			continue
		}
		leftNums = append(leftNums, leftNum)

		rightNum, err := strconv.Atoi(eachLine[1])
		if err != nil {
			fmt.Println(err)
			continue
		}
		rightNums = append(rightNums, rightNum)
	}

	leftSlice := leftNums
	rightSlice := rightNums

	slices.Sort(leftSlice)
	slices.Sort(rightSlice)

	if len(leftNums) != len(rightNums) {
		fmt.Println("THE TWO SLICES ARE NOT THE SAME LENGTH!!!")
	}

	distance := 0
	for i := range leftNums {
		distance += int(math.Abs(float64(leftSlice[i] - rightSlice[i])))
	}

	fmt.Printf("The total distance between each of the points is: %d\n", distance)

	// part 2 number of appearances of nums from left compared to nums to the right nums
	// O(n)
	total := 0
	rightMapNums := make(map[int]int)

	for _, rv := range rightNums {
		rightMapNums[rv]++
	}

	for _, v := range leftNums {
		_, exists := rightMapNums[v]
		if exists {
			total += (rightMapNums[v] * (v))
		}
	}
	// o(n^2)
	// for _, lv := range leftNums {
	// 	fmt.Println("working on ", lv)
	// 	for _, rv := range rightNums {
	// 		if lv == rv {
	// 			count += 1
	// 		}
	// 	}
	// 	fmt.Printf("I found %d this many times %d in the rightnums\n", lv, count)
	// 	total += lv * count
	// 	count = 0
	// }
	fmt.Println("total ", total)
}
