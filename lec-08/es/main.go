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
	"time"

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

	time_start := time.Now()
	wg.Add(2)
	go GetReview(&wg, ch)
	go SetReview(es, &wg, ch)
	wg.Wait()
	time_end := time.Now()
	log.Println(time_end.Sub(time_start))
	// Search(es)
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
	for {
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
			Index: "reviews",
			Body:  bytes.NewReader(line),
		}
		req.Do(context.Background(), es)
	}
}

func Search(es *elasticsearch.Client) {
	time_start := time.Now()
	str := "It does sort of sway a little bit when you push the clothes to make more room, but because the pieces are screwed together I think it's fine and doesn't seem like it will topple over. For my purposes, it works just fine and fits nicely beside the washer without taking up too much room. Definitely has saved me from doing the trip up the stairs with an armful of wet"
	var buf bytes.Buffer
	query := map[string]interface{}{
		"query": map[string]interface{}{
			"match": map[string]interface{}{
				"Content": str,
			},
		},
	}
	if err := json.NewEncoder(&buf).Encode(query); err != nil {
		log.Fatalf("Error encoding query: %s", err)
	}

	// Perform the search request.
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

	time_end := time.Now()
	log.Println(time_end.Sub(time_start))
	log.Println(res)
}
