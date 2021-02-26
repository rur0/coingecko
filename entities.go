package coingecko

import (
	"strings"
	"time"

	"github.com/rur0/decimal"
)

type MarketChart struct {
	Prices       []timedPrice `json:"prices"`
	MarketCaps   []timedPrice `json:"market_caps"`
	TotalVolumes []timedPrice `json:"total_volumes"`
}

type timedPrice struct {
	Time  time.Time
	Value decimal.Decimal
}

func (m *timedPrice) UnmarshalJSON(data []byte) error {
	var err error
	vals := strings.Split(strings.Trim(string(data), "[]"), ",")

	m.Time, err = MsToTime(vals[0])
	if err != nil {
		return err
	}

	m.Value, err = decimal.NewFromString(vals[1])
	if err != nil {
		return err
	}

	return nil
}

// 7906058.858537895
// 0.7496383078861153
// 54158413.15214113
