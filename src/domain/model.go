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

type TaskExecutionLogger struct {
	StartTime   time.Time
	EndTime     time.Time
	TimeElapsed time.Duration
}

func NewTaskExecutionMeter() TaskExecutionLogger {
	return TaskExecutionLogger{}
}

func (taskExecutionMeter *TaskExecutionLogger) Start() {
	taskExecutionMeter.StartTime = time.Now()
}

func (taskExecutionMeter *TaskExecutionLogger) End() {
	taskExecutionMeter.EndTime = time.Now()
	taskExecutionMeter.TimeElapsed = time.Since(taskExecutionMeter.StartTime)
}
