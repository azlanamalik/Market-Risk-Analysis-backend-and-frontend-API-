//this file is for storing stuff in memory
package store

import (
	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
)

type MemoryStore struct {
    positions    map[string]map[string]domain.Position
    latestPrices map[string]domain.PriceTick
    risks        map[string]map[string]domain.PositionRisk
}