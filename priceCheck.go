package main

import (
	"fmt"
	"go-poe-trade/helpers"
	"strconv"
	"strings"
)

func findAndExcludeForbidden(splitName []string, forbidden []string, target string) bool {
	isFound := false
	for _, name := range splitName {
		_, foundForbidden := helpers.Find(forbidden, func(forbiddenName string) bool {
			return forbiddenName == strings.ToLower(name)
		})

		if foundForbidden {
			continue
		}

		_, found := helpers.Find(splitName, func(k string) bool {
			return strings.ToLower(k) == target
		})

		if found {
			isFound = found
			break
		}
	}

	return isFound
}

func priceCheck(text string, state *State) {
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

	checkItems := false
	currency, err := getCurrency(userInput[0], state)

	if err != nil {
		checkItems = true
	}

	if !checkItems {
		chaosEquivalent := currency.ChaosEquivalent * float64(quantity)
		divPrice, chaosPrice := convertChaosToDivs(chaosEquivalent, state.divLine.ChaosEquivalent)
		currencyPrinter(currency.CurrencyTypeName, chaosEquivalent, divPrice, chaosPrice, quantity)
		return
	}

	items, err := getItems(userInput[0], state)

	if err != nil {
		fmt.Println("Couldn't find the item")
		return
	}

	chaos := items.ChaosValue
	divPrice, chaosPrice := convertChaosToDivs(chaos, state.divLine.ChaosEquivalent)
	currencyPrinter(items.Name, items.ChaosValue, divPrice, chaosPrice, quantity)

}
