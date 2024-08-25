package main

import (
	"encoding/json"
	"fmt"
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
