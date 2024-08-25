package poecurrency

import (
	"errors"
	"fmt"
	"go-poe-trade"
	"go-poe-trade/helpers"
	"slices"
	"strconv"
	"strings"
)

func getCurrencyName(text string, state *main.State) (CurrencyDetails, bool) {
	details := state.currency.CurrencyDetails
	// TODO: Run function that checks double names and store them in state
	forbiddenSplitNames := []string{"orb", "scroll", "shard", "lifeforce", "maven", "grand", "ichor"}
	textLower := strings.ToLower(text)

	// TODO: this part is broken atm
	i := slices.IndexFunc(details, func(n CurrencyDetails) bool {
		isCurrencyFound := n.Name == textLower || n.TradeID == textLower || string(n.ID) == textLower

		if !isCurrencyFound {
			tradeIdSplit := strings.Split(n.TradeID, "-")

			for _, v := range tradeIdSplit {
				_, foundForbidden := helpers.Find(forbiddenSplitNames, func(forbiddenName string) bool {
					return forbiddenName == v
				})

				if !foundForbidden {
					_, found := helpers.Find(tradeIdSplit, func(k string) bool {
						return k == textLower
					})
					isCurrencyFound = found
					break
				}
			}
		}

		if !isCurrencyFound {
			nameSliced := strings.Split(n.Name, " ")

			for _, itemPartialName := range nameSliced {
				_, foundForbidden := helpers.Find(forbiddenSplitNames, func(forbiddenName string) bool {
					return forbiddenName == strings.ToLower(itemPartialName)
				})

				if !foundForbidden {
					nameToLowerCase := strings.ToLower(itemPartialName)

					if textLower == nameToLowerCase {
						isCurrencyFound = true
						break
					}
				}
			}
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
	if !found {
		fmt.Println("Currency not found!")
		return
	}

	currency, err := findCurrencyByName(currencyNames.Name, state)
	if err != nil {
		fmt.Println("Couldn't find the currency!")
		return
	}

	chaosEquivalent := currency.ChaosEquivalent * float64(quantity)

	divPrice, chaosPrice := convertChaosToDivs(chaosEquivalent, state.divLine.ChaosEquivalent)
	currencyPrinter(currency.CurrencyTypeName, chaosEquivalent, divPrice, chaosPrice, quantity)
}
