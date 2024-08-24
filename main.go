package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

type State struct {
	currency Currency
	divLine  Lines
	replMode string
}

func getCurrencyName(text string, state *State) (CurrencyDetails, bool) {
	details := state.currency.CurrencyDetails
	// TODO: Run function that checks double names and store them in state
	forbiddenSplitNames := []string{"orb", "scroll", "shard", "lifeforce", "maven", "grand", "ichor"}
	textLower := strings.ToLower(text)

	i := slices.IndexFunc(details, func(n CurrencyDetails) bool {
		isCurrencyFound := n.Name == textLower || n.TradeID == textLower || string(n.ID) == textLower

		if !isCurrencyFound {
			tradeIdSliced := strings.Split(n.TradeID, "-")

			for _, v := range tradeIdSliced {
				_, foundForbidden := slices.BinarySearch(forbiddenSplitNames, v)
				if !foundForbidden {
					_, found := slices.BinarySearch(tradeIdSliced, textLower)
					isCurrencyFound = found
					break
				}
			}
		}

		if !isCurrencyFound {
			nameSliced := strings.Split(n.Name, " ")

			for _, k := range nameSliced {
				nameToLowerCase := strings.ToLower(k)

				if textLower == nameToLowerCase {
					isCurrencyFound = true
					break
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

func isPriceCheck(text string) bool {
	inputSlice := strings.Split(text, " ")

	if inputSlice[0] == "pc" {
		return true
	}

	return false
}

func parsePriceCheck(text string, state *State) {
	input := strings.Split(text, " ")
	if len(input) < 2 {
		fmt.Println("No argument provided for price check")
		return
	}

	slice := input[1:]
	name, found := getCurrencyName(slice[0], state)
	if found {
		fmt.Println("Found: ", name.TradeID)
	} else {
		fmt.Println("Not found")
	}

	if slice[0] == "divine" {
		divPrice, chaosPrice := convertChaosToDivs(state.divLine.ChaosEquivalent, state.divLine.ChaosEquivalent)
		fmt.Printf("Divine price is %d in divs and %.1f in Chaos \n", divPrice, chaosPrice)
	}

}

func replParse(text string, state *State) {
	if text == "" {
		return
	}
	exitCommands := []string{"exit", "e", "quit", "Exit", ":q", "close"}
	_, isExitCommand := slices.BinarySearch(exitCommands, text)

	if isExitCommand {
		os.Exit(0)
		return
	}

	if isPriceCheck(text) {
		parsePriceCheck(text, state)
		return
	}

	fmt.Println("# No valid commant provided")

}

func main() {
	CURRENCY_URL := "https://poe.ninja/api/data/currencyoverview?league=Settlers&type=Currency"

	state := State{}

	getJson(CURRENCY_URL, &state.currency)
	extractDiv(&state)
	printCurrency(&state)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter text: ")
		scanner.Scan()

		text := scanner.Text()
		if len(text) != 0 {
			replParse(text, &state)
		}
	}
}
