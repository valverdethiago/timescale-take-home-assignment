package db

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

type DbConnector struct {
	db *sql.DB
}

func NewDbConnector(config config.AppConfig) DbConnector {
	db, err := sql.Open(config.DBDriver, config.DBSource)
	if err != nil {
		log.Fatal(err)
	}
	return DbConnector{db: db}
}

func (connector *DbConnector) ExecuteQuery(hostname string, interval domain.Interval) domain.QueryExecutionResult {
	stmt, err := connector.db.Prepare(DefaultQuery)
	if err != nil {
		log.Fatal(err)
	}
	qer := domain.NewQueryExecutionResult(hostname, interval)
	rows, err := stmt.Query(interval.StartTime, interval.EndTime, hostname)
	if err != nil {
		log.Fatal(err)
	}
	qer.End()
	defer rows.Close()
	defer stmt.Close()
	return qer
}

func (connector *DbConnector) CloseConnection() {
	connector.db.Close()
}
