package consumer

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"sync"
	"testing"
)

func TestCsvWorker(t *testing.T) {
	CsvWorker("test.csv")
}

func TestCsvExporter(t *testing.T) {
	CsvExporter("test.csv", ExporterWorkerHandle)
}

func ExporterWorkerHandle(product chan []string, ctx context.Context, wg *sync.WaitGroup)  {
	defer wg.Done()
	for {
		select {
		case <-ctx.Done():
			return
		case data, ok := <-product:
			if !ok {
				fmt.Println("end")
				return
			}
			fmt.Println("consumer: ", data)
		}
	}
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
