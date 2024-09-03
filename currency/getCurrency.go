package currency

import (
	"errors"
)

func (t *Currency) GetCurrency(input string) (Line, error) {
	currencyName, found := findCurrencyName(input, t.CurrencyDetails)
	if !found {
		return Line{}, errors.New("Couldn't find the currency name")
	}

	foundCurrency, err := findCurrencyByName(currencyName.Name, t.Lines)

	if err != nil {
		return Line{}, errors.New("Couldn't find the currency")
	}
	return foundCurrency, nil
}
