package model

type Stock struct {
	Ticker        string  `json:"ticker"`
	Price         float64 `json:"price"`
	Change        float64 `json:"change"`
	PercentChange float64 `json:"percent_change"`
	Data          string  `json:"data"`
}
