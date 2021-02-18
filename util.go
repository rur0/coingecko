package coingecko

import (
	"strconv"
	"time"
)

var (
	millisInSecond = time.Second / time.Millisecond
	nsInSecond     = time.Second / time.Nanosecond
)

func MsToTime(mss string) (time.Time, error) {
	ms, err := strconv.ParseInt(mss, 10, 0)
	if err != nil {
		return time.Time{}, err
	}
	return time.Unix(ms/1000, 0), nil
}
