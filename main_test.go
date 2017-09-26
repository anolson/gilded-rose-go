package main

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
	{"Aged Brie", "before sell date with max quality", 5, 50, 50},
	{"Aged Brie", "on sell date", 0, 10, 12},
	{"Aged Brie", "on sell date near max quality", 0, 49, 50},
	{"Aged Brie", "on sell date with max quality", 0, 50, 50},
	{"Aged Brie", "after sell date", -10, 10, 12},
	{"Aged Brie", "after sell date with max quality", -10, 50, 50},
	{"Sulfuras, Hand of Ragnaros", "before sell date", 5, 80, 80},
	{"Sulfuras, Hand of Ragnaros", "on sell date", 0, 80, 80},
	{"Sulfuras, Hand of Ragnaros", "after sell date", -10, 80, 80},
	{"Backstage passes to a TAFKAL80ETC concert", "long before sell date", 11, 10, 11},
	{"Backstage passes to a TAFKAL80ETC concert", "long before sell date at max quality", 11, 50, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "medium close to sell date (upper bound)", 10, 10, 12},
	{"Backstage passes to a TAFKAL80ETC concert", "medium close to sell date (upper bound) at max quality", 10, 50, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "medium close to sell date (lower bound)", 6, 10, 12},
	{"Backstage passes to a TAFKAL80ETC concert", "medium close to sell date (lower bound) at max quality", 6, 50, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "very close to sell date (upper bound)", 5, 10, 13},
	{"Backstage passes to a TAFKAL80ETC concert", "very close to sell date (upper bound) at max quality", 5, 50, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "very close to sell date (lower bound)", 1, 10, 13},
	{"Backstage passes to a TAFKAL80ETC concert", "very close to sell date (lower bound) at max quality", 1, 50, 50},
	{"Backstage passes to a TAFKAL80ETC concert", "on sell date", 0, 10, 0},
	{"Backstage passes to a TAFKAL80ETC concert", "after sell date", -10, 10, 0},
	// {"Conjured Mana Cake", "before sell date", 5, 10, 8},
	// {"Conjured Mana Cake", "before sell date at zero quality", 5, 0, 0},
	// {"Conjured Mana Cake", "on sell date", 0, 10, 6},
	// {"Conjured Mana Cake", "on sell date at zero quality", 0, 0, 0},
	// {"Conjured Mana Cake", "after sell date", -10, 10, 6},
	// {"Conjured Mana Cake", "after sell date at zero quality", -10, 0, 0},
}

func TestGildedRose(t *testing.T) {
	for _, tt := range tests {
		item := New(tt.name, tt.days, tt.quality)
		item.Tick()
		if item.quality != tt.expected {
			t.Fatalf("\n%s %s\nexpected quality: %d\nactual quality: %d", tt.name, tt.description, tt.quality, item.quality)
		}
		if item.name == "Sulfuras, Hand of Ragnaros" {
			if item.daysLeft != tt.days {
				t.Fatalf("\n%s %s\nexpected days left: %d\nactual days left: %d", tt.name, tt.description, tt.days, item.daysLeft)
			}
			continue
		}
		if item.daysLeft != tt.days-1 {
			t.Fatalf("\n%s %s\nexpected days left: %d\nactual days left: %d", tt.name, tt.description, tt.days, item.daysLeft)
		}
	}

}
