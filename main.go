package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type State struct {
	currency Currency
	items    TradeItems
	divLine  Lines
	replMode string
}

func main() {
	CURRENCY_URL := "https://poe.ninja/api/data/currencyoverview?league=Settlers&type=Currency"
	SCARABS_URL := "https://poe.ninja/api/data/itemoverview?league=Settlers&type=Scarab"

	state := State{}

	fetchCurrency(CURRENCY_URL, &state.currency)
	item, err := fetchItem(SCARABS_URL)
	state.items.Lines = append(state.items.Lines, item.Lines...)

	if err != nil {
		fmt.Println("Error fetching scarabs")
	}

	extractDiv(&state)

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("Enter text: ")
		scanner.Scan()

		text := scanner.Text()
		if len(text) != 0 {
			text = strings.Trim(text, " ")
			text = strings.ToLower(text)
			replParse(text, &state)
		}
	}
}
