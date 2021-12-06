package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/fatih/color"
)

type pos struct {
	val   int
	found bool
}

type board struct {
	HasWon bool
	Nums   [][]*pos
}

type coord struct {
	Board *board
	X     int
	Y     int
	pos   *pos
	Line  []*pos
	Col   []*pos
}

func main() {
	file, err := os.Open("input")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)

	boards := []*board{}

	positions := map[int][]*coord{}

	if !scanner.Scan() {
		panic(scanner.Err())
	}

	nums := strings.Split(scanner.Text(), ",")

	for scanner.Scan() {
		if strings.TrimSpace(scanner.Text()) == "" {
			boards = append(boards, &board{
				Nums: [][]*pos{},
			})

			continue
		}

		boardNum := len(boards) - 1
		currentBoard := boards[boardNum]

		spacedLineNums := strings.Split(scanner.Text(), " ")

		currentNums := []*pos{}

		for _, v := range spacedLineNums {
			trimVal := strings.TrimSpace(v)
			if trimVal == "" {
				continue
			}

			num, err := strconv.Atoi(trimVal)
			if err != nil {
				panic(err)
			}

			currentPos := &pos{
				val:   num,
				found: false,
			}

			positions[num] = append(positions[num], &coord{
				Board: currentBoard,
				X:     len(currentNums),
				Y:     len(currentBoard.Nums),
				pos:   currentPos,
			})

			currentNums = append(currentNums, currentPos)
		}

		currentBoard.Nums = append(currentBoard.Nums, currentNums)
	}

	if scanner.Err() != nil {
		panic(scanner.Err())
	}

	for _, coords := range positions {
		for _, coord := range coords {
			for i, line := range coord.Board.Nums {
				if i == coord.Y {
					coord.Line = append(coord.Line, line...)
				}

				coord.Col = append(coord.Col, line[coord.X])
			}
		}
	}

	winners := []*board{}

	var numVal int
	for _, n := range nums {
		numVal, err = strconv.Atoi(n)
		if err != nil {
			panic(err)
		}

		coords, ok := positions[numVal]
		if !ok {
			continue
		}

		for _, coord := range coords {
			coord.pos.found = true

			isWinRow := true
			for _, pos := range coord.Line {
				if !pos.found {
					isWinRow = false
					break
				}
			}

			isWinCol := true
			for _, pos := range coord.Col {
				if !pos.found {
					isWinCol = false
					break
				}
			}

			if isWinRow || isWinCol {
				if coord.Board.HasWon {
					continue
				}

				coord.Board.HasWon = true
				winners = append(winners, coord.Board)
			}
		}

		color.Cyan("Picked Num: %d", numVal)
		for i, b := range boards {
			fmt.Println("Board", i)

			for _, line := range b.Nums {
				for _, n := range line {
					if n.found {
						fmt.Print(color.GreenString("%3d", n.val))
					} else {
						fmt.Print(color.RedString("%3d", n.val))
					}
				}
				fmt.Println("")
			}
		}
		if len(winners) == len(boards) {
			break
		}
	}

	fmt.Println("Last Num:", numVal)

	winner := winners[len(winners)-1]

	sum := 0
	for _, line := range winner.Nums {
		for _, n := range line {
			if n.found {
				continue
			}

			sum += n.val
		}
	}

	fmt.Println("Sum:", sum)
	fmt.Println("Mul:", sum*numVal)

	// 	fmt.Println("Gamma:", gamma)
	// 	fmt.Println("Epsilon:", epsilon)
	// 	fmt.Println("Mul:", gamma*epsilon)
}
