package risk

import (
	"log/slog"
	"fmt"
	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
)

func risk(priceTick domain.PriceTick, position domain.Position) (marketPrice, marketValue, unrealizedPnL float64) {
	logger := slog.Default()
	if !priceTick.Validator() || priceTick.Symbol != position.Symbol {
		logger.Error("error in validation of data", "Validator : " ,priceTick.Validator() , "mismatch : ",  fmt.Sprintf( "price tick symbol : " + priceTick.Symbol + " mismatched with position : " + position.Symbol ) )
		return 0, 0, 0
	}

	marketPrice = priceTick.Midpoint()
	marketValue = position.Quantity * marketPrice
	unrealizedPnL = position.Quantity * (marketPrice - position.AveragePrice)
	return
}
