package domain

import (
	"log"
	"sync"
)

type Task struct {
	Hostname string
	Interval Interval
}

type Worker struct {
	ID              int
	HostIntervalMap map[string][]Interval
	dbConnector     DbConnector
}

func NewWorker(id int, dbConnector DbConnector) Worker {
	return Worker{
		ID:              id,
		HostIntervalMap: make(map[string][]Interval),
		dbConnector:     dbConnector,
	}
}

func (worker *Worker) Process(waitGroup *sync.WaitGroup) {
	workerTaskMeter := NewTaskExecutionMeter()
	workerTaskMeter.Start()
	for hostname := range worker.HostIntervalMap {
		for _, interval := range worker.HostIntervalMap[hostname] {
			intervalTaskMeter := NewTaskExecutionMeter()
			intervalTaskMeter.Start()
			worker.dbConnector.ExecuteQuery(hostname, interval.StartTime, interval.EndTime)
			intervalTaskMeter.End()
			log.Println("intervalTaskMeter: ", intervalTaskMeter)
		}
	}
	workerTaskMeter.End()
	log.Println("workerTaskMeter: ", workerTaskMeter)
	waitGroup.Done()
}
