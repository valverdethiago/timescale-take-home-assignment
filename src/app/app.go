package app

import (
	"github.com/valverdethiago/timescale-take-home-assignment/config"
	"github.com/valverdethiago/timescale-take-home-assignment/domain"
	"github.com/valverdethiago/timescale-take-home-assignment/util"
	"log"
	"sync"
)

type App struct {
	WorkersNumber int
	FilePath      string
	Config        config.AppConfig
	DbConnector   domain.DbConnector
	waitGroup     sync.WaitGroup
	workers       []domain.Worker
}

func NewApp(workersNumber int, filePath string, config config.AppConfig, dbConnector domain.DbConnector) App {
	return App{
		WorkersNumber: workersNumber,
		FilePath:      filePath,
		Config:        config,
		DbConnector:   dbConnector,
	}
}

func (app *App) Process() domain.TaskExecutionMeter {
	hostIntervalMap, err := app.readFile()
	if err != nil {
		log.Fatal(err)
	}
	app.createWorkersPool(*hostIntervalMap)
	generalTaskMeter := domain.NewTaskExecutionMeter()
	generalTaskMeter.Start()
	app.startProcessing()
	generalTaskMeter.End()
	return generalTaskMeter
}

func (app *App) createWorkersPool(hostIntervalMap map[string][]domain.Interval) {
	totalWorkers := app.WorkersNumber
	totalHosts := len(hostIntervalMap)
	if totalWorkers > totalHosts {
		totalWorkers = totalHosts
	}
	for i := 0; i < totalWorkers; i++ {
		app.workers = append(app.workers, domain.NewWorker(i+1, app.DbConnector))
	}
	executionQueue := domain.NewQueue(util.GetKeysFromMap(hostIntervalMap))
	currentWorker := 0
	for {
		if currentWorker >= totalWorkers {
			currentWorker = 0
		}
		if executionQueue.IsEmpty() {
			break
		}
		hostname := executionQueue.Dequeue()
		app.workers[currentWorker].HostIntervalMap[hostname] = hostIntervalMap[hostname]
		currentWorker++
	}
}

func (app *App) readFile() (*map[string][]domain.Interval, error) {
	fileReader := util.FileReader{
		FilePath: app.FilePath,
	}
	err := fileReader.Read()
	if err != nil {
		return nil, err
	}
	return fileReader.Result, nil
}

func (app *App) startProcessing() {
	var waitGroup sync.WaitGroup
	waitGroup.Add(len(app.workers))
	for _, worker := range app.workers {
		go worker.Process(&waitGroup)
	}
	waitGroup.Wait()
}
