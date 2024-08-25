package main

import (
	"errors"
	"fmt"
	"go-poe-trade/helpers"
	"slices"
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

func getCurrencyName(text string, state *State) (CurrencyDetails, bool) {
	details := state.currency.CurrencyDetails
	// TODO: Run function that checks double names and store them in state
	forbiddenSplitNames := []string{"orb", "scroll", "shard", "lifeforce", "maven", "grand", "ichor"}

	i := slices.IndexFunc(details, func(n CurrencyDetails) bool {
		isCurrencyFound := n.Name == text || n.TradeID == text || string(n.ID) == text

		if !isCurrencyFound {
			tradeIdSplit := strings.Split(n.TradeID, "-")

			isFound := findAndExcludeForbidden(tradeIdSplit, forbiddenSplitNames, text)
			isCurrencyFound = isFound
		}

		if !isCurrencyFound {
			nameSplit := strings.Split(n.Name, " ")

			isFound := findAndExcludeForbidden(nameSplit, forbiddenSplitNames, text)
			isCurrencyFound = isFound
		}

		return isCurrencyFound
	})

	if i < 0 {
		return CurrencyDetails{}, false
	}

	return details[i], true
}

func findCurrencyByName(text string, state *State) (Lines, error) {
	found := Lines{}
	for _, curLine := range state.currency.Lines {
		if curLine.CurrencyTypeName == text {
			return curLine, nil
		}
	}
	return found, errors.New("No line found")
}

func parsePriceCheck(text string, state *State) {
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

	currencyNames, found := getCurrencyName(userInput[0], state)
	/* itemName, found := getItemByName(userInput[0], state) */

	if !found {
		fmt.Println("Currency not found!")
		return
	}

	// TODO: Searching for Chaos Orb by name will fail
	// Since Chaos is not in Lines - it's the defautl currency it'd be worth 1 chaos and whatever chaos/divine
	currency, err := findCurrencyByName(currencyNames.Name, state)
	if err != nil {
		fmt.Println("Couldn't find the currency!")
		return
	}

	chaosEquivalent := currency.ChaosEquivalent * float64(quantity)

	divPrice, chaosPrice := convertChaosToDivs(chaosEquivalent, state.divLine.ChaosEquivalent)
	currencyPrinter(currency.CurrencyTypeName, chaosEquivalent, divPrice, chaosPrice, quantity)
}
