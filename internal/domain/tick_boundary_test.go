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
			if got := test.tick.validator(); got != test.valid {
				t.Fatalf("validator() = %v, want %v", got, test.valid)
			}
		})
	}
}
