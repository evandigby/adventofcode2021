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

	lastNum := math.MaxInt64
	numIncreasing := 0
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			panic(err)
		}

		if num > lastNum {
			numIncreasing++
		}

		lastNum = num
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println(numIncreasing)
}
