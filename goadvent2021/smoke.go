package goadvent2021

import (
	"strconv"
	"strings"
)

type HeightMap [][]string
type Point int

func RiskLevelSum(inputs []string) int {
	riskLevelSum := 0
	heightMap := parseHeightMap(inputs)

	for i := 0; i < len(heightMap); i++ {
		for j := 0; j < len(heightMap[i]); j++ {
			if heightMap.isLowPoint(i, j) {
				riskLevel, _ := strconv.Atoi(heightMap[i][j])
				riskLevelSum += riskLevel + 1
			}
		}
	}

	return riskLevelSum
}

func (h HeightMap) isLowPoint(i int, j int) bool {
	point, _ := strconv.Atoi(h[i][j])

	if topIsLower(point, i, j, h) {
		return false
	}
	if leftIsLower(point, i, j, h) {
		return false
	}
	if bottomIsLower(point, i, j, h) {
		return false
	}
	if rightIsLower(point, i, j, h) {
		return false
	}

	return true
}

func rightIsLower(point int, i int, j int, heightMap [][]string) bool {
	if j+1 < len(heightMap[i]) {
		right, _ := strconv.Atoi(heightMap[i][j+1])
		return right <= point
	}
	return false
}

func bottomIsLower(point int, i int, j int, heightMap [][]string) bool {
	if i+1 < len(heightMap) {
		bottom, _ := strconv.Atoi(heightMap[i+1][j])
		return bottom <= point
	}
	return false
}

func leftIsLower(point int, i int, j int, heightMap [][]string) bool {
	if j-1 > -1 {
		left, _ := strconv.Atoi(heightMap[i][j-1])
		return left <= point
	}
	return false
}

func topIsLower(point int, i int, j int, heightMap [][]string) bool {
	if i-1 > -1 {
		top, _ := strconv.Atoi(heightMap[i-1][j])
		return top <= point
	}
	return false
}

func parseHeightMap(inputs []string) HeightMap {
	var heightMap HeightMap
	for _, input := range inputs {
		heightMap = append(heightMap, strings.Split(input, ""))
	}
	return heightMap
}
