package main

import (
	// Context object for Do() methods
	"context" // Log errors and quit
	// Get object methods and attributes
	"fmt"  // Format and print cluster data
	"time" // Set a timeout for the connection
	// Import the Olivere Golang driver for Elasticsearch

	"github.com/olivere/elastic"
)

var (
	client *elastic.Client
	ctx    context.Context
)

const (
	mapping = `{
		"settings":{
			"number_of_shards":1,
			"number_of_replicas":0
		},
		"mappings":{
			"properties":{
				"tags":{
					"type":"keyword"
				},
				"location":{
					"type":"geo_point"
				},
				"suggest_field":{
					"type":"completion"
				}
			}
		}
	}`
)

func main() {
	var err error
	// khai báo một số option của client
	options := []elastic.ClientOptionFunc{
		elastic.SetSniff(true),
		elastic.SetURL("http://localhost:9200"),         // nếu không có dòng này thì mặc định là 127.0.0.1:9200
		elastic.SetHealthcheckInterval(5 * time.Second), // ngưng kết nối sau 5 giây
	}

	// tạo client với các option trên
	client, err = elastic.NewClient(options...)
	if err != nil {
		panic(err)
	} else {
		fmt.Println("client:", client)
	}
	ctx = context.Background()

	DeleteAllIndexes()
	CreateIndex("some_index")
	IndexExists("some_index")
	IndexDocument("index2", "1", "film", mapping)
	GetDocument("index2", "1", "film", mapping)
	ListIndexNames()

	// Xoá một index
	DeleteIndex("index2")
	ListIndexNames()
}
