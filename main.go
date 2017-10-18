package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	rose "github.com/robphoenix/gilded-rose/gilded-rose"
)

var items = []*rose.Item{
	{"+5 Dexterity Vest", 10, 20},
	{"Aged Brie", 2, 0},
	{"Elixir of the Mongoose", 5, 7},
	{"Sulfuras, Hand of Ragnaros", 0, 80},
	{"Backstage passes to a TAFKAL80ETC concert", 15, 20},
	// {"Conjured Mana Cake", 3, 6},
}

func main() {
	rose.UpdateQuality(items...)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "Item\tDays Left\tQuality\t")
	fmt.Fprintln(w, "----\t---------\t-------\t")

	for _, item := range items {
		fmt.Fprintln(w, item)
	}
}
