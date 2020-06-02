package main

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/olivere/elastic"
)

const (
	indexName = "students"
	_type     = "_doc"
)

var (
	client *elastic.Client
	ctx    context.Context
	err    error
)

type Student struct {
	Id           string  `json:"id"`
	Name         string  `json:"name"`
	Age          int64   `json:"age"`
	AverageScore float64 `json:"average_score"`
}

func main() {
	client, err = NewClient()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("client info:", client)

	ctx = context.Background()

	for i := 1; i < 15; i++ {
		newStudent := Student{
			Id:           fmt.Sprint(i),
			Name:         "User " + fmt.Sprint(i),
			Age:          (int64)(40 - 10%i),
			AverageScore: (float64)(100 - (40 - 10%i)),
		}

		CreateIndex(newStudent, indexName, fmt.Sprint(i))
	}

	IndexExists(indexName)
	GetDocument(indexName, "1")
	ListIndexNames()
	DeleteIndex(indexName)
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
func CreateIndex(newStudent Student, index, id string) {
	dataJSON, err := json.Marshal(newStudent)
	js := string(dataJSON)
	_, err = client.Index().
		Index(indexName).
		BodyJson(js).
		Type(_type).
		Id(id).
		Do(ctx)

	if err != nil {
		panic(err)
	}

	fmt.Println("inserted:", newStudent.Name)
}

// Kiểm tra index tồn tại chưa
func IndexExists(index string) {
	isExist, err := client.IndexExists(index).Do(ctx)
	if err != nil {
		fmt.Println("exist (error):", err)
	} else {
		fmt.Println("exist:", isExist)
	}
}

// Lấy document từ index
func GetDocument(index, id string) {
	res, err := client.Get().
		Index(index).
		Id(id).
		Type(_type).
		Do(ctx)

	if err != nil {
		fmt.Println(err)
	} else {
		if res.Found {
			fmt.Println("document info: ", res.Id, res.Version, res.Index, res.Type)
		} else {
			fmt.Println("not found document")
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
