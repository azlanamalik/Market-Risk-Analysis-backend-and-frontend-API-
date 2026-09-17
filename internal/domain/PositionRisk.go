package domain
import (
	"log/slog"
)

type PositionRisk struct {
    PortfolioID   string
    Symbol        string
    MarketPrice   float64
    MarketValue   float64
    UnrealizedPnL float64
}


func (PositionRisk PositionRisk) Validator() bool {
	logger := slog.Default()
	if PositionRisk.PortfolioID == ""{
		return false
	}else if PositionRisk.Symbol == "" {
		logger.Error("invalid position symbol", "symbol", PositionRisk.Symbol)
		return false
	}else if PositionRisk.MarketPrice < 0{
		logger.Error("invalid position MarketPrice", "MarketPrice", PositionRisk.MarketPrice)
		return false
	}else if PositionRisk.MarketValue < 0{
		logger.Error("invalid position MarketValue", "MarketValue", PositionRisk.MarketValue)
		return false
	}else if PositionRisk.UnrealizedPnL < 0{
		logger.Error("invalid position UnrealizedPnL", "UnrealizedPnL", PositionRisk.UnrealizedPnL)
		return false
	}else{
		return true
	}
}