package item

import (
	"testing"
)

var tests = []struct {
	name        string
	description string
	days        int
	quality     int
	expected    int
}{
	{"normal", "before sell date", 5, 10, 9},
	{"normal", "on sell date", 0, 10, 8},
	{"normal", "after sell date", -10, 10, 8},
	{"normal", "of zero quality", 5, 0, 0},
}

func TestGildedRose(t *testing.T) {
	for _, tt := range tests {
		item := New(tt.name, tt.days, tt.quality)
		item.Tick()
		if item.quality != tt.expected {
			t.Fatalf("%s, expected: %d, actual: %d", tt.description, tt.quality, item.quality)
		}
	}

}
