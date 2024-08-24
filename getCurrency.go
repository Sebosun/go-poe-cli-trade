package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type Receive struct {
	ID                int       `json:"id"`
	LeagueID          int       `json:"league_id"`
	PayCurrencyID     int       `json:"pay_currency_id"`
	GetCurrencyID     int       `json:"get_currency_id"`
	SampleTimeUtc     time.Time `json:"sample_time_utc"`
	Count             int       `json:"count"`
	Value             float64   `json:"value"`
	DataPointCount    int       `json:"data_point_count"`
	IncludesSecondary bool      `json:"includes_secondary"`
	ListingCount      int       `json:"listing_count"`
}

type PaySparkLine struct {
	Data        []float64 `json:"data"`
	TotalChange float64   `json:"totalChange"`
}

type ReceiveSparkLine struct {
	Data        []float64 `json:"data"`
	TotalChange float64   `json:"totalChange"`
}

type LowConfidencePaySparkLine struct {
	Data        []float64 `json:"data"`
	TotalChange float64   `json:"totalChange"`
}

type LowConfidenceReceiveSparkLine struct {
	Data        []float64 `json:"data"`
	TotalChange float64   `json:"totalChange"`
}

type CurrencyDetails struct {
	ID      int    `json:"id"`
	Icon    string `json:"icon,omitempty"`
	Name    string `json:"name"`
	TradeID string `json:"tradeId,omitempty"`
}

type Pay struct {
	ID                int       `json:"id"`
	LeagueID          int       `json:"league_id"`
	PayCurrencyID     int       `json:"pay_currency_id"`
	GetCurrencyID     int       `json:"get_currency_id"`
	SampleTimeUtc     time.Time `json:"sample_time_utc"`
	Count             int       `json:"count"`
	Value             float64   `json:"value"`
	DataPointCount    int       `json:"data_point_count"`
	IncludesSecondary bool      `json:"includes_secondary"`
	ListingCount      int       `json:"listing_count"`
}

type Lines struct {
	CurrencyTypeName              string                        `json:"currencyTypeName"`
	Pay                           Pay                           `json:"pay,omitempty"`
	Receive                       Receive                       `json:"receive,omitempty"`
	PaySparkLine                  PaySparkLine                  `json:"paySparkLine"`
	ReceiveSparkLine              ReceiveSparkLine              `json:"receiveSparkLine"`
	ChaosEquivalent               float64                       `json:"chaosEquivalent"`
	LowConfidencePaySparkLine     LowConfidencePaySparkLine     `json:"lowConfidencePaySparkLine"`
	LowConfidenceReceiveSparkLine LowConfidenceReceiveSparkLine `json:"lowConfidenceReceiveSparkLine"`
	DetailsID                     string                        `json:"detailsId"`
}

type Currency struct {
	Lines           []Lines           `json:"lines"`
	CurrencyDetails []CurrencyDetails `json:"currencyDetails"`
}

func getJson(url string, dupa *Currency) error {
	r, err := http.Get(url)

	if err != nil {
		return err
	}

	defer r.Body.Close()
	fmt.Println(r.Body)

	decoder := json.NewDecoder(r.Body)

	err = decoder.Decode(&dupa)

	if err != nil {
		return fmt.Errorf("Error fethcing the url")
	}

	return nil
}
