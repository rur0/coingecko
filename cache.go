package coingecko

import (
	"strings"
	"sync"

	"github.com/ethereum/go-ethereum/common"
)

// PricesCache to use globally
var PricesCache = Prices{}

// UpdatePriceCache updates the price cache to use globally
func UpdatePriceCache(coins []string, currencies []string) error {
	var err error
	PricesCache, err = GetPrices(coins, currencies)
	return err
}

type TokenPricesCache struct {
	prices TokenPrices
	sync.RWMutex
}

func (tp TokenPricesCache) GetPrice(token common.Address, currency string) (float64, bool) {
	tp.RLock()
	defer tp.RUnlock()
	contractPrices, ok := tp.prices[token]
	if !ok {
		return 0, false
	}

	price, ok := contractPrices[strings.ToLower(currency)]
	if !ok {
		return 0, false
	}

	return price, true
}

func (tp TokenPricesCache) GetMcap(token common.Address, currency string) (float64, bool) {
	tp.RLock()
	defer tp.RUnlock()
	tkPrices, ok := tp.prices[token]
	if !ok {
		return 0, false
	}

	mcap, ok := tkPrices[strings.ToLower(currency)+"_market_cap"]
	if !ok {
		return 0, false
	}

	return mcap, true
}

func (tp *TokenPricesCache) SetPrices(prices TokenPrices) {
	tp.Lock()
	tp.prices = prices
	tp.Unlock()
}

// TokenPricesCache to use globally
var TokenPricesCached = TokenPricesCache{}

// UpdateTokenPriceCache updates the token price cache to use globally
func UpdateTokenPriceCache(contractAddrs []common.Address, currencies []string) error {
	prices, err := GetTokenPrices(contractAddrs, currencies)
	if err != nil {
		return err
	}
	TokenPricesCached.SetPrices(prices)
	return nil
}
