// this file is for storing stuff in memory
package store

import (
	"errors"

	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
)

type MemoryStore struct {
	positions    map[string]map[string]domain.Position
	latestPrices map[string]domain.PriceTick
	risks        map[string]map[string]domain.PositionRisk
}

func (store *MemoryStore) InsertPosition(position domain.Position) error {
	if store.positions == nil {
		store.positions = make(map[string]map[string]domain.Position)
	}
	if store.positions[position.PortfolioID] == nil {
		store.positions[position.PortfolioID] = make(map[string]domain.Position)
	}
	if _, exists := store.positions[position.PortfolioID][position.Symbol]; exists {
		return errors.New("position already exists")
	}

	store.positions[position.PortfolioID][position.Symbol] = position
	return nil
}

func (store *MemoryStore) UpdatePosition(position domain.Position) error {
	if store.positions == nil || store.positions[position.PortfolioID] == nil {
		return errors.New("position does not exist")
	}
	if _, exists := store.positions[position.PortfolioID][position.Symbol]; !exists {
		return errors.New("position does not exist")
	}

	store.positions[position.PortfolioID][position.Symbol] = position
	return nil
}

func (store *MemoryStore) GetPosition(portfolioID, symbol string) (domain.Position, error) {
	if store.positions == nil || store.positions[portfolioID] == nil {
		return domain.Position{}, errors.New("position does not exist")
	}

	position, exists := store.positions[portfolioID][symbol]
	if !exists {
		return domain.Position{}, errors.New("position does not exist")
	}

	return position, nil
}

func (store *MemoryStore) InsertPriceTick(priceTick domain.PriceTick) error {
	if store.latestPrices == nil {
		store.latestPrices = make(map[string]domain.PriceTick)
	}
	if _, exists := store.latestPrices[priceTick.Symbol]; exists {
		return errors.New("price tick already exists")
	}

	store.latestPrices[priceTick.Symbol] = priceTick
	return nil
}

func (store *MemoryStore) UpdatePriceTick(priceTick domain.PriceTick) error {
	if store.latestPrices == nil {
		return errors.New("price tick does not exist")
	}
	if _, exists := store.latestPrices[priceTick.Symbol]; !exists {
		return errors.New("price tick does not exist")
	}

	store.latestPrices[priceTick.Symbol] = priceTick
	return nil
}

func (store *MemoryStore) GetPriceTick(symbol string) (domain.PriceTick, error) {
	if store.latestPrices == nil {
		return domain.PriceTick{}, errors.New("price tick does not exist")
	}

	priceTick, exists := store.latestPrices[symbol]
	if !exists {
		return domain.PriceTick{}, errors.New("price tick does not exist")
	}

	return priceTick, nil
}

func (store *MemoryStore) InsertRisk(risk domain.PositionRisk) error {
	if store.risks == nil {
		store.risks = make(map[string]map[string]domain.PositionRisk)
	}
	if store.risks[risk.PortfolioID] == nil {
		store.risks[risk.PortfolioID] = make(map[string]domain.PositionRisk)
	}
	if _, exists := store.risks[risk.PortfolioID][risk.Symbol]; exists {
		return errors.New("risk already exists")
	}

	store.risks[risk.PortfolioID][risk.Symbol] = risk
	return nil
}

func (store *MemoryStore) UpdateRisk(risk domain.PositionRisk) error {
	if store.risks == nil || store.risks[risk.PortfolioID] == nil {
		return errors.New("risk does not exist")
	}
	if _, exists := store.risks[risk.PortfolioID][risk.Symbol]; !exists {
		return errors.New("risk does not exist")
	}

	store.risks[risk.PortfolioID][risk.Symbol] = risk
	return nil
}

func (store *MemoryStore) GetRisk(portfolioID, symbol string) (domain.PositionRisk, error) {
	if store.risks == nil || store.risks[portfolioID] == nil {
		return domain.PositionRisk{}, errors.New("risk does not exist")
	}

	risk, exists := store.risks[portfolioID][symbol]
	if !exists {
		return domain.PositionRisk{}, errors.New("risk does not exist")
	}

	return risk, nil
}
