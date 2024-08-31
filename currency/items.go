package currency

import (
	"maps"
	"strings"
)

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
