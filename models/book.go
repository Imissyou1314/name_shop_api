package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Book struct {
	Title    string `json:"title"`
	FileName string `json:"fileName"`
	Tag      string `json:"tag"`
	Desc     string `json:"desc"`
	Total    int    `json:"total"`
}

func GetAllBooks() ([]Book, error) {
	return getFileBookJsonContent()
}

func getFileBookJsonContent() ([]Book, error) {
	bytes, err := ioutil.ReadFile("./data/book.json")
	if err != nil {
		fmt.Println("打开文件失败")
		return nil, err
	}
	bookList := []Book{}
	err = json.Unmarshal(bytes, &bookList)
	return bookList, err
}
