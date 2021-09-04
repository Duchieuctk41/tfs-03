package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"io"
	"log"
	"os"
	"sync"

	"github.com/elastic/go-elasticsearch/v6"
	"github.com/elastic/go-elasticsearch/v6/esapi"
)

type Review struct {
	Vote    string
	Title   string
	Content string
}

func main() {
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error creating the client: %s", err)
	}
	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()
	ch := make(chan Review)
	var wg sync.WaitGroup

	wg.Add(2)
	go GetReview(&wg, ch)
	go SetReview(es, &wg, ch)
	wg.Wait()

	log.Println("success")
}

func GetReview(wg *sync.WaitGroup, ch chan Review) {
	defer wg.Done()
	defer close(ch)
	file, err := os.Open("csv/train.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	r := csv.NewReader(file)
	for i := 0; i < 3; i++ {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		item := Review{
			Vote:    record[0],
			Title:   record[1],
			Content: record[2],
		}
		ch <- item
	}
}

func SetReview(es *elasticsearch.Client, wg *sync.WaitGroup, ch chan Review) {
	defer wg.Done()
	for {
		item, ok := <-ch
		if !ok {
			return
		}
		line, _ := json.Marshal(item)
		req := esapi.IndexRequest{
			Index:      "reviews",
			DocumentID: "1",
			Body:       bytes.NewReader(line),
		}
		req.Do(context.Background(), es)
	}
}
