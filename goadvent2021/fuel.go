package goadvent2021

import (
	"math"
	"sort"
	"strconv"
	"strings"
)

type BurnRate int

const (
	ConstantBurnRate BurnRate = iota
	VariableBurnRate
)

func FuelCost(inputs []string, burnRates ...BurnRate) int {
	var burnRate BurnRate

	if len(burnRates) == 0 {
		burnRate = ConstantBurnRate
	} else {
		burnRate = burnRates[0]
	}

	positions := parsePositions(inputs)
	fuelCost := computeFuelCost(positions, burnRate)

	return fuelCost
}

func computeFuelCost(positions []int, burnRate BurnRate) int {
	sumMap := make(map[int]int)
	largestNumber := positions[len(positions)-1]
	lowestCost := math.MaxInt64

	for i := 0; i < largestNumber; i++ {
		temp := 0
		for _, p := range positions {
			if burnRate == VariableBurnRate {
				temp += sum(Abs(p-i), sumMap)
			} else {
				temp += Abs(p - i)
			}
		}
		if temp < lowestCost {
			lowestCost = temp
		}
	}
	return lowestCost
}

func parsePositions(inputs []string) []int {
	var positions []int
	temp := strings.Split(inputs[0], ",")
	for _, position := range temp {
		p, _ := strconv.Atoi(position)
		positions = append(positions, p)
	}
	sort.Ints(positions)
	return positions
}

func sum(num int, sumMap map[int]int) int {
	if _, ok := sumMap[num]; ok {
		return sumMap[num]
	}
	s := 0
	for i := 1; i <= num; i++ {
		s += i
	}
	sumMap[num] = s
	return s
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
