package coingecko

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
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
