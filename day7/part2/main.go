package main

import (
	"fmt"
	"io/ioutil"
	"math"
	"strconv"
	"strings"
)

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	crabPositionStrings := strings.Split(string(data), ",")

	crabPositions := []int{}
	min := math.MaxInt64
	max := 0
	for _, s := range crabPositionStrings {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		crabPositions = append(crabPositions, v)

		if v < min {
			min = v
		}
		if v > max {
			max = v
		}
	}

	minFuel := math.MaxInt64

	nums := make([]int, max+1)
	for i := range nums {
		if i == 0 {
			continue
		}
		nums[i] = nums[i-1] + i
	}

	for x := min; x <= max; x++ {
		totalFuel := 0
		for _, v := range crabPositions {
			totalFuel += nums[int(math.Abs(float64(v-x)))]
		}

		if totalFuel < minFuel {
			minFuel = totalFuel
		}
	}

	fmt.Println("Min Fuel:", minFuel)
}
