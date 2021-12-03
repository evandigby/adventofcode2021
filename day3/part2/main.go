package main

import (
	"bufio"
	"fmt"
	"os"
)

func narrow(vals []string, most bool) string {
	strLen := len(vals[0])
	for currentPosition := 0; currentPosition < strLen; currentPosition++ {
		bitCount := 0

		for _, num := range vals {
			switch num[currentPosition] {
			case '0':
				bitCount -= 1
			case '1':
				bitCount += 1
			}
		}

		newVals := make([]string, 0, len(vals))

		for _, v := range vals {
			current := byte('0')
			if bitCount >= 0 {
				current = byte('1')
			}

			if most {
				if v[currentPosition] == current {
					newVals = append(newVals, v)
				}
			} else {
				if v[currentPosition] != current {
					newVals = append(newVals, v)
				}
			}
		}

		fmt.Println("newVals:", newVals)
		if len(newVals) == 1 {
			return newVals[0]
		}

		vals = newVals
	}

	return vals[0]
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	nums := []string{}

	for scanner.Scan() {
		nums = append(nums, scanner.Text())
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	oxyVal := narrow(nums, true)
	c02Val := narrow(nums, false)

	c02 := 0
	oxy := 0
	for _, v := range oxyVal {
		oxy <<= 1
		switch v {
		case '1':
			oxy |= 1
		}
	}

	for _, v := range c02Val {
		c02 <<= 1
		switch v {
		case '1':
			c02 |= 1
		}
	}

	fmt.Println("c02:", c02)
	fmt.Println("oxy:", oxy)
	fmt.Println("Mul:", c02*oxy)
}
