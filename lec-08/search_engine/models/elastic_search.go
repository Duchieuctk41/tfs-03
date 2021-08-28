package models

import (
	"fmt"
	"log"
	"context"
	"github.com/olivere/elastic"
	"reflect"
	"time"
)

type Student struct {
	Name string `json:"name"`
}

const (
	indexName = "tfs03"
	indexMapping = `{
		"settings": {
			"number_of_shards":1,
			"number_of_replicas:0
		},
		"mappings": {
			"name":"text"
		}
	}`
)

var client *elastic.Client

func NewElasticSearchClient() *elastic.Client {
	var err error
	connected := false
	retries := 0

	for connected == false {
		client, err = elastic.NewClient(
			elastic.SetURL("http://localhost:9200"), elastic.SetSniff(false))
			if err != nil {
				if retries == 5 {
					log.Fatal(err)
				}
				fmt.Println("Elasticsearch isn't ready for connection", 5-retries, "less")
				retries++
				time.Sleep(3* time.Second)
			} else {
				connected = true
			}
	}

	esversion, err := client.ElasticsearchVersion("http://localhost:9200")
	if err != nil {
		panic(err)
	}
	fmt.Printf("Elasticsearch version %s\n", esversion)

	return client
}

func CheckExistsIndex() {
	exists := ExistsIndex(indexName)

	if !exists {
		CreateIndex(indexName)
	}

}

func ExistsIndex(i string) bool {
	exists, err := client.IndexExists(i).Do(context.TODO())
	if err != nil {
		log.Fatal(err)
	}

	return exists
}

func CreateIndex(i string) {
	createIndex, err := client.CreateIndex(i).
	Body(indexMapping).
	Do(context.Background())

	if err != nil {
		fmt.Println(err)
		return
	}

	if !createIndex.Acknowledged {
		log.Println("CreateIndex was not acknowledged. Check that timeout value is correct.")
	}
}

func SearchContent(input string)[]Student {
	students := []Student{}

	ctx := context.Background()

	q := elastic.NewMultiMatchQuery(input, "name").
	Type("most_fields").
	Fuzziness("2")
	result, err := client.Search().
	Index(indexName).
	Query(q).
	From(0).Size(20).
	Sort("_score", false).
	Do(ctx)
	if err != nil {
		log.Fatal(err)
	}

	var ttyp Student
	for _, student := range result.Each(reflect.TypeOf(ttyp)) {
		p := student.(Student)
		students = append(students, p)
	}
	return students
}
