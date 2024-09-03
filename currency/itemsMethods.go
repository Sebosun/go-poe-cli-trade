package currency

import (
	"errors"
	"maps"
	"strings"
)

func (t *TradeItems) ParseSharedNames() {
	accumulator := make(map[string]int)

	for _, k := range t.Lines {
		splitNames := strings.Split(k.Name, " ")

		for _, name := range splitNames {

			nameLower := strings.ToLower(name)
			val, ok := accumulator[nameLower]

			if ok {
				accumulator[nameLower] = val + 1
			} else {
				accumulator[nameLower] = 1
			}
		}
	}

	for key, value := range maps.All(accumulator) {
		if value > 1 {
			t.SharedNames = append(t.SharedNames, key)
		}
	}
}

func (t *TradeItems) FindItems(input string) (ItemLine, error) {
	forbiddenNames := t.SharedNames

	var exactMatch ItemLine
	exactMatchFound := false

	var closeMatch ItemLine
	closeMatchFound := false

	for _, item := range t.Lines {
		exactMatchFound = strings.ToLower(item.Name) == input || strings.ToLower(item.DetailsID) == input
		if exactMatchFound {
			exactMatch = item
			break
		}

		detailsSplit := strings.Split(item.DetailsID, "-")

		found := findAndExcludeForbidden(detailsSplit, forbiddenNames, input)

		if found {
			closeMatchFound = true
			closeMatch = item
		}

		nameSplit := strings.Split(item.Name, " ")

		found = findAndExcludeForbidden(nameSplit, forbiddenNames, input)

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

	return ItemLine{}, errors.New("Didnt find the item")
}
