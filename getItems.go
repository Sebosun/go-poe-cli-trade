package main

import (
	"errors"
	"slices"
	"strings"
)

func getItems(input string, state *State) (ItemLine, error) {
	forbiddenSplitNames := []string{"scarab"}

	i := slices.IndexFunc(state.items.Lines, func(item ItemLine) bool {
		isFound := item.Name == input || item.DetailsID == input

		if !isFound {
			detailsSplit := strings.Split(item.DetailsID, "-")

			found := findAndExcludeForbidden(detailsSplit, forbiddenSplitNames, input)
			isFound = found
		}

		if !isFound {
			nameSplit := strings.Split(item.Name, " ")

			found := findAndExcludeForbidden(nameSplit, forbiddenSplitNames, input)
			isFound = found
		}

		return isFound
	})

	if i >= 0 {
		return state.items.Lines[i], nil
	}

	return ItemLine{}, errors.New("Didnt find the item")
}
