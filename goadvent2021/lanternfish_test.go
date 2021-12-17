package goadvent2021_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/timwea/goadvent/goadvent2021"
	"testing"
)

func TestLanternfishGrownRate(t *testing.T) {

	given := []string{"3,4,3,1,2"}
	numOfDays := 256

	expected := 26984457539

	actual := goadvent2021.LanternfishGrowthRate(given, numOfDays)
	assert.Equal(t, expected, actual)
}
