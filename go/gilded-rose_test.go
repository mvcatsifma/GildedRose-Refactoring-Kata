package main

import "testing"

func Test_DecreasesQualityAndSellInOfNormalItem(t *testing.T) {
	items := []*Item{{name: "aName", quality: 1, sellIn: 1}}
	expectedQuality := 0
	expectedSellIn := 0

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_DecreasesQualityOfNormalItemBeyondSellByDate(t *testing.T) {
	items := []*Item{{name: "aName", quality: 2, sellIn: -1}}
	expectedQuality := 0
	expectedSellIn := -2

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_DoesNotDecreaseQualityAndSellInOfSulfurasItem(t *testing.T) {
	items := []*Item{{name: itemTypeSulfuras, quality: sulfurasQuality, sellIn: 1}}
	expectedQuality := 80
	expectedSellIn := 1

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_DoesNotDecreaseQualityAndSellInOfExpiredSulfurasItem(t *testing.T) {
	items := []*Item{{name: itemTypeSulfuras, quality: sulfurasQuality, sellIn: -1}}
	expectedQuality := 80
	expectedSellIn := -1

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_IncreasesByOneQualityOfBrieItem(t *testing.T) {
	items := []*Item{
		{name: itemTypeBrie, quality: 0, sellIn: 1},
	}
	expectedQuality := 1
	expectedSellIn := 0

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_IncreasesByTwoQualityOfExpiredBrieItem(t *testing.T) {
	items := []*Item{
		{name: itemTypeBrie, quality: 0, sellIn: -1},
	}
	expectedQuality := 2
	expectedSellIn := -2

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_IncreasesQualityOfBackstagePassWithSellInFive(t *testing.T) {
	items := []*Item{
		{name: itemTypeBackstagePass, quality: 0, sellIn: 5},
	}
	expectedQuality := 3
	expectedSellIn := 4

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_IncreasesQualityOfBackstagePassWithSellInTen(t *testing.T) {
	items := []*Item{
		{name: itemTypeBackstagePass, quality: 0, sellIn: 10},
	}
	expectedQuality := 2
	expectedSellIn := 9

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_IncreasesQualityOfBackstagePassWithSellInMoreThanTen(t *testing.T) {
	items := []*Item{
		{name: itemTypeBackstagePass, quality: 0, sellIn: 11},
	}
	expectedQuality := 1
	expectedSellIn := 10

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_QualityOfExpiredBackstagePassesIsZero(t *testing.T) {
	items := []*Item{
		{name: itemTypeBackstagePass, quality: 10, sellIn: -1},
	}
	expectedQuality := 0
	expectedSellIn := -2

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_QualityNeverLessThanZero(t *testing.T) {
	items := []*Item{{name: "aName", quality: 0, sellIn: 1}}
	expectedQuality := 0
	expectedSellIn := 0

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_QualityOfBrieNeverMoreThanFifty(t *testing.T) {
	items := []*Item{
		{name: itemTypeBrie, quality: 50, sellIn: -1},
	}
	expectedQuality := 50
	expectedSellIn := -2

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_QualityOfBackstagePassWithSellInFiveNeverMoreThanFifty(t *testing.T) {
	items := []*Item{
		{name: itemTypeBackstagePass, quality: 50, sellIn: 5},
	}
	expectedQuality := 50
	expectedSellIn := 4

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_QualityOfBackstagePassWithSellInTenNeverMoreThanFifty(t *testing.T) {
	items := []*Item{
		{name: itemTypeBackstagePass, quality: 50, sellIn: 10},
	}
	expectedQuality := 50
	expectedSellIn := 9

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_QualityOfConjuredItemDegradesTwiceAsFast(t *testing.T) {
	items := []*Item{
		{name: itemTypeConjured, quality: 10, sellIn: 1},
	}
	expectedQuality := 8
	expectedSellIn := 0

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}

func Test_QualityOfExpiredConjuredItemDegradesTwiceAsFast(t *testing.T) {
	items := []*Item{
		{name: itemTypeConjured, quality: 10, sellIn: -1},
	}
	expectedQuality := 6
	expectedSellIn := -2

	UpdateQuality(items)

	item := items[0]
	if item.quality != expectedQuality {
		t.Errorf("quality: Expected %v but got %v ", expectedQuality, item.quality)
	}
	if item.sellIn != expectedSellIn {
		t.Errorf("sellIn: Expected %v but got %v ", expectedSellIn, item.sellIn)
	}
}
