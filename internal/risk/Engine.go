package risk

import (
	"log/slog"
	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
	"context"
)
func risk(priceTick domain.PriceTick, position domain.Position) (domain.PositionRisk) {
	logger := slog.Default()
	valid := priceTick.Validator() 
	if !valid|| priceTick.Symbol != position.Symbol {
		logger.Error("error in validation of data", "Validator : " ,valid , "mismatch : ",  "price tick symbol : " + priceTick.Symbol + " mismatched with position : " + position.Symbol)
		return domain.PositionRisk{
		PortfolioID: "error loading",
		Symbol: "error loading",
		MarketPrice: 1.0,
		MarketValue: 1.0,
		UnrealizedPnL: 1.0,
	}//dont terminate as thining on a scale of full network we are sending constant packages it is bad if we terminate program on a slight change of dat
	}

	marketPrice := priceTick.Midpoint()
	marketValue := position.Quantity * marketPrice
	unrealizedPnL := position.Quantity * (marketPrice - position.AveragePrice)
	return domain.PositionRisk{
		PortfolioID: position.PortfolioID,
		Symbol: position.Symbol,
		MarketPrice: marketPrice,
		MarketValue: marketValue,
		UnrealizedPnL: unrealizedPnL,
	}
}

type RiskRepository interface {
    SaveTick(ctx context.Context, tick domain.PriceTick) error
    GetPositionsBySymbol(ctx context.Context, symbol string) ([]domain.Position, error)
    SaveRisk(ctx context.Context, marketPrice, marketValue, unrealizedPnL float64) error
}//ctx context.Context key as now we can control a cancellation request quicker

type RiskProcessor struct {
    store RiskRepository
}//this is to use the RiskRepo: i dont care what you do (and who you are) just have those 3 method signitures for SQL implimentation
//Dependency Inversion Principle (from SOLID) layer of abstraction so we donthaveto worry about the actual details of code.
//Thinknig in the future on implimentation with SQL and live data
