package main

import "fmt"

func currencyPrinter(name string, rawChaos float64, divineEquivalent int64, chaosRemaining float64, quantity int) {
	if quantity <= 1 {
		fmt.Printf("%-35s \t Raw: %.2f | Divs %d Chaos %.2f\n", name, rawChaos, divineEquivalent, chaosRemaining)
		return
	}

	fmt.Printf("%v %-35s \t Raw: %.2f | Divs %d Chaos %.2f\n", quantity, name, rawChaos, divineEquivalent, chaosRemaining)
}
