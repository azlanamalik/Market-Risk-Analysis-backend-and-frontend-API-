package domain

import (
	"log/slog"
	"time"
)

type PriceTick struct {
	EventID    string    `json:"event_id"`
	Symbol     string    `json:"symbol"`
	Bid        float64   `json:"bid"`
	Ask        float64   `json:"ask"`
	ObservedAt time.Time `json:"observed_at"`
} //creates it in form of json when we do REST implimentation

func (PriceTick PriceTick) Midpoint() float64 {
	return (PriceTick.Bid + PriceTick.Ask) / 2
}
func (PriceTick PriceTick) validator() bool {
	logger := slog.Default()
	if PriceTick.EventID == "" {
		logger.Error("invalid price EventID", "event_id", PriceTick.EventID)
		return false
	}
	if PriceTick.Symbol == "" {
		logger.Error("invalid price Symbol", "Symbol", PriceTick.Symbol)
		return false
	}
	if PriceTick.Ask <= 0 {
		logger.Error("invalid price Ask", "Ask", PriceTick.Ask)
		return false
	}
	if PriceTick.Bid <= 0 {
		logger.Error("invalid price Bid", "Bid", PriceTick.Bid)
		return false
	}
	if PriceTick.Ask < PriceTick.Bid {
		logger.Error("invalid price Ask/Bid", "ask", PriceTick.Ask, "bid", PriceTick.Bid)
		return false
	}
	if PriceTick.ObservedAt.IsZero() {
		logger.Error("invalid price Time", "Time", PriceTick.ObservedAt)
		return false
	}
	return true
}
