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

func CsvWorker(fileName string) {
	f, err := os.Open(fileName)
	if err != nil {
		fmt.Println(err)
		return
	}
	r := csv.NewReader(f)
	product := make(chan []string)
	res := make(chan interface{})
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	initProducerHandle(r, product)
	wg := initWorkerHandle(product, res, ctx)
	go func() {
		wg.Wait()
		close(res)
	}()
	consumeHandle(res)
}

func initProducerHandle(reader *csv.Reader, producer chan []string)  {
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
	close(producer)
}

func initWorkerHandle(product chan []string, res chan interface{}, ctx context.Context) *sync.WaitGroup {
	wg := new(sync.WaitGroup)
	for i := 0; i < runtime.NumCPU(); i++ {
		wg.Add(1)
		go workerHandle(product, res, ctx)
	}
	return wg
}

func workerHandle(product chan []string, res chan interface{}, ctx context.Context) {
	for {
		select {
		case data, ok := <-product:
			if !ok {
				return
			}
			fmt.Println(data)
			res <- data
		case <-ctx.Done():
			return
		}
	}
}

func consumeHandle(res chan interface{})  {
	for data := range res {
		fmt.Println(data)
	}
}
