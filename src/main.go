package main

import (
	"bufio"
	"flag"
	"fmt"
	_ "github.com/lib/pq"
	config "github.com/valverdethiago/timescale-take-home-assignment/config"
	"github.com/valverdethiago/timescale-take-home-assignment/db"
	"github.com/valverdethiago/timescale-take-home-assignment/domain"
	"log"
	"os"
	"strings"
)

var (
	invalidLineError = fmt.Errorf("invalid line format")
)

func main() {

	numberOfWorkers := flag.Int("workers", 10, "Number of workers")
	filePath := flag.String("file", "../query_params.csv", "File Path")

	flag.Parse()

	fmt.Printf("Number of workers: %d\n", *numberOfWorkers)
	fmt.Printf("File: %s\n", *filePath)
	fmt.Println("tail:", flag.Args())

	config := config.LoadEnvConfig()
	dbConnector := db.NewDbConnector(config)

	hostIntervalMap, err := readFile(*filePath)
	if err != nil {
		log.Fatal(err)
	}
	for hostname, elements := range *hostIntervalMap {
		fmt.Println("hostname:", hostname)
		for _, interval := range elements {
			qer := dbConnector.ExecuteQuery(hostname, interval)
			fmt.Println("\tInterval: ", interval, "| QueryExecutionResult: ", qer)
		}
	}
	defer dbConnector.CloseConnection()

}

func readFile(filePath string) (*map[string][]domain.Interval, error) {

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	result := make(map[string][]domain.Interval)

	scanner := bufio.NewScanner(file)
	scanner.Scan() //skiping the first line
	for scanner.Scan() {
		line := scanner.Text()
		fileEntry, err := parseLine(line)
		if err != nil {
			return nil, err
		}
		if _, exists := result[fileEntry.Hostname]; exists {
			result[fileEntry.Hostname] = append(result[fileEntry.Hostname], fileEntry.Interval)
		} else {
			result[fileEntry.Hostname] = []domain.Interval{fileEntry.Interval}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return &result, nil
}

func parseLine(line string) (*domain.FileEntry, error) {
	array := strings.Split(line, ",")
	if len(array) < 3 {
		return nil, invalidLineError
	}
	return &domain.FileEntry{
		Hostname: array[0],
		Interval: domain.Interval{
			StartTime: array[1],
			EndTime:   array[2],
		},
	}, nil
}
