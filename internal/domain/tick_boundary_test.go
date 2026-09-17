package domain

import (
	"testing"
	"time"
)

func TestPriceTickValidatorBoundaries(t *testing.T) {
	validTick := PriceTick{
		EventID:    "event-1",
		Symbol:     "GOLD",
		Bid:        100,
		Ask:        100,
		ObservedAt: time.Unix(1, 0),
	}

	tests := []struct {
		name  string
		tick  PriceTick
		valid bool
	}{
		{name: "equal bid and ask is valid", tick: validTick, valid: true},
		{name: "zero bid is invalid", tick: func() PriceTick { tick := validTick; tick.Bid = 0; return tick }(), valid: false},
		{name: "zero ask is invalid", tick: func() PriceTick { tick := validTick; tick.Ask = 0; return tick }(), valid: false},
		{name: "ask below bid is invalid", tick: func() PriceTick { tick := validTick; tick.Ask = 99; return tick }(), valid: false},
		{name: "missing event id is invalid", tick: func() PriceTick { tick := validTick; tick.EventID = ""; return tick }(), valid: false},
		{name: "missing symbol is invalid", tick: func() PriceTick { tick := validTick; tick.Symbol = ""; return tick }(), valid: false},
		{name: "zero observation time is invalid", tick: func() PriceTick { tick := validTick; tick.ObservedAt = time.Time{}; return tick }(), valid: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.tick.Validator(); got != test.valid {
				t.Fatalf("validator() = %v, want %v", got, test.valid)
			}
		})
	}
}

func TestPositionValidatorBoundaries(t *testing.T) {
	validPosition := Position{
		PortfolioID:  "portfolio-1",
		Symbol:       "GOLD",
		Quantity:     10,
		AveragePrice: 100,
	}

	tests := []struct {
		name     string
		position Position
		valid    bool
	}{
		{name: "valid long position passes", position: validPosition, valid: true},
		{name: "valid short position passes", position: func() Position { position := validPosition; position.Quantity = -10; return position }(), valid: true},
		{name: "zero quantity fails", position: func() Position { position := validPosition; position.Quantity = 0; return position }(), valid: false},
		{name: "missing portfolio fails", position: func() Position { position := validPosition; position.PortfolioID = ""; return position }(), valid: false},
		{name: "missing symbol fails", position: func() Position { position := validPosition; position.Symbol = ""; return position }(), valid: false},
		{name: "zero average price fails", position: func() Position { position := validPosition; position.AveragePrice = 0; return position }(), valid: false},
		{name: "negative average price fails", position: func() Position { position := validPosition; position.AveragePrice = -1; return position }(), valid: false},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := test.position.Validator(); got != test.valid {
				t.Fatalf("validator() = %v, want %v", got, test.valid)
			}
		})
	}
}
