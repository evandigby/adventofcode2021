package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	horizontalPosition := 0
	depth := 0

	for scanner.Scan() {
		sp := strings.Split(scanner.Text(), " ")

		num, err := strconv.Atoi(sp[1])
		if err != nil {
			panic(err)
		}

		switch sp[0] {
		case "forward":
			horizontalPosition += num
		case "down":
			depth += num
		case "up":
			depth -= num
		}

	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	fmt.Println("Horizontal:", horizontalPosition)
	fmt.Println("Depth:", depth)
	fmt.Println("Mul:", horizontalPosition*depth)
}
