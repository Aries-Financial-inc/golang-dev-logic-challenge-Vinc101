package tests

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"

	"github.com/Vinc101/golang-dev-logic-challenge-Vinc101/model"
	"github.com/Vinc101/golang-dev-logic-challenge-Vinc101/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestOptionsContractModelValidation(t *testing.T) {
	validContract := model.OptionsContract{
		Type:           model.Call,
		StrikePrice:    100.0,
		Bid:            10.0,
		Ask:            12.0,
		ExpirationDate: time.Now().AddDate(0, 1, 0),
		LongShort:      model.Long,
	}
	assert.NoError(t, validateOptionsContract(validContract))

	// Test invalid OptionsContract (invalid type)
	invalidContractType := model.OptionsContract{
		Type:           "invalid",
		StrikePrice:    100.0,
		Bid:            10.0,
		Ask:            12.0,
		ExpirationDate: time.Now().AddDate(0, 1, 0),
		LongShort:      model.Long,
	}
	assert.Error(t, validateOptionsContract(invalidContractType))

	// Test invalid OptionsContract (invalid long/short type)
	invalidLongShort := model.OptionsContract{
		Type:           model.Call,
		StrikePrice:    100.0,
		Bid:            10.0,
		Ask:            12.0,
		ExpirationDate: time.Now().AddDate(0, 1, 0),
		LongShort:      "invalid",
	}
	assert.Error(t, validateOptionsContract(invalidLongShort))
}

func validateOptionsContract(contract model.OptionsContract) error {
	validTypes := map[model.OptionsType]bool{
		model.Call: true,
		model.Put:  true,
	}
	if !validTypes[contract.Type] {
		return fmt.Errorf("invalid option type")
	}

	validLongShort := map[model.LongShortType]bool{
		model.Long:  true,
		model.Short: true,
	}
	if !validLongShort[contract.LongShort] {
		return fmt.Errorf("invalid long/short type")
	}
	return nil
}

func TestAnalysisEndpoint(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	contracts := []model.OptionsContract{
		{
			Type:           model.Call,
			StrikePrice:    100.0,
			Bid:            10.0,
			Ask:            12.0,
			ExpirationDate: time.Now().AddDate(0, 1, 0),
			LongShort:      model.Long,
		},
		{
			Type:           model.Put,
			StrikePrice:    80.0,
			Bid:            5.0,
			Ask:            6.0,
			ExpirationDate: time.Now().AddDate(0, 1, 0),
			LongShort:      model.Short,
		},
	}

	contractJSON, _ := json.Marshal(contracts)
	req, _ := http.NewRequest(http.MethodPost, "/analyze", bytes.NewBuffer(contractJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var result model.AnalysisResult
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	assert.NoError(t, err)
	assert.NotEmpty(t, result.GraphData)
	assert.NotZero(t, result.MaxProfit)
	assert.NotZero(t, result.MaxLoss)
	assert.NotEmpty(t, result.BreakEvenPoints)
}

func TestIntegration(t *testing.T) {
	gin.SetMode(gin.TestMode)
	router := routes.SetupRouter()

	contracts := []model.OptionsContract{
		{
			Type:           model.Call,
			StrikePrice:    100.0,
			Bid:            10.0,
			Ask:            12.0,
			ExpirationDate: time.Now().AddDate(0, 1, 0),
			LongShort:      model.Long,
		},
		{
			Type:           model.Put,
			StrikePrice:    80.0,
			Bid:            5.0,
			Ask:            6.0,
			ExpirationDate: time.Now().AddDate(0, 1, 0),
			LongShort:      model.Short,
		},
	}

	contractJSON, _ := json.Marshal(contracts)
	req, _ := http.NewRequest(http.MethodPost, "/analyze", bytes.NewBuffer(contractJSON))
	req.Header.Set("Content-Type", "application/json")

	resp := httptest.NewRecorder()
	router.ServeHTTP(resp, req)

	assert.Equal(t, http.StatusOK, resp.Code)

	var result model.AnalysisResult
	err := json.Unmarshal(resp.Body.Bytes(), &result)
	assert.NoError(t, err)

	assert.NotEmpty(t, result.GraphData)
	assert.NotZero(t, result.MaxProfit)
	assert.NotZero(t, result.MaxLoss)
	assert.NotEmpty(t, result.BreakEvenPoints)
}
