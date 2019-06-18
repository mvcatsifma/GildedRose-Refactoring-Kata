package main

import "math"

type Item struct {
	name            string
	sellIn, quality int
}

const itemTypeSulfuras = "Sulfuras, Hand of Ragnaros"
const itemTypeBrie = "Aged Brie"
const itemTypeBackstagePass = "Backstage passes to a TAFKAL80ETC concert"
const itemTypeConjured = "Conjured"

const maxQuality = 50
const minQuality = 0

const sulfurasQuality = 80

const defaultQualityIncrease = 1
const defaultQualityDecrease = 1

const expiredQualityIncrease = 2
const expiredQualityDecrease = 2

func UpdateQuality(items []*Item) {
	for _, item := range items {
		updateQuality(item, item.sellIn < 0)
		updateSellIn(item)
	}
}

func updateQuality(item *Item, expired bool) {
	if !expired {
		handle(item)
	} else {
		handleExpired(item)
	}
}

func handle(item *Item) {
	switch item.name {
	case itemTypeBackstagePass:
		increase := defaultQualityIncrease // quality of backstage passes increases by 1
		if item.sellIn <= 5 {
			increase = 3 // quality increases by 3 when there are 5 days or less
		} else if item.sellIn <= 10 {
			increase = 2 // quality increases by 2 when there are 10 days or less
		}
		item.quality = int(math.Min(maxQuality, float64(item.quality+increase)))
	case itemTypeBrie:
		item.quality = int(math.Min(maxQuality, float64(item.quality+defaultQualityIncrease))) // "Aged Brie" increases in Quality the older it gets
	case itemTypeSulfuras:
		// sulfuras' quality does not change
	case itemTypeConjured:
		item.quality = int(math.Max(minQuality, float64(item.quality-defaultQualityDecrease*2))) // by default items decrease in quality
	default:
		item.quality = int(math.Max(minQuality, float64(item.quality-defaultQualityDecrease))) // by default items decrease in quality
	}
}

func handleExpired(item *Item) {
	switch item.name {
	case itemTypeBackstagePass:
		item.quality = minQuality // expired backstage passes have zero value
	case itemTypeBrie:
		item.quality = int(math.Min(maxQuality, float64(item.quality+expiredQualityIncrease))) // "Aged Brie" increases in Quality the older it gets
	case itemTypeSulfuras:
		// sulfuras' quality does not change
	case itemTypeConjured:
		item.quality = int(math.Max(minQuality, float64(item.quality-expiredQualityDecrease*2))) // by default items decrease in quality
	default:
		item.quality = int(math.Max(minQuality, float64(item.quality-expiredQualityDecrease))) // expired items degrade twice as fast
	}
}

func updateSellIn(item *Item) {
	switch item.name {
	case itemTypeSulfuras:
		// sulfuras' sellIn does not change
	default:
		item.sellIn = item.sellIn - 1
	}
}
