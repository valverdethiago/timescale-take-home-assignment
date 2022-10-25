package domain

import (
	"fmt"
	"log"
	"sync"
)

type Task struct {
	Hostname string
	Interval Interval
}

type Worker struct {
	ID              string
	HostIntervalMap map[string][]Interval
	dbConnector     DbConnector
}

func NewWorker(id int, dbConnector DbConnector) Worker {
	return Worker{
		ID:              fmt.Sprintf("Worker-%d", id),
		HostIntervalMap: make(map[string][]Interval),
		dbConnector:     dbConnector,
	}
}

func (worker *Worker) Process(waitGroup *sync.WaitGroup, state chan<- []TaskExecutionLogger) {
	var results []TaskExecutionLogger
	for hostname := range worker.HostIntervalMap {
		for _, interval := range worker.HostIntervalMap[hostname] {
			intervalTaskMeter := NewTaskExecutionMeter()
			intervalTaskMeter.Start()
			err := worker.dbConnector.ExecuteQuery(hostname, interval.StartTime, interval.EndTime)
			if err != nil {
				log.Println(err)
			}
			intervalTaskMeter.End()
			results = append(results, intervalTaskMeter)
		}
	}
	state <- results
	waitGroup.Done()
}
