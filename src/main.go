package main

import (
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	"github.com/valverdethiago/timescale-take-home-assignment/adapter"
	"github.com/valverdethiago/timescale-take-home-assignment/app"
	config "github.com/valverdethiago/timescale-take-home-assignment/config"
)

var (
	invalidLineError = fmt.Errorf("invalid line format")
)

func main() {

	numberOfWorkers := flag.Int("workers", 2, "Number of workers")
	filePath := flag.String("file", "../query_params.csv", "File Path")

	flag.Parse()

	fmt.Printf("Number of workers: %d\n", *numberOfWorkers)
	fmt.Printf("File: %s\n", *filePath)
	fmt.Println("tail:", flag.Args())

	config := config.LoadEnvConfig()
	dbConnector := adapter.NewTimescaleAdapter(config)

	application := app.NewApp(*numberOfWorkers, *filePath, config, dbConnector)
	application.Process()
}
