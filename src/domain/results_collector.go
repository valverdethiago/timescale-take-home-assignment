package domain

import (
	"fmt"
	"math"
)

type ResultsCollector struct {
	state                chan []TaskExecutionLogger
	results              []TaskExecutionLogger
	minimumQueryTime     int64
	averageQueryTime     float64
	maximumQueryTime     int64
	globalExecutionMeter TaskExecutionLogger
}

func NewResultsCollector(state chan []TaskExecutionLogger) ResultsCollector {
	return ResultsCollector{
		state:   state,
		results: []TaskExecutionLogger{},
	}

}

func (rc *ResultsCollector) SetGlobalExecutionResult(result TaskExecutionLogger) {
	rc.globalExecutionMeter = result
}

func (rc *ResultsCollector) CollectResults() {
	for {
		workerResults := <-rc.state
		rc.appendWorkerResults(workerResults)
	}
}

func (rc *ResultsCollector) appendWorkerResults(results []TaskExecutionLogger) {
	for _, result := range results {
		rc.results = append(rc.results, result)
	}
}

func (rc *ResultsCollector) PrintReport() {
	rc.calculateMetrics()
	fmt.Printf("Global execution time: %d ms\n", rc.globalExecutionMeter.TimeElapsed.Milliseconds())
	fmt.Printf("Minimum query time: %d ms\n", rc.minimumQueryTime)
	fmt.Printf("Maximum query time: %d ms\n", rc.maximumQueryTime)
	fmt.Printf("Average query time: %f ms\n", rc.averageQueryTime)
}

func (rc *ResultsCollector) calculateMetrics() {
	rc.minimumQueryTime = math.MaxInt64
	rc.maximumQueryTime = math.MinInt64
	var totalExecutionTime int64 = 0
	for _, meter := range rc.results {
		totalExecutionTime += meter.TimeElapsed.Milliseconds()
		if meter.TimeElapsed.Milliseconds() > rc.maximumQueryTime {
			rc.maximumQueryTime = meter.TimeElapsed.Milliseconds()
		}
		if meter.TimeElapsed.Milliseconds() < rc.minimumQueryTime {
			rc.minimumQueryTime = meter.TimeElapsed.Milliseconds()
		}
	}
	rc.averageQueryTime = float64(totalExecutionTime) / float64(len(rc.results))

}
