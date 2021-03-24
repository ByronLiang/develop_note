package consumer

import (
	"context"
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"runtime"
	"sync"
)

type ExporterHandle func(product chan []string, ctx context.Context, wg *sync.WaitGroup)

func CsvExporter(fileName string, workerHandle ExporterHandle) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go signalWatch(cancel, os.Interrupt, os.Kill)
	producer := initProducer(f)
	wg := initWorker(producer, ctx, workerHandle)
	wg.Wait()
}

func initProducer(file *os.File) chan []string {
	producer := make(chan []string)
	go func(file *os.File, producer chan []string) {
		reader := csv.NewReader(file)
		defer func() {
			close(producer)
			file.Close()
			// producer finish
		}()
		for {
			content, err := reader.Read()
			if err == io.EOF {
				break
			} else if err != nil {
				fmt.Println(err)
				break
			}
			producer <- content
		}
	}(file, producer)

	return producer
}

func initWorker(product chan []string, ctx context.Context, workerHandle ExporterHandle) *sync.WaitGroup {
	wg := new(sync.WaitGroup)
	fmt.Println("cpu: ", runtime.NumCPU())
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		// worker process
		go workerHandle(product, ctx, wg)
	}
	return wg
}
