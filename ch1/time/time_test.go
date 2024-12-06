package time_test

import (
	"testing"
	"time"
)

func TestTicker(t *testing.T) {
	ticker := time.NewTicker(time.Second)
	var count = 1
	for tc := range ticker.C {
		t.Logf("current ticker: %#v\n", tc)
		count++
		if count > 3 {
			ticker.Stop()
			return
		}
	}
}

func TestTruncate(t *testing.T) {
	d := time.Minute * 5
	nexttime:= time.Now().Truncate(d).Add(d)
	t.Logf("current time: %#v, delay time: %#v\n", nexttime, time.Until(nexttime))
}
