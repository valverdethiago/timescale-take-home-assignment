package domain

import "time"

type FileEntry struct {
	Hostname string
	Interval Interval
}

type Interval struct {
	StartTime string
	EndTime   string
}

type QueryExecutionResult struct {
	Hostname string
	Interval Interval
	StartTs  time.Time
	EndTs    time.Time
	Elapsed  time.Duration
}

func NewQueryExecutionResult(hostname string, interval Interval) QueryExecutionResult {
	result := QueryExecutionResult{
		Hostname: hostname,
		Interval: interval,
	}
	result.start()
	return result
}

func (queryExecutionResult *QueryExecutionResult) start() {
	queryExecutionResult.StartTs = time.Now()
}
func (queryExecutionResult *QueryExecutionResult) End() {
	queryExecutionResult.EndTs = time.Now()
	queryExecutionResult.Elapsed = time.Since(queryExecutionResult.StartTs)
}
