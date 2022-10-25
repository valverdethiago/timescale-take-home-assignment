package domain

type DbConnector interface {
	ExecuteQuery(hostname string, startTime string, endTime string) error
}
