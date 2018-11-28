package web

import (
	"github.com/Sfeir/golang-200/statistics"
	"net/http"
	"time"
)

// StatisticsMiddleware is the middleware to record request statistics
type StatisticsMiddleware struct {
	// TODO use a pointer to a statistics struct 'Stat' to build the middleware
	stats *statistics.Statistics
}

// NewStatisticsMiddleware creates a new statistics middleware
func NewStatisticsMiddleware(duration time.Duration) *StatisticsMiddleware {
	// TODO return a newly initialized pointer to a StatisticsMiddleware

	return &StatisticsMiddleware{
		statistics.NewStatistics(duration),
	}
}

func (sm *StatisticsMiddleware) ServeHTTP(rw http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	// TODO count the hit on the API and call the next middleware
	sm.stats.PlusOne()
	next(rw, r)
}
