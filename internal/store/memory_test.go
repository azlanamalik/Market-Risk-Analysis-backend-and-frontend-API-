package store

import (
	"context"
	"errors"
	"testing"

	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
)

func testPosition(portfolioID, symbol string, quantity float64) domain.Position {
	return domain.Position{
		PortfolioID:  portfolioID,
		Symbol:       symbol,
		Quantity:     quantity,
		AveragePrice: 100,
	}
}

func TestInsertAndGetPositionBySymbol(t *testing.T) {
	store := &MemoryStore{}//allows to create and query the functions it means make memstore then get the location
	position := testPosition("portfolio-1", "DEMO", 10)

	if err := store.InsertPosition(position); err != nil {
		t.Fatalf("insert position: %v", err)
	}

	positions, err := store.GetPositionsBySymbol(context.Background(), "DEMO")
	if err != nil {
		t.Fatalf("get positions: %v", err)
	}
	if len(positions) != 1 {
		t.Fatalf("expected 1 position, got %d", len(positions))
	}
	if positions[0] != position {
		t.Fatalf("got %+v, want %+v", positions[0], position)
	}
}

func TestUpdatePositionDoesNotCreateDuplicate(t *testing.T) {
	store := &MemoryStore{}
	position := testPosition("portfolio-1", "DEMO", 10)

	if err := store.InsertPosition(position); err != nil {
		t.Fatalf("insert position: %v", err)
	}

	updated := testPosition("portfolio-1", "DEMO", 25)
	if err := store.UpdatePosition(updated); err != nil {
		t.Fatalf("update position: %v", err)
	}

	positions, err := store.GetPositionsBySymbol(context.Background(), "DEMO")
	if err != nil {
		t.Fatalf("get positions: %v", err)
	}
	if len(positions) != 1 {
		t.Fatalf("expected 1 position after update, got %d", len(positions))
	}
	if positions[0].Quantity != 25 {
		t.Fatalf("expected quantity 25, got %v", positions[0].Quantity)
	}
}

func TestSameSymbolCanExistInDifferentPortfolios(t *testing.T) {
	store := &MemoryStore{}

	for _, position := range []domain.Position{
		testPosition("portfolio-1", "DEMO", 10),
		testPosition("portfolio-2", "DEMO", 20),
	} {
		if err := store.InsertPosition(position); err != nil {
			t.Fatalf("insert position: %v", err)
		}
	}

	positions, err := store.GetPositionsBySymbol(context.Background(), "DEMO")
	if err != nil {
		t.Fatalf("get positions: %v", err)
	}
	if len(positions) != 2 {
		t.Fatalf("expected 2 positions, got %d", len(positions))
	}
}

func TestUnknownSymbolReturnsNoPositions(t *testing.T) {
	store := &MemoryStore{}

	positions, err := store.GetPositionsBySymbol(context.Background(), "UNKNOWN")
	if err != nil {
		t.Fatalf("get positions: %v", err)
	}
	if len(positions) != 0 {
		t.Fatalf("expected no positions, got %d", len(positions))
	}
}

func TestChangingReturnedSliceDoesNotChangeStore(t *testing.T) {
	store := &MemoryStore{}
	position := testPosition("portfolio-1", "DEMO", 10)

	if err := store.InsertPosition(position); err != nil {
		t.Fatalf("insert position: %v", err)
	}

	positions, err := store.GetPositionsBySymbol(context.Background(), "DEMO")
	if err != nil {
		t.Fatalf("get positions: %v", err)
	}
	positions[0] = testPosition("changed", "CHANGED", 999)

	stored, err := store.GetPosition("portfolio-1", "DEMO")
	if err != nil {
		t.Fatalf("get stored position: %v", err)
	}
	if stored != position {
		t.Fatalf("store changed unexpectedly: got %+v, want %+v", stored, position)
	}
}

func TestCancelledContextReturnsCancellation(t *testing.T) {
	store := &MemoryStore{}
	ctx, cancel := context.WithCancel(context.Background())
	cancel()

	_, err := store.GetPositionsBySymbol(ctx, "DEMO")
	if !errors.Is(err, context.Canceled) {
		t.Fatalf("expected context.Canceled, got %v", err)
	}
}
