package main

import "fmt"

// Item describes an item sold by the Gilded Rose Inn
type Item struct {
	name    string
	days    int
	quality int
}

// New creates a new Item
func New(name string, days, quality int) *Item {
	return &Item{
		name:    name,
		days:    days,
		quality: quality,
	}
}

// String implements the Stringer interface
func (i *Item) String() string {
	return fmt.Sprintf("%s: %d days left, quality is %d", i.name, i.days, i.quality)
}

// Tick ages the item by a day, and updates the quality of the item
func (i *Item) Tick() {
	if i.name != "Aged Brie" && i.name != "Backstage passes to a TAFKAL80ETC concert" {
		if i.quality > 0 {
			if i.name != "Sulfuras, Hand of Ragnaros" {
				i.quality = i.quality - 1
			}
		}
	} else {
		if i.quality < 50 {
			i.quality = i.quality + 1
			if i.name == "Backstage passes to a TAFKAL80ETC concert" {
				if i.days < 11 {
					if i.quality < 50 {
						i.quality = i.quality + 1
					}
				}
				if i.days < 6 {
					if i.quality < 50 {
						i.quality = i.quality + 1
					}
				}
			}
		}
	}

	if i.name != "Sulfuras, Hand of Ragnaros" {
		i.days = i.days - 1
	}

	if i.days < 0 {
		if i.name != "Aged Brie" {
			if i.name != "Backstage passes to a TAFKAL80ETC concert" {
				if i.quality > 0 {
					if i.name != "Sulfuras, Hand of Ragnaros" {
						i.quality = i.quality - 1
					}
				}
			} else {
				i.quality = i.quality - i.quality
			}
		} else {
			if i.quality < 50 {
				i.quality = i.quality + 1
			}
		}
	}
}
