package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	// extract to separate function
	LEAGUE := "Settlers"

	// CURRENCY := []string{"Fragments", "KalguuranRune"}
	ITEMS := []string{"Scarabs", "Tattoo", "Omen", "DivinationCard", "Artifact", "Oil"}
	CURRENCY_URL := fmt.Sprintf("https://poe.ninja/api/data/currencyoverview?league=%s&type=Currency", LEAGUE)
	ITEMS_BASE_URL := fmt.Sprintf("https://poe.ninja/api/data/itemoverview?league=%s&type=", LEAGUE)

	state := State{}

	err := fetchCurrency(CURRENCY_URL, &state.currency)
	if err != nil {
		return
	}

	for _, item := range ITEMS {
		item, err := fetchItem(fmt.Sprintf(ITEMS_BASE_URL + item))

		time.Sleep(500)

		if err != nil {
			fmt.Println("Error fetching scarabs")
			continue
		}

		state.items.Lines = append(state.items.Lines, item.Lines...)
	}

	state.extractDiv()

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
