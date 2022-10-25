package util

import (
	"bufio"
	"fmt"
	"github.com/valverdethiago/timescale-take-home-assignment/domain"
	"os"
	"strings"
)

var (
	invalidLineError = fmt.Errorf("invalid line format")
)

type FileReader struct {
	FilePath string
	Result   *map[string][]domain.Interval
}

func (fileReader *FileReader) Read() error {

	file, err := os.Open(fileReader.FilePath)
	if err != nil {
		return err
	}
	defer file.Close()
	result := make(map[string][]domain.Interval)

	scanner := bufio.NewScanner(file)
	scanner.Scan() //skipping the first line
	for scanner.Scan() {
		line := scanner.Text()
		fileEntry, err := parseLine(line)
		if err != nil {
			return err
		}
		if _, exists := result[fileEntry.Hostname]; exists {
			result[fileEntry.Hostname] = append(result[fileEntry.Hostname], fileEntry.Interval)
		} else {
			result[fileEntry.Hostname] = []domain.Interval{fileEntry.Interval}
		}
	}

	if err := scanner.Err(); err != nil {
		return err
	}
	fileReader.Result = &result
	return nil
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
