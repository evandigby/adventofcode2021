package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	var bits []int

	for scanner.Scan() {
		if len(bits) == 0 {
			bits = make([]int, len(scanner.Text()))
		}

		fmt.Println(scanner.Text())
		for i, v := range scanner.Text() {
			switch v {
			case '0':
				bits[i] -= 1
			case '1':
				bits[i] += 1
			}
		}
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	epsilon := 0
	gamma := 0

	for _, v := range bits {
		epsilon <<= 1
		gamma <<= 1
		if v > 0 {
			epsilon |= 1
		} else {
			gamma |= 1
		}
	}

	fmt.Println("Gamma:", gamma)
	fmt.Println("Epsilon:", epsilon)
	fmt.Println("Mul:", gamma*epsilon)
}
