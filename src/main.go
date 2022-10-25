package main

import (
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/valverdethiago/timescale-take-home-assignment/adapter"
	"github.com/valverdethiago/timescale-take-home-assignment/app"
	"github.com/valverdethiago/timescale-take-home-assignment/config"
)

func main() {

	numberOfWorkers := flag.Int("workers", 10, "Number of workers")
	filePath := flag.String("file", "../query_params.csv", "File Path")

	flag.Parse()

	fmt.Printf("Number of workers: %d\n", *numberOfWorkers)
	fmt.Printf("File: %s\n", *filePath)
	fmt.Println("tail:", flag.Args())

	appConfig := config.LoadEnvConfig()
	dbConnector := adapter.NewTimescaleAdapter(appConfig)

	application := app.NewApp(*numberOfWorkers, *filePath, appConfig, dbConnector)
	results := application.Process()
	results.PrintReport()
	//application.PrintExecutionReport()
}
