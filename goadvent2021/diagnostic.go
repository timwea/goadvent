package goadvent2021

import (
	"fmt"
	"strconv"
)

type Power struct {
	gamma   string
	epsilon string
}

type LifeSupport struct {
	oxygen string
	co2    string
}

type DataPoint string
type Diagnostic int

const (
	PowerConsumption Diagnostic = iota
	LifeSupportRating
)

var BinarySize int

func ReportDiagnostic(dataPoints []string, diagnostics ...Diagnostic) int {
	var diagnostic Diagnostic

	if len(diagnostics) == 0 {
		diagnostic = PowerConsumption
	} else {
		diagnostic = diagnostics[0]
	}

	var report int

	if diagnostic == PowerConsumption {
		report = reportPowerConsumption(dataPoints)
	} else {
		report = reportLifeSupportRating(dataPoints)
	}
	return report
}

func reportPowerConsumption(dataPoints []string) int {
	dataPointsMap := convertToMap(dataPoints)
	power := Power{"", ""}

	for position := 0; position < BinarySize; position++ {
		var mostCommonBit string
		var leastCommonBit string
		for range dataPoints {
			mostCommonBit, leastCommonBit = getCommonBits(position, dataPointsMap)
		}
		power.gamma += mostCommonBit
		power.epsilon += leastCommonBit
		fmt.Println(power)
	}

	powerConsumption := computeDiagnostic(power.gamma, power.epsilon)
	return powerConsumption
}

func reportLifeSupportRating(dataPoints []string) int {
	dataPointsMap1 := convertToMap(dataPoints)
	dataPointsMap2 := convertToMap(dataPoints)
	lifeSupport := LifeSupport{"", ""}
	lifeSupport.oxygen = getOxygenGenRating(dataPointsMap1)
	lifeSupport.co2 = getC02Rating(dataPointsMap2)
	lifeSupportRating := computeDiagnostic(lifeSupport.co2, lifeSupport.oxygen)
	return lifeSupportRating

}

func getC02Rating(dataPoints map[string]string) string {
	var co2Rating string
out:
	for position := 0; position < BinarySize; position++ {
		mostCommonBit, _ := getCommonBits(position, dataPoints)
		for dataPoint := range dataPoints {
			if len(dataPoints) == 1 {
				co2Rating = dataPoint
				break out
			}
			if string(dataPoint[position]) == mostCommonBit || (string(dataPoint[position]) == "1" && mostCommonBit == "E") {
				delete(dataPoints, dataPoint)
			}
		}
	}
	return co2Rating
}

func getOxygenGenRating(dataPoints map[string]string) string {
	var oxygenGenRating string
out:
	for position := 0; position < BinarySize; position++ {
		_, leastCommonBit := getCommonBits(position, dataPoints)
		for dataPoint := range dataPoints {
			if len(dataPoints) == 1 {
				oxygenGenRating = dataPoint
				break out
			}
			if string(dataPoint[position]) == leastCommonBit || (string(dataPoint[position]) == "0" && leastCommonBit == "E") {
				delete(dataPoints, dataPoint)
			}
		}
	}

	return oxygenGenRating
}

func getCommonBits(position int, dataPoints map[string]string) (string, string) {
	numOfZeroes := 0
	numOfOnes := 0

	for dataPoint := range dataPoints {
		if string(dataPoint[position]) == "0" {
			numOfZeroes++
		} else {
			numOfOnes++
		}
	}

	if numOfZeroes > numOfOnes {
		return "0", "1"
	} else if numOfOnes > numOfZeroes {
		return "1", "0"
	}

	return "E", "E"
}

func convertToMap(dataPoints []string) map[string]string {
	dataPointsMap := make(map[string]string)
	for _, dataPoint := range dataPoints {
		dataPointsMap[dataPoint] = ""
	}
	return dataPointsMap
}

func computeDiagnostic(v1 string, v2 string) int {
	return convertToInt(v1) * convertToInt(v2)
}

func convertToInt(v string) int {
	vInt, err := strconv.ParseInt(v, 2, 64)
	if err != nil {
		fmt.Println(err)
	}
	return int(vInt)
}
