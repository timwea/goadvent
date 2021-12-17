package main

import (
	"fmt"
	"github.com/timwea/goadvent/goadvent2021"
	"github.com/timwea/goadvent/internal/puzzles"
	"os"
	"time"
)

func main() {
	year := 2021
	day := 9

	inputs, err := puzzles.Fetch(year, day, os.Getenv("SESSION_ID"))

	if err != nil {
		fmt.Printf("error fetching Day %d input: %v", day, err)
		os.Exit(1)
	}

	start := time.Now()

	answer := goadvent2021.RiskLevelSum(inputs)
	fmt.Println(answer)

	fmt.Println("puzzle solved in", time.Since(start))

}
