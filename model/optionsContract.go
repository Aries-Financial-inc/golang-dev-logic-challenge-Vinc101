package model

import (
	"time"
)

// OptionsType defines the type of the option: call or put
type OptionsType string

const (
	Call OptionsType = "call"
	Put  OptionsType = "put"
)

// LongShortType defines if the option is long or short
type LongShortType string

const (
	Long  LongShortType = "long"
	Short LongShortType = "short"
)

// OptionsContract represents the data structure of an options contract
type OptionsContract struct {
	Type           OptionsType   `json:"type"`
	StrikePrice    float64       `json:"strike_price"`
	Bid            float64       `json:"bid"`
	Ask            float64       `json:"ask"`
	ExpirationDate time.Time     `json:"expiration_date"`
	LongShort      LongShortType `json:"long_short"`
}

// AnalysisResponse represents the data structure of the analysis result
type AnalysisResponse struct {
	XYValues        []XYValue `json:"xy_values"`
	MaxProfit       float64   `json:"max_profit"`
	MaxLoss         float64   `json:"max_loss"`
	BreakEvenPoints []float64 `json:"break_even_points"`
}

// XYValue represents a pair of X and Y values
type XYValue struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

type AnalysisResult struct {
	GraphData       []XYValue `json:"graph_data"`
	MaxProfit       float64   `json:"max_profit"`
	MaxLoss         float64   `json:"max_loss"`
	BreakEvenPoints []float64 `json:"break_even_points"`
}
