package main

import (
	"errors"
	"slices"
	"strings"
)

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

func findCurrencyByName(text string, state *State) (Line, error) {
	found := Line{}
	for _, curLine := range state.currency.Lines {
		if curLine.CurrencyTypeName == text {
			return curLine, nil
		}
	}
	return found, errors.New("No line found")
}

func getCurrency(input string, state *State) (Line, error) {
	currencyName, found := getCurrencyName(input, state)
	if !found {
		return Line{}, errors.New("Couldn't find the currency name")
	}

	currency, err := findCurrencyByName(currencyName.Name, state)

	if err != nil {
		return Line{}, errors.New("Couldn't find the currency")
	}
	return currency, nil
}
