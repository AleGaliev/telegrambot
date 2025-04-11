package checknewdate

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestCheckNewDate(t *testing.T) {
	newDateTimeList := map[string][]string{
		"2025-03-24": {"12:30"},
		"2025-03-25": {"14:30", "14:40", "14:50", "15:00"},
		"2025-03-26": {"14:30", "14:40", "14:50"},
	}

	oldDateTimeList := map[string][]string{
		"2025-03-25": {"14:20", "14:30", "14:40", "14:50"},
		"2025-03-26": {"14:30", "14:50"},
		"2025-03-27": {"14:30", "14:50"},
	}
	changeDataTime := CheckNewDate(newDateTimeList, oldDateTimeList)
	assert.Equal(t, map[string][]string{
		"2025-03-24": {"12:30"},
		"2025-03-25": {"15:00"},
		"2025-03-26": {"14:40"},
	}, changeDataTime)
}
