package goadvent2021

import (
	"strconv"
	"strings"
)

type Position struct {
	horizontal int
	depth      int
	aim        int
}

type Command struct {
	direction string
	units     int
}

type Path uint

const (
	Undirected Path = iota
	Directed
)

func NavigateCourse(inputs []string, paths ...Path) int {
	var path Path

	if len(paths) == 0 {
		path = Undirected
	} else {
		path = paths[0]
	}

	course := parseCourse(inputs, path)
	var position Position

	if path == Undirected {
		position = undirectedNavigation(course)
	} else {
		position = directedNavigation(course)
	}
	return position.horizontal * position.depth
}

func directedNavigation(course []Command) Position {
	position := Position{0, 0, 0}
	for _, command := range course {
		switch command.direction {
		case "up":
			position.aim -= command.units
		case "down":
			position.aim += command.units
		default:
			position.horizontal += command.units
			position.depth += position.aim * command.units
		}
	}
	return position
}

func undirectedNavigation(course []Command) Position {
	position := Position{0, 0, 0}
	for _, command := range course {
		switch command.direction {
		case "forward":
			position.horizontal += command.units
		case "down":
			position.depth += command.units
		default:
			position.depth -= command.units
		}
	}
	return position

}

func parseCourse(inputs []string, path Path) []Command {

	var individualCommands []Command

	for _, input := range inputs {
		i := strings.Split(input, " ")
		direction := i[0]
		units, _ := strconv.Atoi(i[1])

		command := Command{
			direction: direction,
			units:     units,
		}

		individualCommands = append(individualCommands, command)

	}
	return individualCommands
}
