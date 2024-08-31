package main

import (
	"encoding/json"
	"fmt"
	"go-poe-trade/currency"
	"net/http"
)

func fetchCurrency(url string, currency *Currency) error {
	r, err := http.Get(url)

	if err != nil {
		return err
	}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&currency)

	if err != nil {
		return fmt.Errorf("Error fethcing the url")
	}

	return nil
}

func fetchItem(url string) (currency.TradeItems, error) {
	basicItem := currency.TradeItems{}

	r, err := http.Get(url)

	if err != nil {
		return basicItem, err
	}

	defer r.Body.Close()

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&basicItem)

	if err != nil {
		return basicItem, nil
	}

	return basicItem, nil
}
