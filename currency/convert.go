package currency

func ConvertChaosToDivs(itemPriceChaos float64, priceDiv float64) (int64, float64) {
	// casting as int so we have a nice even number
	divineEquivalent := int64(0)

	// if item price is lower than price of div
	// obvs its 0 div
	if itemPriceChaos >= priceDiv {
		divineEquivalent = int64(itemPriceChaos / priceDiv)
	}
	chaosReminder := itemPriceChaos - priceDiv*float64(divineEquivalent)
	return divineEquivalent, chaosReminder
}
