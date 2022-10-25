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

type TaskExecutionMeter struct {
	//Hostname    string
	//Interval    Interval
	StartTime   time.Time
	EndTime     time.Time
	TimeElapsed time.Duration
}

func NewTaskExecutionMeter() TaskExecutionMeter {
	return TaskExecutionMeter{}
}

func (taskExecutionMeter *TaskExecutionMeter) Start() {
	taskExecutionMeter.StartTime = time.Now()
}

func (taskExecutionMeter *TaskExecutionMeter) End() {
	taskExecutionMeter.EndTime = time.Now()
	taskExecutionMeter.TimeElapsed = time.Since(taskExecutionMeter.StartTime)
}
