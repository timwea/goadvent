package goadvent2021_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/timwea/goadvent/goadvent2021"
	"testing"
)

func TestRiskLevelSum(t *testing.T) {
	given := []string{
		"2199943210",
		"3987894921",
		"9856789892",
		"8767896789",
		"9899965678",
	}
	
	expected := 15

	actual := goadvent2021.RiskLevelSum(given)
	assert.Equal(t, expected, actual)
}
