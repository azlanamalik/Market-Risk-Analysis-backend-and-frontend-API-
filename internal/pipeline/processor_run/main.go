package main

import (
	"context"
	"fmt"
	"time"

	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/pipeline"
	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/store"
)

func main() {
	memoryStore := store.MemoryStore{}
	ctx := context.Background()

	memoryStore.InsertPosition(ctx, domain.Position{
		PortfolioID:  "growth-portfolio",
		Symbol:       "AAPL",
		Quantity:     100,
		AveragePrice: 225.50,
	})
	memoryStore.InsertPosition(ctx, domain.Position{
		PortfolioID:  "income-portfolio",
		Symbol:       "AAPL",
		Quantity:     40,
		AveragePrice: 228.25,
	})

	tick := domain.PriceTick{
		EventID:    "quick-check-1",
		Symbol:     "AAPL",
		Bid:        229.90,
		Ask:        230.10,
		ObservedAt: time.Now(),
	}

	pipeline.Processor(ctx, tick, &memoryStore)

	storedTick, _ := memoryStore.GetPriceTick(ctx, "AAPL")
	fmt.Printf("processor ran: %+v\n", storedTick)
}
