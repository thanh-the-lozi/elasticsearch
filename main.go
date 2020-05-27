package main

import (
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
	err    error
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
	client, err = NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("client info:", client)

	ctx = context.Background()

	DeleteAllIndexes()
	// CreateIndex("some_index")
	// IndexExists("s")
	// IndexDocument("index2", "1", "film", mapping)
	// GetDocument("students", "1", "", mapping)
	// ListIndexNames()

	// // Xoá một index
	// DeleteIndex("index2")
	// ListIndexNames()
}

//=============================//

// Tạo client
func NewClient() (*elastic.Client, error) {
	// khai báo một số option của client
	options := []elastic.ClientOptionFunc{
		elastic.SetSniff(true),
		elastic.SetURL("http://localhost:9200"),         // nếu không có dòng này thì mặc định là 127.0.0.1:9200
		elastic.SetHealthcheckInterval(5 * time.Second), // ngưng kết nối sau 5 giây
	}

	// tạo client với các option trên
	return elastic.NewClient(options...)
}

// Tạo index
func CreateIndex(index string) {
	createdIndex, err := client.CreateIndex(index).BodyString(mapping).Do(ctx)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("create: ", createdIndex)
}

// Kiểm tra index tồn tại chưa
func IndexExists(index string) {
	exist, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		fmt.Println("exist:", err)
	} else {
		fmt.Println("exist:", exist)
	}
}

// Đánh index cho document
func IndexDocument(index, id, _type string, document interface{}) {
	put1, _ := client.Index().
		Index(index).
		BodyJson(document).
		Type(_type).
		Id(id).
		Do(ctx)

	fmt.Println("index: ", put1.Id, put1.Version, put1.Index, put1.Type)
}

// Lấy document từ index
func GetDocument(index, id, _type string, document interface{}) {
	res, err := client.Get().Index(index).Id(id).Type(_type).Do(ctx)
	if err != nil {
		fmt.Println("get: ", err)
	} else {
		if res.Found {
			fmt.Println("get res: ", res.Id, res.Version, res.Index, res.Type)
		} else {
			fmt.Println("not found")
		}
	}
}

// Xóa index
func DeleteIndex(index string) {
	_, err := client.DeleteIndex(index).Do(ctx)

	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println("deleted index: ", index)
}

// Lấy danh sách tên index
func ListIndexNames() {
	names, _ := client.IndexNames()
	for _, name := range names {
		fmt.Println(name)
	}
}

// Xóa tất cả index
func DeleteAllIndexes() {
	names, _ := client.IndexNames()
	for _, name := range names {
		DeleteIndex(name)
	}
}
