package adapter

import (
	"database/sql"
	"github.com/valverdethiago/timescale-take-home-assignment/config"
	"github.com/valverdethiago/timescale-take-home-assignment/domain"
	"log"
)

const (
	DefaultQuery string = "   SELECT TIME_BUCKET('1 minute', ts) AS EVERY_MIN, AVG(usage) " +
		"                       FROM CPU_USAGE " +
		"                      WHERE ts BETWEEN $1 AND $2 " +
		"                        AND host = $3 " +
		"                   GROUP BY every_min"
)

type TimescaleAdapter struct {
	db *sql.DB
}

func NewTimescaleAdapter(config config.AppConfig) domain.DbConnector {
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	return &TimescaleAdapter{db: db}
}

func (adapter *TimescaleAdapter) ExecuteQuery(hostname string, startTime string, endTime string) error {
	stmt, err := adapter.db.Prepare(DefaultQuery)
	if err != nil {
		return err
	}
	rows, err := stmt.Query(startTime, endTime, hostname)
	if err != nil {
		return err
	}
	defer rows.Close()
	defer stmt.Close()
	return nil
}

func (adapter *TimescaleAdapter) CloseConnection() {
	adapter.db.Close()
}
