package main

import (
	"fmt"
	"github.com/blevesearch/bleve/v2"
)

func main() {
	// 创建一个索引
	index, err := bleve.New("build", bleve.NewIndexMapping())
	if err != nil {
		panic(err)
	}
	defer index.Close()

	// 定义文档结构
	type Document struct {
		Title string `json:"title"`
		Body  string `json:"body"`
	}

	// 添加文档
	documents := []Document{
		{Title: "Go 语言教程", Body: "Go 语言是一种开源的编程语言"},
		{Title: "Python 语言教程", Body: "Python 语言是一种解释型的编程语言"},
		{Title: "xxxxx", Body: "嘻嘻嘻嘻嘻嘻"},
	}

	for _, doc := range documents {
		id := fmt.Sprintf("%d", hash(doc.Title)) // 使用标题生成一个简单的ID
		err = index.Index(id, doc)
		if err != nil {
			panic(err)
		}
	}

	// 执行搜索
	searchRequest := bleve.NewSearchRequest(bleve.NewMatchQuery("编程"))
	// 显式要求返回字段信息
	searchRequest.Fields = []string{"*"}
	searchResult, err := index.Search(searchRequest)
	if err != nil {
		panic(err)
	}

	// 打印搜索结果
	for _, hit := range searchResult.Hits {
		fmt.Printf("Found: %+v\n", hit.Fields)
	}
}

// 简单的哈希函数用于生成文档ID
func hash(s string) int {
	h := 0
	for _, r := range s {
		h = 31*h + int(r)
	}
	return h
}
