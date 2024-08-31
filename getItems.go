package main

import (
	"errors"
	"go-poe-trade/currency"
	"strings"
)

func getItems(input string, state *State) (currency.ItemLine, error) {
	forbiddenSplitNames := state.items.SharedNames

	var exactMatch currency.ItemLine
	exactMatchFound := false

	var closeMatch currency.ItemLine
	closeMatchFound := false

	for _, item := range state.items.Lines {
		exactMatchFound = strings.ToLower(item.Name) == input || strings.ToLower(item.DetailsID) == input
		if exactMatchFound {
			exactMatch = item
			break
		}

		detailsSplit := strings.Split(item.DetailsID, "-")

		found := findAndExcludeForbidden(detailsSplit, forbiddenSplitNames, input)

		if found {
			closeMatchFound = true
			closeMatch = item
		}

		nameSplit := strings.Split(item.Name, " ")

		found = findAndExcludeForbidden(nameSplit, forbiddenSplitNames, input)

		if found {
			closeMatchFound = true
			closeMatch = item
		}
	}

	if exactMatchFound {
		return exactMatch, nil
	}

	if closeMatchFound {
		return closeMatch, nil
	}

	return currency.ItemLine{}, errors.New("Didnt find the item")
}
