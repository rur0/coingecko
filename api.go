package coingecko

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

// TokenPrices example:
// currently only supports ethereum
/*
{
  "0x000...": {
    "usd": 0.01259181
  }
}
*/
type TokenPrices map[string]map[string]float64

var Client = &http.Client{
	Timeout: time.Second * 5,
}

func GetTokenPrices(contractAddrs []string, currencies []string) (TokenPrices, error) {
	q := url.Values{}
	q.Add("contract_addresses", strings.Join(contractAddrs, ","))
	q.Add("vs_currencies", strings.Join(currencies, ","))
	resp, err := Client.Get("https://api.coingecko.com/api/v3/simple/token_price/ethereum?" + q.Encode())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("error: " + string(bBytes))
	}

	tkPrices := &TokenPrices{}
	err = json.NewDecoder(resp.Body).Decode(tkPrices)
	if err != nil {
		return nil, err
	}

	return *tkPrices, nil
}

//Prices holds cryptocurrency prices to fiat
/*
example:
{
  "ethereum": {
    "gbp": 892.65,
    "usd": 1215.77,
    "eur": 985.75
  }
}
*/
type Prices map[string]map[string]float64

func GetPrices(coins []string, currencies []string) (Prices, error) {
	q := url.Values{}
	q.Add("ids", strings.Join(coins, ","))
	q.Add("vs_currencies", strings.Join(currencies, ","))
	resp, err := Client.Get("https://api.coingecko.com/api/v3/simple/price?" + q.Encode())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		bBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("error: " + string(bBytes))
	}

	prices := &Prices{}
	err = json.NewDecoder(resp.Body).Decode(prices)
	if err != nil {
		return nil, err
	}

	return *prices, nil
}

func GetMarketChart(coin, vsCurrency, days string) (*MarketChart, error) {
	q := url.Values{}
	q.Add("vs_currency", vsCurrency)
	q.Add("days", days)

	resp, err := Client.Get("https://api.coingecko.com/api/v3/coins/" + strings.ToLower(coin) + "/market_chart?" + q.Encode())
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		bBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New("error: " + string(bBytes))
	}

	marketChart := &MarketChart{}
	err = json.NewDecoder(resp.Body).Decode(marketChart)
	if err != nil {
		return nil, err
	}

	return marketChart, nil
}
