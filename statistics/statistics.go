package statistics

import (
	"fmt"
	logger "github.com/sirupsen/logrus"
	"time"
)

const (
	statisticsChannelSize = 1000
)

// TODO define a Statistics struct with an uint8 chan, an uint32 counter, a start time and logging period duration
// Statistics is the worker to persist the request statistics
type Statistics struct {
	ch                    chan uint8
	counter               uint32
	startTime             time.Time
	loggingPeriodDuration time.Duration
}

// NewStatistics creates a new statistics structure and launches its worker routine
func NewStatistics(loggingPeriod time.Duration) *Statistics {
	// TODO build a new Statistics object with a sized channel, initialized counter and start time
	// TODO and logging period as param

	ch := make(chan uint8, statisticsChannelSize)

	stats := &Statistics{
		ch,
		0,
		time.Now(),
		loggingPeriod,
	}
	// TODO launch the run in a separate Go routine in background
	go stats.run()

	// TODO return the initialized and started object
	return stats
}

// PlusOne is used to send a statistics hit increment
func (sw *Statistics) PlusOne() {
	// TODO push a hit in the statistics channel
	select {
	case sw.ch <- 1:
	default:
		logger.Warn("message PlusOne not sent")
	}
}

func (sw *Statistics) run() {
	// TODO build a new time Ticker from the logging period
	ticker := time.Tick(sw.loggingPeriodDuration)

	// TODO build a infinite loop and the channel selection inside
	for {
		select {
		case msg := <-sw.ch:
			sw.counter += uint32(msg)
		case tick := <-ticker:
			tick.Clock()
			rate := float64(sw.counter) / sw.loggingPeriodDuration.Seconds()
			logger.Info("apps has received ", fmt.Sprintf("%.2f", rate), " per second (period: ", sw.loggingPeriodDuration, ")")
			sw.counter = 0
		}
	}
	// TODO build a first select case from the statistics channel
	// TODO add the hit count to the counter and log it as debug level

	// TODO build a second case on the time Ticker chan
	// TODO retrieve the elapsed time since start
	// TODO log the hit/sec rate
	// TODO reset the counter and the start time
}
