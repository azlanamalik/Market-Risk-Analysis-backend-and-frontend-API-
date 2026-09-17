package domain

import (
	"log/slog"
)

type Position struct {
	PortfolioID  string  `json:"portfolio_id"`
	Symbol       string  `json:"symbol"`
	Quantity     float64 `json:"quantity"` // make sure this is signed to signify long or short
	AveragePrice float64 `json:"average_price"`
}

/*
•	Positive quantity means long exposure.
•	Negative quantity means short exposure.
•	Zero means no open exposure and should normally not be stored as an active position.
*/
func (Position Position) Validator() bool {
	logger := slog.Default()
	if Position.PortfolioID == "" {
		logger.Error("invalid position portfolio ID", "portfolio_id", Position.PortfolioID)
		return false
	}
	if Position.Symbol == "" {
		logger.Error("invalid position symbol", "symbol", Position.Symbol)
		return false
	}
	if Position.Quantity == 0 {
		logger.Error("invalid position quantity", "quantity", Position.Quantity)
		return false
	}
	if Position.AveragePrice <= 0 {
		logger.Error("invalid position average price", "average_price", Position.AveragePrice)
		return false
	}
	return true
}
