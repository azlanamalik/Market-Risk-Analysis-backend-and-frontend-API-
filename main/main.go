package main

import (
	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/cmd/simulator"
	"context"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	listOfDat := [] string {"AAPL","MSFT","GOOGL","AMZN","TSLA"}
	simulator.Run(ctx,listOfDat)
}
