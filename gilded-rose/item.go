package gildedrose

import "fmt"

// Item describes an item sold by the Gilded Rose Inn
type Item struct {
	name    string
	days    int
	quality int
}

// String outputs a string representation of an Item
func (i *Item) String() string {
	return fmt.Sprintf("%s: %d days left, quality is %d", i.name, i.days, i.quality)
}

func (i *Item) normalTick() {
	i.days--
	if i.quality == 0 {
		return
	}
	i.quality--
	if i.days <= 0 {
		i.quality--
	}
}

func (i *Item) brieTick() {
	i.days--
	if i.quality >= 50 {
		return
	}
	if i.days <= 0 {
		i.quality++
	}
	if i.quality < 50 {
		i.quality++
	}
}

func (i *Item) sulfurasTick() {

}

func (i *Item) backstageTick() {
	i.days--
	if i.quality >= 50 {
		return
	}
	if i.days < 0 {
		i.quality = 0
		return
	}
	i.quality++
	if i.days < 10 {
		i.quality++
	}
	if i.days < 5 {
		i.quality++
	}
}

// New creates a new Item
func New(name string, days, quality int) *Item {
	return &Item{
		name:    name,
		days:    days,
		quality: quality,
	}
}

// UpdateQuality ages the item by a day, and updates the quality of the item
func UpdateQuality(items []*Item) {
	for _, item := range items {
		if item.name == "normal" {
			item.normalTick()
			continue
		}
		if item.name == "Aged Brie" {
			item.brieTick()
			continue
		}
		if item.name == "Sulfuras, Hand of Ragnaros" {
			item.sulfurasTick()
			continue
		}
		if item.name == "Backstage passes to a TAFKAL80ETC concert" {
			item.backstageTick()
			continue
		}
	}
}
