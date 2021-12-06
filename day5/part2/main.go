package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type vent struct {
	Start coord
	End   coord
}

type coord struct {
	X int
	Y int
}

func getVentCoord(coordVals []string) coord {
	x, err := strconv.Atoi(coordVals[0])
	if err != nil {
		panic(err)
	}

	y, err := strconv.Atoi(coordVals[1])
	if err != nil {
		panic(err)
	}

	return coord{X: x, Y: y}
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	vents := []*vent{}

	maxX := 0
	maxY := 0
	for scanner.Scan() {
		coords := strings.Split(scanner.Text(), " -> ")

		coord1 := strings.Split(coords[0], ",")
		coord2 := strings.Split(coords[1], ",")

		vent := vent{
			Start: getVentCoord(coord1),
			End:   getVentCoord(coord2),
		}

		if vent.Start.X > maxX {
			maxX = vent.Start.X
		}

		if vent.End.X > maxX {
			maxX = vent.End.X
		}

		if vent.Start.Y > maxY {
			maxY = vent.Start.Y
		}

		if vent.End.Y > maxY {
			maxY = vent.End.Y
		}

		vents = append(vents, &vent)
	}

	fmt.Println("maxX", maxX)
	fmt.Println("maxY", maxY)

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	board := [][]int{}

	for x := 0; x <= maxX; x++ {
		board = append(board, make([]int, maxY+1))
	}

	for _, vent := range vents {
		xInc := 1
		yInc := 1

		if vent.Start.X > vent.End.X {
			xInc = -1
		} else if vent.Start.X < vent.End.X {
			xInc = 1
		} else {
			xInc = 0
		}

		if vent.Start.Y > vent.End.Y {
			yInc = -1
		} else if vent.Start.Y < vent.End.Y {
			yInc = 1
		} else {
			yInc = 0
		}
		for x, y := vent.Start.X, vent.Start.Y; x != (vent.End.X+xInc) || y != (vent.End.Y+yInc); {
			board[x][y]++
			x += xInc
			y += yInc
		}

	}

	moreThanTwo := 0
	for _, line := range board {
		for _, place := range line {
			if place >= 2 {
				moreThanTwo++
			}
		}
	}

	fmt.Println("More Than Two:", moreThanTwo)
}
