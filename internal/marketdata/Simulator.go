package marketdata

//this is going to simulate the data at first before we actually impliment

import (
	"context"
	"errors"
	"sync"
	"sync/atomic"
	"time"

	"github.com/azlanamalik/Market-Risk-Analysis-backend-and-frontend-API-/internal/domain"
)

type Simulator struct {
	initialPrices map[string]float64//create multiple at different prices
	interval      time.Duration//how often do we produce a tick
	nextID        atomic.Uint64//counts to create unique ID
	spread 		map[string]float64 // dictates the volatility of each symbol MUST BE THE SAME AS Initial prices
}

func NewSimulator(
	initialPrice map[string]float64,
	interval      time.Duration,
	spread 		map[string]float64) (*Simulator, error){
	//add error checks later this isnt vital right now as we are SIMULATING data
	return &Simulator{initialPrices: initialPrice , interval: interval, spread: spread} , nil
}

/*
streams data in intervals specified when creating func
param: ctx : if the program stops we need to stop the streaming
symbols: the specified symbols to stream (useful if we want to stream on seperate go routines for different symbols)

output: would like to output domain.ticks
*/
func (Simulator *Simulator) Stream(
	ctx context.Context,
	symbols []string) domain.PriceTick{
		ticks := make(chan domain.PriceTick,len(symbols) * 2) // allow for 2 ticks
		errors := make(chan error , len(symbols))
		var workers sync.WaitGroup
		for index,symbol := range symbols{
			
		}


	

}
