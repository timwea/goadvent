package goadvent2021

import (
	"strconv"
	"strings"
)

type LanternfishSchool struct {
	school map[int]int
}

func (ls *LanternfishSchool) processDay() {
	temp := LanternfishSchool{
		school: make(map[int]int),
	}
	for k, v := range ls.school {
		newKey := k - 1
		if newKey < 0 {
			temp.updateSchool(6, v)
			temp.addNewSpawn(v)
		} else {
			temp.updateSchool(newKey, v)
		}
	}
	ls.school = temp.school
}

func (ls *LanternfishSchool) updateSchool(newKey int, v int) {
	if _, ok := ls.school[newKey]; ok {
		ls.school[newKey] = ls.school[newKey] + v
	} else {
		ls.school[newKey] = v
	}
}

func (ls *LanternfishSchool) addNewSpawn(v int) {
	ls.school[8] = v
}

func LanternfishGrowthRate(inputs []string, numOfDays int) int {
	lanternfishSchool := LanternfishSchool{
		school: parseInput(strings.Split(inputs[0], ",")),
	}
	growthRate := computeGrowthRate(lanternfishSchool, numOfDays)
	return growthRate
}

func computeGrowthRate(lanternfishSchool LanternfishSchool, numOfDays int) int {
	for i := 0; i < numOfDays; i++ {
		lanternfishSchool.processDay()
	}
	var lanternfishTotal int
	for _, v := range lanternfishSchool.school {
		lanternfishTotal += v
	}
	return lanternfishTotal
}

func parseInput(inputs []string) map[int]int {
	lanternfishes := make(map[int]int)

	for i := 0; i < len(inputs); i++ {
		days, _ := strconv.Atoi(inputs[i])
		if _, ok := lanternfishes[days]; ok {
			lanternfishes[days]++
		} else {
			lanternfishes[days] = 1
		}
	}
	return lanternfishes
}
