package store

import (
	"context"
	"errors"
	"fmt"
	"sync"
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
	store := &MemoryStore{} //allows to create and query the functions it means make memstore then get the location
	position := testPosition("portfolio-1", "DEMO", 10)

	if err := store.InsertPosition(context.Background(), position); err != nil {
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

	if err := store.InsertPosition(context.Background(), position); err != nil {
		t.Fatalf("insert position: %v", err)
	}

	updated := testPosition("portfolio-1", "DEMO", 25)
	if err := store.UpdatePosition(context.Background(), updated); err != nil {
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
		if err := store.InsertPosition(context.Background(), position); err != nil {
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

	if err := store.InsertPosition(context.Background(), position); err != nil {
		t.Fatalf("insert position: %v", err)
	}

	positions, err := store.GetPositionsBySymbol(context.Background(), "DEMO")
	if err != nil {
		t.Fatalf("get positions: %v", err)
	}
	positions[0] = testPosition("changed", "CHANGED", 999)

	stored, err := store.GetPosition(context.Background(), "portfolio-1", "DEMO")
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
func TestConcurrentReadsFromMemory(t *testing.T) {
	store := &MemoryStore{}
	position := testPosition("portfolio-1", "DEMO", 10)
	if err := store.InsertPosition(context.Background(), position); err != nil {
		t.Fatalf("insert position: %v", err)
	}

	var waitGroup sync.WaitGroup
	stream := make(chan error, 4)
	for range 4 {
		waitGroup.Add(1)
		go func() {
			defer waitGroup.Done()
			_, err := store.GetPosition(context.Background(), "portfolio-1", "DEMO")
			stream <- err
		}()
	}

	waitGroup.Wait()
	for range 4 {
		if err := <-stream; err != nil {
			t.Errorf("concurrent read failed: %v", err)
		}
	}
}

func TestConcurrentReadersAndWriters(t *testing.T) {
	store := &MemoryStore{}
	const positionCount = 100

	// Create records first so every writer can update an existing position.
	for i := 0; i < positionCount; i++ {
		position := testPosition(fmt.Sprintf("portfolio-%d", i), "DEMO", 10)
		if err := store.InsertPosition(context.Background(), position); err != nil {
			t.Fatalf("insert position: %v", err)
		}
	}

	var waitGroup sync.WaitGroup
	errorsFromGoroutines := make(chan error, positionCount*2)

	// Start writers. Each writer updates a different position.
	for i := 0; i < positionCount; i++ {
		waitGroup.Add(1)
		go func(index int) {
			defer waitGroup.Done() //so issue here is that main func will end before the func return we need to
			//defer until it is finished
			position := testPosition(fmt.Sprintf("portfolio-%d", index), "DEMO", 20)
			if err := store.UpdatePosition(context.Background(), position); err != nil {
				errorsFromGoroutines <- err
			}
		}(i)
	}

	// Start readers while the writers are updating the store.
	for i := 0; i < positionCount; i++ {
		waitGroup.Add(1) //always do this
		go func() {
			defer waitGroup.Done() //always do this as main func will end before you can write run it
			if _, err := store.GetPositionsBySymbol(context.Background(), "DEMO"); err != nil {
				errorsFromGoroutines <- err
			}
		}()
	}

	waitGroup.Wait()
	close(errorsFromGoroutines)

	for err := range errorsFromGoroutines {
		t.Errorf("concurrent operation failed: %v", err)
	}
}

//get github copilot to write the rest of the test cases as they are the same just different variables (just to save time)
