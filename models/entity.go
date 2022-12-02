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

type ContentWords struct {
	Words []string `json:"words"`
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

	// filter key words
	resEntity.Keys = filterNameKey(splitWords(resEntity.Entity.Content))
	return resEntity, nil
}

func GetEntityWordsByNameAndIndex(name string, indexValue int) (resWords ContentWords, err error) {
	entityList, err := getFileJsonContent(name)
	if err != nil {
		fmt.Println("解析json 失败")
		return resWords, err
	}
	var lenValue = len(entityList)
	var entity = Entity{}
	fmt.Printf("index %d =====> len %d", indexValue, lenValue)
	if indexValue >= lenValue {
		entity = entityList[lenValue-1]
	} else {
		entity = entityList[indexValue]
	}
	resWords.Words = strings.FieldsFunc(entity.Content, splitString)
	return resWords, err
}

func splitString(c rune) bool {
	return c == '。' || c == '；' || c == '？' || c == '　'
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

// split words
func splitWords(content string) (words []string) {
	var seg jiebago.Segmenter
	seg.LoadDictionary("../data/dictionary.txt")

	// 拆分
	for word := range seg.Cut(content, true) {
		words = append(words, word)
	}
	return words
}

// filter key words
func filterNameKey(words []string) []string {
	filterKey := utils.Filter(words, func(key string) bool {
		// 过滤的字段
		hideKeys := getLimitKeyWords()
		return !strings.Contains(hideKeys, key) && key != "　" && len(key) > 1 && len(key) <= 8
	})
	return utils.RemoveDuplicateEle(filterKey)
}

func getLimitKeyWords() string {
	words, err := ioutil.ReadFile("./data/limit.txt")
	if err != nil {
		// 默认值
		return " .。，；“？：、！《》只为上下淫死鬼无之乎者也吗妈爸奶爷让退败狗犬狼哉兮矣而以杀傻笨血你我他她它的得"
	}
	return string(words)
}
