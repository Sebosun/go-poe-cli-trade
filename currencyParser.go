package main

import (
	"fmt"
)

func extractDiv(state *State) {
	DIVINE_ID := "divine-orb"
	curLines := state.currency.Lines
	for n := range len(curLines) {
		curLine := curLines[n]
		if curLine.DetailsID == DIVINE_ID {
			state.divLine = curLine
			break
		}
	}
}

func convertChaosToDivs(itemPriceChaos float64, priceDiv float64) (int64, float64) {
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

func printCurrency(state *State) {
	for n := range len(state.currency.Lines) {
		curLine := state.currency.Lines[n]

		divine := state.divLine
		divineEquivalent, chaosReminder := convertChaosToDivs(curLine.ChaosEquivalent, divine.ChaosEquivalent)

		fmt.Printf("Pay %v chaos for %s aka %s. or %d divines %.2f chaoses \n", curLine.ChaosEquivalent, curLine.CurrencyTypeName, curLine.DetailsID, divineEquivalent, chaosReminder)
		/* fmt.Printf("Result %v %s %v %f \n", curLine.ChaosEquivalent, curLine.CurrencyTypeName, divineEquivalent, chaosReminder) */
	}
}
