package controllers

import (
	"encoding/json"
	"math"
	"net/http"

	"github.com/Vinc101/golang-dev-logic-challenge-Vinc101/model"
)

func AnalysisHandler(w http.ResponseWriter, r *http.Request) {
	var contracts []model.OptionsContract
	if err := json.NewDecoder(r.Body).Decode(&contracts); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	if len(contracts) > 4 {
		http.Error(w, "Cannot process more than 4 options contracts", http.StatusBadRequest)
		return
	}

	xyValues := calculateXYValues(contracts)
	maxProfit := calculateMaxProfit(contracts)
	maxLoss := calculateMaxLoss(contracts)
	breakEvenPoints := calculateBreakEvenPoints(contracts)

	response := model.AnalysisResponse{
		XYValues:        xyValues,
		MaxProfit:       maxProfit,
		MaxLoss:         maxLoss,
		BreakEvenPoints: breakEvenPoints,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func CalculateAnalysis(contracts []model.OptionsContract) model.AnalysisResponse {
	xyValues := calculateXYValues(contracts)
	maxProfit := calculateMaxProfit(contracts)
	maxLoss := calculateMaxLoss(contracts)
	breakEvenPoints := calculateBreakEvenPoints(contracts)

	return model.AnalysisResponse{
		XYValues:        xyValues,
		MaxProfit:       maxProfit,
		MaxLoss:         maxLoss,
		BreakEvenPoints: breakEvenPoints,
	}
}

func calculateXYValues(contracts []model.OptionsContract) []model.XYValue {
	var xyValues []model.XYValue

	// Define a range for X values (underlying prices at expiry)
	for price := 0.0; price <= 2*findMaxStrikePrice(contracts); price += 1.0 {
		profitLoss := calculateProfitLossAtPrice(contracts, price)
		xyValues = append(xyValues, model.XYValue{X: price, Y: profitLoss})
	}

	return xyValues
}

func calculateMaxProfit(contracts []model.OptionsContract) float64 {
	xyValues := calculateXYValues(contracts)
	return findMaxYValue(xyValues)
}

func calculateMaxLoss(contracts []model.OptionsContract) float64 {
	xyValues := calculateXYValues(contracts)
	return findMinYValue(xyValues)
}

func calculateBreakEvenPoints(contracts []model.OptionsContract) []float64 {
	xyValues := calculateXYValues(contracts)
	return findBreakEvenPoints(xyValues)
}

func findMaxStrikePrice(contracts []model.OptionsContract) float64 {
	maxPrice := 0.0
	for _, contract := range contracts {
		if contract.StrikePrice > maxPrice {
			maxPrice = contract.StrikePrice
		}
	}
	return maxPrice
}

func calculateProfitLossAtPrice(contracts []model.OptionsContract, price float64) float64 {
	profitLoss := 0.0
	for _, contract := range contracts {
		if contract.Type == model.Call {
			if contract.LongShort == model.Long {
				profitLoss += math.Max(0, price-contract.StrikePrice) - contract.Ask
			} else {
				profitLoss += contract.Bid - math.Max(0, price-contract.StrikePrice)
			}
		} else {
			if contract.LongShort == model.Long {
				profitLoss += math.Max(0, contract.StrikePrice-price) - contract.Ask
			} else {
				profitLoss += contract.Bid - math.Max(0, contract.StrikePrice-price)
			}
		}
	}
	return profitLoss
}

func findMaxYValue(xyValues []model.XYValue) float64 {
	maxValue := xyValues[0].Y
	for _, value := range xyValues {
		if value.Y > maxValue {
			maxValue = value.Y
		}
	}
	return maxValue
}

func findMinYValue(xyValues []model.XYValue) float64 {
	minValue := xyValues[0].Y
	for _, value := range xyValues {
		if value.Y < minValue {
			minValue = value.Y
		}
	}
	return minValue
}

func findBreakEvenPoints(xyValues []model.XYValue) []float64 {
	var breakEvenPoints []float64
	for i := 1; i < len(xyValues); i++ {
		if (xyValues[i-1].Y < 0 && xyValues[i].Y >= 0) || (xyValues[i-1].Y > 0 && xyValues[i].Y <= 0) {
			breakEvenPoints = append(breakEvenPoints, xyValues[i].X)
		}
	}
	return breakEvenPoints
}
