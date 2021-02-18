package coingecko

// PricesCache to use globally
var PricesCache = Prices{}

// UpdatePriceCache updates the price cache to use globally
func UpdatePriceCache(coins []string, currencies []string) error {
	var err error
	PricesCache, err = GetPrices(coins, currencies)
	return err
}

// TokenPricesCache to use globally
var TokenPricesCache = TokenPrices{}

// UpdateTokenPriceCache updates the token price cache to use globally
func UpdateTokenPriceCache(contractAddrs []string, currencies []string) error {
	var err error
	TokenPricesCache, err = GetTokenPrices(contractAddrs, currencies)
	return err
}
