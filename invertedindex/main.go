package main

import (
	"fmt"
	"strings"
)

// 倒排索引结构体
type InvertedIndex map[string][]int

// 创建倒排索引
func createInvertedIndex(documents []string) InvertedIndex {
	invertedIndex := make(InvertedIndex)
	for id, doc := range documents {
		words := strings.Split(doc, " ")
		for _, word := range words {
			word = strings.ToLower(word)
			if ids, ok := invertedIndex[word]; ok {
				// 如果单词已经在倒排索引中，则追加文档ID
				ids = append(ids, id)
				invertedIndex[word] = ids
			} else {
				// 如果单词不在倒排索引中，则添加一个新的键值对
				invertedIndex[word] = []int{id}
			}
		}
	}
	return invertedIndex
}

// 查询单词出现在哪些文档中
func searchWord(word string, invertedIndex InvertedIndex) []int {
	word = strings.ToLower(word)
	if ids, ok := invertedIndex[word]; ok {
		return ids
	}
	return nil
}

func main() {
	documents := []string{
		"hello world",
		"goodbye world",
		"hello goodbye",
	}

	// 创建倒排索引
	invertedIndex := createInvertedIndex(documents)

	// 查询单词
	word := "hello"
	ids := searchWord(word, invertedIndex)
	fmt.Printf("单词'%s'出现在文档: %v\n", word, ids)
}
