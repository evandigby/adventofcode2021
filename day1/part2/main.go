package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	lastSum := math.MaxInt64
	numIncreasing := 0

	nums := make([]int, 0, 3)

	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		nums = append(nums, num)
		if len(nums) < 3 {
			continue
		}

		sum := 0
		for _, v := range nums {
			sum += v
		}

		if sum > lastSum {
			numIncreasing++
		}

		lastSum = sum

		nums = nums[1:]
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(numIncreasing)
}
