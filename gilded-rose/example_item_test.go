// +build example

package gildedrose

import (
	"fmt"
	"os"
	"text/tabwriter"
)

func ExampleGildedRose() {
	var items []*Item

	items = append(items, &Item{"normal", 5, 5})
	items = append(items, &Item{"Aged Brie", 5, 5})
	items = append(items, &Item{"Sulfuras, Hand of Ragnaros", 5, 5})
	items = append(items, &Item{"Backstage passes to a TAFKAL80ETC concert", 5, 5})
	// when you're ready to uncomment the line for the conjured item
	// you will also have to update the expected output below
	//
	// items = append(items, &Item{"Conjured Mana Cake", 5, 5})

	UpdateQuality(items...)

	w := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	defer w.Flush()

	fmt.Fprintln(w, "Item\tDays Left\tQuality\t")
	fmt.Fprintln(w, "----\t---------\t-------\t")

	for _, item := range items {
		fmt.Fprintln(w, item)
	}

	// Output:
	// Item                                       Days Left  Quality
	// ----                                       ---------  -------
	// normal                                     4          4
	// Aged Brie                                  4          6
	// Sulfuras, Hand of Ragnaros                 5          5
	// Backstage passes to a TAFKAL80ETC concert  4          8
}
