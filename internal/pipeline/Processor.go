package pipeline

import (
	"context"
	"fmt"
	"log/slog"
	"strconv"

	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/store"
)

/*
putting it all together
processor.Process(tick) signiture
	↓ to impliment
validate tick
	↓
store.SavePrice(tick)
	↓
store.GetPositionsBySymbol(tick.Symbol)
	↓
for each position
	↓
risk.Calculate(tick, position)
	↓
store.SaveRisk(result)
*/
//key part store locally then process then add to store else we wont validate before moving it to the store
func Processor(ctx context.Context, tick domain.PriceTick, store *store.MemoryStore) {
	valid := tick.Validator()
	if !valid {
		slog.Warn("issue with processor: the tick is not valid")
	}
	if err := store.InsertPriceTick(ctx, tick); err != nil {
		slog.Warn("issue with storing the pricetick")
	}
	dataPosBySymb, errGetPosBySymbol := store.GetPositionsBySymbol(ctx, tick.Symbol)
	if errGetPosBySymbol != nil {
		slog.Warn("issue with storing the pricetick")
	}
	for index, position := range dataPosBySymb {
		print("here\n")//best debugging tool :)
		fmt.Println("position", strconv.Itoa(index), ":", position)
	}
	
}
