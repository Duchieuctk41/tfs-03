package main

import (
	"bytes"
	"context"
	"encoding/csv"
	"encoding/json"
	"fmt"
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
	// connect db
	es, err := elasticsearch.NewDefaultClient()
	if err != nil {
		log.Fatalf("Error when create es client: %s", err)
	}

	res, err := es.Info()
	if err != nil {
		log.Fatalf("Error when get response %s", err)
	}

	defer res.Body.Close()

	// scan csv and push to db
	wg := sync.WaitGroup{}
	ch := make(chan Review)

	wg.Add(2)

	go ScanCSVFile(&wg, ch)
	go PushCSVToDatabase(es, &wg, ch)

	wg.Wait()
	fmt.Println("done")
}

func ScanCSVFile(wg *sync.WaitGroup, ch chan<- Review) {
	defer wg.Done()
	defer close(ch)

	file, err := os.Open("csv/train.csv")
	if err != nil {
		log.Fatal(err)
	}

	r := csv.NewReader(file)

	for {
		record, err := r.Read()
		if err == io.EOF {
			log.Fatal()
		}
		if err != nil {
			log.Fatal()
		}

		item := Review{
			Vote:    record[0],
			Title:   record[1],
			Content: record[2],
		}
		ch <- item
	}
}

func PushCSVToDatabase(es *elasticsearch.Client, wg *sync.WaitGroup, ch <-chan Review) {
	defer wg.Done()
	for {
		item, ok := <-ch
		if !ok {
			return
		}

		line, _ := json.Marshal(item) // parsing struct to JSON

		// create db
		req := esapi.IndexRequest{
			Index: "test",
			Body:  bytes.NewReader(line),
		}

		req.Do(context.Background(), es)
	}
}

func Search(es *elasticsearch.Client) {
	str := "i found something"

	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"Content": str,
			},
		},
	}

	fmt.Printf("%T", query)

	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	res, err := es.Search(
		es.Search.WithContext(context.Background()),
		es.Search.WithIndex("reviews"),
		es.Search.WithBody(&buf),
		es.Search.WithTrackTotalHits(true),
		es.Search.WithPretty(),
	)
	if err != nil {
		log.Fatalf("Error getting response: %s", err)
	}
	defer res.Body.Close()

}
