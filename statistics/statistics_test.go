package statistics

import (
	"testing"
	"time"
)

const (
	statPeriod = 500 * time.Millisecond
)

func TestStatistics(t *testing.T) {

	statistics := NewStatistics(statPeriod)

	//TODO: implement this test to test the statistics middleware (test the value of its counter)
	//TODO: you will need time.Sleep(statPeriod) because of the statistics ticker

	go func() {
		statistics.PlusOne()
	}()

	time.Sleep(statPeriod / 2)

	if statistics.counter != 1 {
		t.Error("Statistics counter should be 1. Actual value is ", statistics.counter)
	}

	time.Sleep(statPeriod)

	if statistics.counter != 0 {
		t.Error("Statistics counter should be 0. Actual value is ", statistics.counter)
	}
}
