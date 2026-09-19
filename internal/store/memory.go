// this file is for storing stuff in memory
package store

import (
	"context"
	"errors"
	"sync"
	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
)


type MemoryStore struct {
	mu sync.RWMutex//allows for a safe way to access memory on several go routines
	positions    map[string]map[string]domain.Position     //[portID] [ symbol] position . prints the position with that stock
	latestPrices map[string]domain.PriceTick               //[ symbol] domain.PriceTick . prints the lastest price info
	risks        map[string]map[string]domain.PositionRisk //[portID] [symbol] position . prints the LATEST RISK FOR OPRTOLIO AND THE SYMB
}
///we have created this safe for single threads but now we should look to add concurrency 
func (store *MemoryStore) InsertPosition(position domain.Position) error {
	store.mu.Lock()
	defer store.mu.Unlock()
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
	store.mu.Lock()
	defer store.mu.Unlock()
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
	store.mu.RLock()
	defer store.mu.RUnlock()
	if store.positions == nil || store.positions[portfolioID] == nil {
		return domain.Position{}, errors.New("position does not exist")
	}

	position, exists := store.positions[portfolioID][symbol]
	if !exists {
		return domain.Position{}, errors.New("position does not exist")
	}

	return position, nil
}

func (store *MemoryStore) GetPositionsBySymbol(ctx context.Context, symbol string) ([]domain.Position, error) {
	store.mu.RLock()
	defer store.mu.RUnlock()
	if err := ctx.Err(); err != nil {
		return nil, err
	}

	positions := make([]domain.Position, 0)
	for _, portfolioPositions := range store.positions {
		if position, exists := portfolioPositions[symbol]; exists {
			positions = append(positions, position)
		}
	}

	return positions, nil
}

func (store *MemoryStore) InsertPriceTick(priceTick domain.PriceTick) error {
	store.mu.Lock()
	defer store.mu.Unlock()
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
	store.mu.Lock()
	defer store.mu.Unlock()
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
	store.mu.RLock()
	defer store.mu.RUnlock()
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
	store.mu.Lock()
	defer store.mu.Unlock()
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
	store.mu.Lock()
	defer store.mu.Unlock()
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
	store.mu.RLock()
	defer store.mu.RUnlock()
	if store.risks == nil || store.risks[portfolioID] == nil {
		return domain.PositionRisk{}, errors.New("risk does not exist")
	}

	risk, exists := store.risks[portfolioID][symbol]
	if !exists {
		return domain.PositionRisk{}, errors.New("risk does not exist")
	}

	return risk, nil
}
