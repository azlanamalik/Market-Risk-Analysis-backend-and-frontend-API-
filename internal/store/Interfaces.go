package store
import (
	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
	"context"
)

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
