package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type State struct {
	currency Currency
	divLine  Lines
	replMode string
}

func main() {
	CURRENCY_URL := "https://poe.ninja/api/data/currencyoverview?league=Settlers&type=Currency"

	state := State{}

	getJson(CURRENCY_URL, &state.currency)
	extractDiv(&state)
	/* printCurrency(&state) */

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
