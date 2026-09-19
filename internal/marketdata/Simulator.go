package marketdata

//this is going to simulate the data at first before we actually impliment

import (
	"context"
	"fmt"
	"math/rand"
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
func (simulator *Simulator) Stream(
	ctx context.Context,//only worry about this in loops most code that is run without loops are too fast to worry about context
	symbols []string) (<-chan domain.PriceTick, <-chan error){
		ticks := make(chan domain.PriceTick,len(symbols) * 2) // allow for 2 ticks
		errorsChannel := make(chan error , len(symbols))
		var workers sync.WaitGroup
		for index, symbol := range symbols {
			startingPrice, exists := simulator.initialPrices[symbol]
			if !exists {
				errorsChannel <- fmt.Errorf("no starting price configured for %s", symbol)
				continue
			}

			workers.Add(1)
			go func(symbol string, startingPrice float64, seed int64) {
				defer workers.Done()
				simulator.streamSymbol(ctx, ticks, symbol, startingPrice, seed)
			}(symbol, startingPrice, time.Now().UnixNano()+int64(index))
		}

		go func() {
			workers.Wait()
			close(ticks)
			close(errorsChannel)
		}()

		return ticks, errorsChannel//return output channels
		//finish first create streamSymbol
		}



/*
we need to stream for each symbol: it goes like this stream -> invokes stream symbol and that provides the data for each symbol

input:
context - incase we decide to terminate we need to close as this is a constantly running program
out - the channel we write to for PriceTick
symbol - what symbol are we streaming
price - the original price of the symbol
seed - new add as it makes it easier to test data

output:
out : only thing that is returned but it is a channel so we dont neeed to worry about that :)


*/
func (simulator *Simulator) streamSymbol(
	ctx context.Context,
	out chan<- domain.PriceTick,
	symbol string,
	price float64,
	seed int64,
) {
	random := rand.New(rand.NewSource(seed))
	//logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	ticker := time.NewTicker(simulator.interval)//interval event creation
	defer ticker.Stop()//ensures we stop
	for {//while true so it doesnt stop
		select {
		case <-ctx.Done()://program stopped context
			return
		case timestamp := <-ticker.C://write the timestamp 
			price *= 1 + ((random.Float64() - 0.5) / 1000)
			spread := price * 0.0001
			tick := domain.PriceTick{
				EventID: fmt.Sprintf("tick-%d", simulator.nextID.Add(1)),
				Symbol:  symbol,
				Bid:     price - spread/2,
				Ask:     price + spread/2,
				ObservedAt:    timestamp.UTC(),
			}
			//logger.Info("hello this is azlan", "tick output",tick)
			select {
			case <-ctx.Done()://if we stopped the program during the actual run
				return
			case out <- tick://outputs
			}
		}
	}
}