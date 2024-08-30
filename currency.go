package main

import (
	"fmt"
	"maps"
	"strings"
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

type Line struct {
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
	Lines           []Line            `json:"lines"`
	CurrencyDetails []CurrencyDetails `json:"currencyDetails"`
}

type ItemLine struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Icon      string `json:"icon"`
	BaseType  string `json:"baseType"`
	StackSize int    `json:"stackSize"`
	ItemClass int    `json:"itemClass"`
	Sparkline struct {
		Data        []float64 `json:"data"`
		TotalChange float64   `json:"totalChange"`
	} `json:"sparkline"`
	LowConfidenceSparkline struct {
		Data        []float64 `json:"data"`
		TotalChange float64   `json:"totalChange"`
	} `json:"lowConfidenceSparkline"`
	ImplicitModifiers []any `json:"implicitModifiers"`
	ExplicitModifiers []struct {
		Text     string `json:"text"`
		Optional bool   `json:"optional"`
	} `json:"explicitModifiers"`
	FlavourText  string  `json:"flavourText"`
	ChaosValue   float64 `json:"chaosValue"`
	ExaltedValue float64 `json:"exaltedValue"`
	DivineValue  float64 `json:"divineValue"`
	Count        int     `json:"count"`
	DetailsID    string  `json:"detailsId"`
	TradeInfo    []any   `json:"tradeInfo"`
	ListingCount int     `json:"listingCount"`
}

type TradeItems struct {
	Lines       []ItemLine `json:"lines"`
	SharedNames []string
}

func (t *TradeItems) ParseSharedNames() {
	accumulator := make(map[string]int)

	for _, k := range t.Lines {
		splitNames := strings.Split(k.Name, " ")

		for _, name := range splitNames {

			nameLower := strings.ToLower(name)
			val, ok := accumulator[nameLower]

			if ok {
				accumulator[nameLower] = val + 1
			} else {
				accumulator[nameLower] = 1
			}
		}
	}

	for key, value := range maps.All(accumulator) {
		if value > 1 {
			t.SharedNames = append(t.SharedNames, key)
		}
	}

}
