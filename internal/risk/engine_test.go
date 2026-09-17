package risk

import (
	"testing"
	"time"

	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
)

func TestRiskBoundaries(t *testing.T) {
	validTick := domain.PriceTick{
		EventID:    "event-1",
		Symbol:     "GOLD",
		Bid:        99,
		Ask:        101,
		ObservedAt: time.Unix(1, 0),
	}
	validPosition := domain.Position{
		PortfolioID:  "portfolio-1",
		Symbol:       "GOLD",
		Quantity:     10,
		AveragePrice: 100,
	}

	tests := []struct {
		name     string
		tick     domain.PriceTick
		position domain.Position
		want     domain.PositionRisk
	}{
		{name: "equal bid and ask", tick: func() domain.PriceTick { tick := validTick; tick.Bid = 100; tick.Ask = 100; return tick }(), position: validPosition, want: domain.PositionRisk{PortfolioID: "portfolio-1", Symbol: "GOLD", MarketPrice: 100, MarketValue: 1000, UnrealizedPnL: 0}},
		{name: "spread uses midpoint", tick: validTick, position: validPosition, want: domain.PositionRisk{PortfolioID: "portfolio-1", Symbol: "GOLD", MarketPrice: 100, MarketValue: 1000, UnrealizedPnL: 0}},
		{name: "long profit", tick: domain.PriceTick{EventID: "event-1", Symbol: "GOLD", Bid: 109, Ask: 111, ObservedAt: time.Unix(1, 0)}, position: validPosition, want: domain.PositionRisk{PortfolioID: "portfolio-1", Symbol: "GOLD", MarketPrice: 110, MarketValue: 1100, UnrealizedPnL: 100}},
		{name: "long loss", tick: domain.PriceTick{EventID: "event-1", Symbol: "GOLD", Bid: 89, Ask: 91, ObservedAt: time.Unix(1, 0)}, position: validPosition, want: domain.PositionRisk{PortfolioID: "portfolio-1", Symbol: "GOLD", MarketPrice: 90, MarketValue: 900, UnrealizedPnL: -100}},
		{name: "short profit", tick: domain.PriceTick{EventID: "event-1", Symbol: "GOLD", Bid: 89, Ask: 91, ObservedAt: time.Unix(1, 0)}, position: func() domain.Position { position := validPosition; position.Quantity = -10; return position }(), want: domain.PositionRisk{PortfolioID: "portfolio-1", Symbol: "GOLD", MarketPrice: 90, MarketValue: -900, UnrealizedPnL: 100}},
		{name: "short loss", tick: domain.PriceTick{EventID: "event-1", Symbol: "GOLD", Bid: 109, Ask: 111, ObservedAt: time.Unix(1, 0)}, position: func() domain.Position { position := validPosition; position.Quantity = -10; return position }(), want: domain.PositionRisk{PortfolioID: "portfolio-1", Symbol: "GOLD", MarketPrice: 110, MarketValue: -1100, UnrealizedPnL: -100}},
		{name: "flat price", tick: validTick, position: validPosition, want: domain.PositionRisk{PortfolioID: "portfolio-1", Symbol: "GOLD", MarketPrice: 100, MarketValue: 1000, UnrealizedPnL: 0}},
		{name: "invalid tick", tick: func() domain.PriceTick { tick := validTick; tick.Bid = 0; return tick }(), position: validPosition, want: domain.PositionRisk{PortfolioID: "error loading", Symbol: "error loading", MarketPrice: 1, MarketValue: 1, UnrealizedPnL: 1}},
		{name: "symbol mismatch", tick: func() domain.PriceTick { tick := validTick; tick.Symbol = "SILVER"; return tick }(), position: validPosition, want: domain.PositionRisk{PortfolioID: "error loading", Symbol: "error loading", MarketPrice: 1, MarketValue: 1, UnrealizedPnL: 1}},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			got := risk(test.tick, test.position)
			if got != test.want {
				t.Fatalf("risk() = %+v, want %+v", got, test.want)
			}
		})
	}
}
