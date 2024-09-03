package currency

import (
	"errors"
	"go-poe-trade/helpers"
	"slices"
	"strings"
)

func findAndExcludeForbidden(splitName []string, forbidden []string, input string) bool {
	for _, name := range splitName {
		splitTarget := strings.Split(input, " ")
		for _, target := range splitTarget {
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
				return true
			}
		}
	}

	return false
}

func findCurrencyName(text string, curDetails []CurrencyDetails) (CurrencyDetails, bool) {
	// TODO: Run function that checks double names and store them in state
	forbiddenSplitNames := []string{"orb", "scroll", "shard", "lifeforce", "maven", "grand", "ichor"}

	i := slices.IndexFunc(curDetails, func(n CurrencyDetails) bool {
		isCurrencyFound := n.Name == text || n.TradeID == text

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

	return curDetails[i], true
}

func findCurrencyByName(text string, lines []Line) (Line, error) {
	found := Line{}
	for _, curLine := range lines {
		if curLine.CurrencyTypeName == text {
			return curLine, nil
		}
	}
	return found, errors.New("No line found")
}
