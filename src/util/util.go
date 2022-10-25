package util

import "github.com/valverdethiago/timescale-take-home-assignment/domain"

func GetKeysFromMap(source map[string][]domain.Interval) []string {
	var result []string
	for key := range source {
		result = append(result, key)
	}
	return result
}
