package main

import "fmt"

func ExampleGildedRose() {
	var item *Item

	// A normal item
	item = New("normal", 5, 5)
	fmt.Println(item)
	item.Tick()
	fmt.Println(item)

	// Aged Brie
	item = New("Aged Brie", 5, 5)
	fmt.Println(item)
	item.Tick()
	fmt.Println(item)

	// Sulfuras
	item = New("Sulfuras, Hand of Ragnaros", 5, 5)
	fmt.Println(item)
	item.Tick()
	fmt.Println(item)

	// Backstage passes
	item = New("Backstage passes to a TAFKAL80ETC concert", 5, 5)
	fmt.Println(item)
	item.Tick()
	fmt.Println(item)

	// // Conjured
	// item = New("Conjured Mana Cake", 5, 5)
	// fmt.Println(item)
	// item.Tick()
	// fmt.Println(item)

	// Output:
	// normal: 5 days left, quality is 5
	// normal: 4 days left, quality is 4
	// Aged Brie: 5 days left, quality is 5
	// Aged Brie: 4 days left, quality is 6
	// Sulfuras, Hand of Ragnaros: 5 days left, quality is 5
	// Sulfuras, Hand of Ragnaros: 5 days left, quality is 5
	// Backstage passes to a TAFKAL80ETC concert: 5 days left, quality is 5
	// Backstage passes to a TAFKAL80ETC concert: 4 days left, quality is 8
}
