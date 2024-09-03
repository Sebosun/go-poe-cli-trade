package main

import (
	"fmt"
	"go-poe-trade/currency"
	"go-poe-trade/helpers"
	"os"
	"strconv"
	"strings"
	"time"
)

type State struct {
	currency currency.Currency
	items    currency.TradeItems
	divLine  currency.Line
	replMode string
}

func (state *State) extractDiv() {
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

func (s *State) exportToCSV() error {
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
	for _, line := range s.currency.Lines {
		divine := s.divLine
		divineEquivalent, chaosReminder := currency.ConvertChaosToDivs(line.ChaosEquivalent, divine.ChaosEquivalent)

		dateString := fmt.Sprintf("%v,%v,%v,%v\n", line.CurrencyTypeName, helpers.FloatCSV(line.ChaosEquivalent), divineEquivalent, helpers.FloatCSV(chaosReminder))
		file.WriteString(dateString)
	}

	fmt.Printf("Saved as %v/%v on the drive\n", workingDirectory, FILE_NAME)
	return nil

}

func (state *State) printAllCurrency() {
	for n := range len(state.currency.Lines) {
		curLine := state.currency.Lines[n]

		divine := state.divLine
		divineEquivalent, chaosReminder := currency.ConvertChaosToDivs(curLine.ChaosEquivalent, divine.ChaosEquivalent)

		helpers.PrintCurrency(curLine.CurrencyTypeName, curLine.ChaosEquivalent, divineEquivalent, chaosReminder, 1)
	}

	for _, n := range state.items.Lines {
		chaos := n.ChaosValue
		divPrice, chaosPrice := currency.ConvertChaosToDivs(chaos, state.divLine.ChaosEquivalent)
		helpers.PrintCurrency(n.Name, n.ChaosValue, divPrice, chaosPrice, 1)
	}
}

func (state *State) priceCheck(text string) {
	input := strings.Split(text, " ")

	if len(input) < 2 {
		fmt.Println("No argument provided for price check")
		return
	}

	userInput := input[1:]

	quantity, quantityError := strconv.Atoi(userInput[0])

	if quantityError == nil {
		// Hack, cutting user input for example: pc 10 divine, we are left with 'divine'
		userInput = userInput[1:]
	} else {
		quantity = 1
	}

	isCurrencyFound := false
	userInputJoined := helpers.SliceJoinStrings(userInput, " ")
	cur, err := state.currency.GetCurrency(text)

	if err == nil {
		isCurrencyFound = true
	}

	if isCurrencyFound {
		chaosEquivalent := cur.ChaosEquivalent * float64(quantity)
		divPrice, chaosPrice := currency.ConvertChaosToDivs(chaosEquivalent, state.divLine.ChaosEquivalent)
		helpers.PrintCurrency(cur.CurrencyTypeName, chaosEquivalent, divPrice, chaosPrice, quantity)
		return
	}

	items, err := state.items.FindItems(userInputJoined)

	if err != nil {
		fmt.Println("Couldn't find the item")
		return
	}

	chaos := items.ChaosValue
	divPrice, chaosPrice := currency.ConvertChaosToDivs(chaos, state.divLine.ChaosEquivalent)
	helpers.PrintCurrency(items.Name, items.ChaosValue, divPrice, chaosPrice, quantity)
}
