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
		name                                      string
		tick                                      domain.PriceTick
		position                                  domain.Position
		wantMarketPrice, wantMarketValue, wantPnL float64
	}{
		{name: "equal bid and ask", tick: func() domain.PriceTick { tick := validTick; tick.Bid = 100; tick.Ask = 100; return tick }(), position: validPosition, wantMarketPrice: 100, wantMarketValue: 1000, wantPnL: 0},
		{name: "spread uses midpoint", tick: validTick, position: validPosition, wantMarketPrice: 100, wantMarketValue: 1000, wantPnL: 0},
		{name: "long profit", tick: domain.PriceTick{EventID: "event-1", Symbol: "GOLD", Bid: 109, Ask: 111, ObservedAt: time.Unix(1, 0)}, position: validPosition, wantMarketPrice: 110, wantMarketValue: 1100, wantPnL: 100},
		{name: "long loss", tick: domain.PriceTick{EventID: "event-1", Symbol: "GOLD", Bid: 89, Ask: 91, ObservedAt: time.Unix(1, 0)}, position: validPosition, wantMarketPrice: 90, wantMarketValue: 900, wantPnL: -100},
		{name: "short profit", tick: domain.PriceTick{EventID: "event-1", Symbol: "GOLD", Bid: 89, Ask: 91, ObservedAt: time.Unix(1, 0)}, position: func() domain.Position { position := validPosition; position.Quantity = -10; return position }(), wantMarketPrice: 90, wantMarketValue: -900, wantPnL: 100},
		{name: "short loss", tick: domain.PriceTick{EventID: "event-1", Symbol: "GOLD", Bid: 109, Ask: 111, ObservedAt: time.Unix(1, 0)}, position: func() domain.Position { position := validPosition; position.Quantity = -10; return position }(), wantMarketPrice: 110, wantMarketValue: -1100, wantPnL: -100},
		{name: "flat price", tick: validTick, position: validPosition, wantMarketPrice: 100, wantMarketValue: 1000, wantPnL: 0},
		{name: "invalid tick", tick: func() domain.PriceTick { tick := validTick; tick.Bid = 0; return tick }(), position: validPosition, wantMarketPrice: 0, wantMarketValue: 0, wantPnL: 0},
		{name: "symbol mismatch", tick: func() domain.PriceTick { tick := validTick; tick.Symbol = "SILVER"; return tick }(), position: validPosition, wantMarketPrice: 0, wantMarketValue: 0, wantPnL: 0},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			marketPrice, marketValue, pnl := risk(test.tick, test.position)
			if marketPrice != test.wantMarketPrice || marketValue != test.wantMarketValue || pnl != test.wantPnL {
				t.Fatalf("risk() = (%v, %v, %v), want (%v, %v, %v)", marketPrice, marketValue, pnl, test.wantMarketPrice, test.wantMarketValue, test.wantPnL)
			}
		})
	}
}
