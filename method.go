package main

import (
	"fmt"
)

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
		fmt.Println(err)
	} else {
		fmt.Println(exist)
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
