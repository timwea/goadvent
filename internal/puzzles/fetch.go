package puzzles

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"time"
)

func Fetch(year int, day int, sessionID string) ([]string, error) {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://adventofcode.com/%d/day/%d/input", year, day), nil)

	session := &http.Cookie{
		Name:   "session",
		Value:  sessionID,
		Domain: "adventofcode.com",
	}
	req.AddCookie(session)

	start := time.Now()

	response, err := http.DefaultClient.Do(req)

	scanner := bufio.NewScanner(response.Body)

	var inputs []string

	for scanner.Scan() {
		input := scanner.Text()
		inputs = append(inputs, input)
	}

	fmt.Println("inputs fetched in", time.Since(start))

	if err != nil {
		fmt.Printf("encountered an error fetching Day %s input: %v", os.Getenv("DAY"), err)
		return nil, err
	}

	return inputs, nil
}
