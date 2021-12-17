package goadvent2021_test

import (
	"github.com/timwea/goadvent/goadvent2021"
	"testing"

	"github.com/stretchr/testify/assert"
)

func FuelCostTest(t *testing.T) {
	given := []string{
		"16,1,2,0,4,2,7,1,2,14",
	}

	expected := 168

	actual := goadvent2021.FuelCost(given, goadvent2021.VariableBurnRAte)
	assert.Equal(t, expected, actual)
}
