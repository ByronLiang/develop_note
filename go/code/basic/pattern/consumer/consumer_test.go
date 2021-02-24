package consumer

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"testing"
)

func TestCondConsumer(t *testing.T) {
	CondConsumer()
}

func TestCsvWorker(t *testing.T) {
	CsvWorker("test.csv")
}

func TestReadCsv(t *testing.T) {
	f, err := os.Open("test.csv")
	if err != nil {
		fmt.Println(err)
		return
	}
	r := csv.NewReader(f)
	defer f.Close()
	for {
		content, err := r.Read()
		if err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
			break
		}
		id, name, age := content[0], content[1], content[2]
		fmt.Printf("id: %v; name: %v; age: %v \n", id, name, age)
	}
}
