package timefilter

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestTimeTilter(t *testing.T) {
	dateTime, _ := TimeFilter(map[string][]string{
		"2021-04-04": {"9:00"},
		"2021-04-05": {"13:54"},
		"2021-04-06": {"13:54"},
		"2021-04-07": {"15:54"},
	}, "9:00", "14:00")
	assert.Equal(t, map[string][]string{
		"2021-04-05": {"13:54"},
		"2021-04-06": {"13:54"},
	}, dateTime)
}
