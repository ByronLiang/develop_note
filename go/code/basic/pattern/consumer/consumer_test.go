package consumer

import "testing"

func TestCondConsumer(t *testing.T) {
	CondConsumer()
}

func TestCsvWorker(t *testing.T) {
	CsvWorker("./sample-csv.csv")
}
