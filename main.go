package main

import "github.com/robphoenix/gilded-rose/item"

var items = []struct {
	name    string
	days    int
	quality int
}{
	{"+5 Dexterity Vest", 10, 20},
	{"Aged Brie", 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	{"Conjured Mana Cake", 3, 6},
}

func main() {
	for _, v := range items {
		item := item.New(v.name, v.days, v.quality)
		item.Tick()
	}
}
