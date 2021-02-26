package coingecko

import (
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/common"
	"github.com/rur0/decimal"
)

func TestBigDecimal(t *testing.T) {
	fmt.Println(decimal.NewFromFloat(0.7496383078861153).Mul(decimal.NewFromInt(100)))
}

func TestMarketChart(t *testing.T) {
	chart, err := GetMarketChart("zenfuse", "usd", "1")
	if err != nil {
		t.Fatal(err)
	}
	for _, price := range chart.Prices {
		fmt.Println(price)
	}
}

func TestCache(t *testing.T) {
	tkAddr := common.HexToAddress("0xb1e9157c2fdcc5a856c8da8b2d89b6c32b3c1229")

	err := UpdateTokenPriceCache([]common.Address{tkAddr}, []string{"eth"})
	if err != nil {
		t.Fatal(err)
	}
	price, ok := TokenPricesCached.GetPrice(tkAddr, "eth")
	if !ok {
		t.Fatal("price not found")
	}

	mcap, ok := TokenPricesCached.GetMcap(tkAddr, "eth")
	if !ok {
		t.Fatal("mcap not found")
	}

	fmt.Println(price, mcap)
}
