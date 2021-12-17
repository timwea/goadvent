package goadvent2021_test

import (
	"github.com/timwea/goadvent/goadvent2021"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReportPowerConsumptionDiagnostic(t *testing.T) {
	given := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	expected := 198

	actual := goadvent2021.ReportDiagnostic(given)
	assert.Equal(t, expected, actual)
}

func TestReportLifeSupportRatingDiagnostic(t *testing.T) {
	given := []string{
		"00100",
		"11110",
		"10110",
		"10111",
		"10101",
		"01111",
		"00111",
		"11100",
		"10000",
		"11001",
		"00010",
		"01010",
	}

	expected := 230

	actual := goadvent2021.ReportDiagnostic(given, goadvent2021.LifeSupportRating)
	assert.Equal(t, expected, actual)
}
