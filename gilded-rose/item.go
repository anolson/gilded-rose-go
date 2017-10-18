package gildedrose

import "fmt"

// Updater is the interface that wraps the Update method.
type Updater interface {
	// Update updates the Quality and Days remaining.
	Update()
}

// Item describes an item sold by the Gilded Rose Inn.
type Item struct {
	Name    string
	Days    int
	Quality int
}

// Update updates the Quality and Days remaining.
func (i Item) Update() {
	// noop
}

// String outputs a string representation of an Item.
func (i *Item) String() string {
	return fmt.Sprintf("%s\t%d\t%d\t", i.Name, i.Days, i.Quality)
}

// Normal is a normal item.
type Normal struct {
	*Item
}

// Update updates the Quality and Days remaining.
func (i Normal) Update() {
	i.Days--
	if i.Quality == 0 {
		return
	}
	i.Quality--
	if i.Days <= 0 {
		i.Quality--
	}
}

// Brie is an "Aged Brie".
type Brie struct {
	*Item
}

// Update updates the Quality and Days remaining.
func (i Brie) Update() {
	i.Days--
	if i.Quality >= 50 {
		return
	}
	if i.Days <= 0 {
		i.Quality++
	}
	if i.Quality < 50 {
		i.Quality++
	}
}

// Backstage is "Backstage passes to a TAFKAL80ETC concert".
type Backstage struct {
	*Item
}

// Update updates the Quality and Days remaining.
func (i Backstage) Update() {
	i.Days--
	if i.Quality >= 50 {
		return
	}
	if i.Days < 0 {
		i.Quality = 0
		return
	}
	i.Quality++
	if i.Days < 10 {
		i.Quality++
	}
	if i.Days < 5 {
		i.Quality++
	}
}

// UpdateQuality ages the item by a day, and updates the Quality of the item.
func UpdateQuality(items ...*Item) {
	for _, item := range items {
		var t Updater
		switch item.Name {
		case "normal":
			t = Normal{item}
		case "Aged Brie":
			t = Brie{item}
		case "Backstage passes to a TAFKAL80ETC concert":
			t = Backstage{item}
		default:
			t = item
		}
		t.Update()
	}
}
