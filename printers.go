package main

import (
	"fmt"
	"go-poe-trade/helpers"
	"os"
	"time"
)

func currencyPrinter(name string, rawChaos float64, divineEquivalent int64, chaosRemaining float64, quantity int) {
	if quantity <= 1 {
		fmt.Printf("%-35s \t Raw: %.2f | Divs %d Chaos %.2f\n", name, rawChaos, divineEquivalent, chaosRemaining)
		return
	}

	fmt.Printf("%v %-35s \t Raw: %.2f | Divs %d Chaos %.2f\n", quantity, name, rawChaos, divineEquivalent, chaosRemaining)
}

func exportToCSV(state *State) error {
	FILE_NAME := "currency.csv"
	file, err := os.Create(FILE_NAME)

	workingDirectory, err := os.Getwd()

	if err != nil {
		fmt.Printf("Error saving currency to the disk: %v \n", err)
		return err
	}

	defer file.Close()

	time := time.Now().Local()
	currentTime := time.UTC()
	file.WriteString(fmt.Sprintf("currency name,chaos,div,chaos remainer,%v\n", currentTime))
	for _, line := range state.currency.Lines {
		divine := state.divLine
		divineEquivalent, chaosReminder := convertChaosToDivs(line.ChaosEquivalent, divine.ChaosEquivalent)

		dateString := fmt.Sprintf("%v,%v,%v,%v\n", line.CurrencyTypeName, helpers.FloatCSV(line.ChaosEquivalent), divineEquivalent, helpers.FloatCSV(chaosReminder))
		file.WriteString(dateString)
	}

	fmt.Printf("Saved as %v/%v on the drive\n", workingDirectory, FILE_NAME)
	return nil
}
