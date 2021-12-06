package main

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

type dayGroup struct {
	currentDays int
	totalFish   int
}

func main() {
	data, err := ioutil.ReadFile("input")
	if err != nil {
		panic(err)
	}

	fishStrings := strings.Split(string(data), ",")

	dayGroups := []*dayGroup{}

	dayGroupMap := map[int]*dayGroup{}

	for _, s := range fishStrings {
		v, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}

		g, ok := dayGroupMap[v]
		if !ok {
			g = &dayGroup{
				currentDays: v,
				totalFish:   0,
			}
			dayGroupMap[v] = g
			dayGroups = append(dayGroups, g)
		}

		g.totalFish++
	}

	const numDays = 256

	const freshDays = 8
	const oldDays = 6

	totalFish := 0
	for i := 0; i < numDays; i++ {
		dayGroupMap = map[int]*dayGroup{}

		for _, dg := range dayGroups {
			if dg.currentDays > 0 {
				dg.currentDays--

				newGroup, ok := dayGroupMap[dg.currentDays]
				if !ok {
					newGroup = &dayGroup{
						currentDays: dg.currentDays,
						totalFish:   0,
					}
					dayGroupMap[dg.currentDays] = newGroup
				}
				newGroup.totalFish += dg.totalFish
			} else if dg.currentDays == 0 {
				oldFish, ok := dayGroupMap[oldDays]
				if !ok {
					oldFish = &dayGroup{
						currentDays: oldDays,
						totalFish:   0,
					}
					dayGroupMap[oldFish.currentDays] = oldFish
				}

				oldFish.totalFish += dg.totalFish

				freshFish, ok := dayGroupMap[freshDays]
				if !ok {
					freshFish = &dayGroup{
						currentDays: freshDays,
						totalFish:   0,
					}

					dayGroupMap[freshFish.currentDays] = freshFish
				}

				freshFish.totalFish += dg.totalFish
			}
		}

		dayGroups = dayGroups[:0]

		totalFish = 0
		for _, dg := range dayGroupMap {
			dayGroups = append(dayGroups, dg)
			totalFish += dg.totalFish
		}

		fmt.Printf("Day %d: ", i)
		for _, dg := range dayGroups {
			fmt.Printf("%d (%d), ", dg.currentDays, dg.totalFish)
		}
		fmt.Println()
		fmt.Println("Total Fish:", totalFish)
	}
}
