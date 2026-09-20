package marketdata

import (
	"fmt"
	"testing"
	"time"
	"context"
)


func TestSimulator(t *testing.T){
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	simulator ,err := NewSimulator(
		map[string]float64{
		"AAPL":  230.50,
		"MSFT":  515.20,
		"GOOGL": 252.80,
		"AMZN":  220.40,
		"TSLA":  430.10,
	},
	1 * time.Second,
	map[string]float64{
		"AAPL":  0.10,
		"MSFT":  0.15,
		"GOOGL": 0.08,
		"AMZN":  0.12,
		"TSLA":  0.25,
	},
	)
	if (err != nil){
		t.Fatalf("issue with initialising the simulator")
	}
	symbol_arr := [] string {"AAPL","MSFT","GOOGL"}
	ticks_chan,_:= simulator.Stream(ctx,symbol_arr)
	for {
		tick := <-ticks_chan
		fmt.Println(tick)
	}
}
//gives a valid output 
// now we can just worry about sending the data as it is slightly out of order but that can be done via sinks not surfice level code
//we dont need to worry about changing anything here as normal for data to be out of order