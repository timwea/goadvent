package goadvent2021_test

import (
	"github.com/timwea/goadvent/goadvent2021"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNavigateDirectedCourse(t *testing.T) {
	given := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	expected := 150

	actual := goadvent2021.NavigateCourse(given)
	assert.Equal(t, expected, actual)
}

func TestNavigateUndirectedCourse(t *testing.T) {
	given := []string{
		"forward 5",
		"down 5",
		"forward 8",
		"up 3",
		"down 8",
		"forward 2",
	}

	expected := 900

	actual := goadvent2021.NavigateCourse(given, goadvent2021.Directed)
	assert.Equal(t, expected, actual)
}
