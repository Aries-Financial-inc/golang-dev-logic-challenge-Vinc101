package routes

import (
	"net/http"

	"github.com/Vinc101/golang-dev-logic-challenge-Vinc101/controllers"
	"github.com/Vinc101/golang-dev-logic-challenge-Vinc101/model"
	"github.com/gin-gonic/gin"
)

// AnalysisResult structure for the response body

// GraphPoint structure for X & Y values of the risk & reward graph
type GraphPoint struct {
	X float64 `json:"x"`
	Y float64 `json:"y"`
}

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.POST("/analyze", func(c *gin.Context) {
		var contracts []model.OptionsContract

		if err := c.ShouldBindJSON(&contracts); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		if len(contracts) > 4 {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Cannot process more than 4 options contracts"})
			return
		}

		responseBody := controllers.CalculateAnalysis(contracts)

		response := model.AnalysisResult{
			GraphData:       responseBody.XYValues,
			MaxProfit:       responseBody.MaxProfit,
			MaxLoss:         responseBody.MaxLoss,
			BreakEvenPoints: responseBody.BreakEvenPoints,
		}

		c.JSON(http.StatusOK, response)
	})

	return router
}
