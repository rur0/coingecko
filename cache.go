package coingecko

import "sync"

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

func (tp TokenPricesCache) GetPrice(contract string, currency string) (float64, bool) {
	tp.RLock()
	defer tp.RUnlock()
	contractPrices, ok := tp.prices[contract]
	if !ok {
		return 0, false
	}

	price, ok := contractPrices[currency]
	if !ok {
		return 0, false
	}

	return price, ok
}

func (tp *TokenPricesCache) SetPrices(prices TokenPrices) {
	tp.Lock()
	tp.prices = prices
	tp.Unlock()
}

// TokenPricesCache to use globally
var TokenPricesCached = TokenPricesCache{}

// UpdateTokenPriceCache updates the token price cache to use globally
func UpdateTokenPriceCache(contractAddrs []string, currencies []string) error {
	var err error
	prices, err := GetTokenPrices(contractAddrs, currencies)
	if err != nil {
		return err
	}
	TokenPricesCached.SetPrices(prices)
	return nil
}
