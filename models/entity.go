package models

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"name_shop_api/utils"
	"strings"

	"github.com/wangbin/jiebago"
)

type Entity struct {
	Content string `json:"content"`
	Title   string `json:"title"`
	Author  string `json:"author"`
	Book    string `json:"book"`
	Dynasty string `json:"dynasty"`
}

type ResEntity struct {
	Entity Entity
	Keys   []string
}

func GetEntityByName(name string) ([]ResEntity, error) {
	entityList, err := getFileJsonContent(name)
	if err != nil {
		fmt.Println("解析json 失败")
		return nil, err
	}

	resultList := []ResEntity{}
	for _, entity := range entityList {
		resEntity := ResEntity{
			Entity: entity,
			Keys:   nil,
		}
		resultList = append(resultList, resEntity)
	}

	return resultList, nil
}

func GetEntityByNameAndIndex(name string, indexValue int) (resEntity ResEntity, err error) {
	entityList, err := getFileJsonContent(name)
	if err != nil {
		fmt.Println("解析json 失败")
		return resEntity, err
	}
	var lenValue = len(entityList)
	fmt.Printf("index %d =====> len %d", indexValue, lenValue)
	if indexValue >= lenValue {
		resEntity.Entity = entityList[lenValue-1]
	} else {
		resEntity.Entity = entityList[indexValue]
	}

	// resEntity.Keys = getContentKeys(resEntity.Entity.Content)

	// jieba 分词工具
	// jieba := gojieba.NewJieba()
	// defer jieba.Free()

	// seg 分词工具
	var seg jiebago.Segmenter
	seg.LoadDictionary("../data/dictionary.txt")

	var words []string

	if len(resEntity.Entity.Content) > 2000 {
		// 分段拆分
		contentKeys := strings.FieldsFunc(resEntity.Entity.Content, splitString)
		for _, wordStr := range contentKeys {
			for word := range seg.Cut(wordStr, true) {
				words = append(words, word)
			}
		}
	} else {
		for word := range seg.Cut(resEntity.Entity.Content, true) {
			words = append(words, word)
		}
	}

	// filter key words
	filterKey := utils.Filter(words, func(key string) bool {
		return key != "　" && key != "." && key != "，" && key != "。" && key != "；" && key != "“" && key != "？" && key != "：" && key != "、" && key != "！" && len(key) > 1 && len(key) <= 8
	})

	resEntity.Keys = filterKey
	return resEntity, nil
}

func getFileJsonContent(name string) ([]Entity, error) {
	bytes, err := ioutil.ReadFile("./data/" + name + ".json")
	if err != nil {
		fmt.Println("打开文件失败")
		return nil, err
	}
	entityList := []Entity{}
	err = json.Unmarshal(bytes, &entityList)
	return entityList, err
}

// func getContentKeys(content string) []string {
// 	var words []string

// 	jieba := gojieba.NewJieba()
// 	defer jieba.Free()

// 	if len(content) > 2000 {
// 		// 分段拆分
// 		contentKeys := strings.FieldsFunc(content, splitString)
// 		for _, wordStr := range contentKeys {
// 			words = append(words, jieba.Cut(wordStr, true)...)
// 		}
// 	} else {
// 		words = jieba.Cut(content, true)
// 	}

// 	// filter key words
// 	filterKey := utils.Filter(words, func(key string) bool {
// 		return key != "　" && key != "." && key != "，" && key != "。" && key != "；" && key != "“" && key != "？" && key != "：" && key != "、" && key != "！" && len(key) > 1 && len(key) <= 8
// 	})
// 	return filterKey
// }

func splitString(r rune) bool {
	return r == '。' || r == '？' || r == '；' || r == '！'
}
