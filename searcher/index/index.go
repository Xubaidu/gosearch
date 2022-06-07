package index

import (
	"errors"
	"go-search/searcher/word"
	"go-search/storage"
	"log"
)

type Index struct {
	DocList []int64
	Count   int64
	Token   string
}

func InsertRevIndex(doc string) error {
	// FIXME: add storage.GetDoc(doc) function
	if _, err := storage.GetDoc(doc); err == nil {
		err = errors.New("doc already exists")
		log.Println(err)
		return err
	}
	// FIXME: add GetDocNum() function
	docID := storage.GetDocNum(doc) + 1
	tokens := word.Tokenizer(doc)
	for _, token := range tokens {
		var index Index
		if err := storage.Get(token, &index); err != nil {
			index.DocList = []int64{docID}
			index.Count = 1
		} else {
			index.DocList = append(index.DocList, docID)
			index.Count++
		}
		err := storage.Set(token, &Index{
			DocList: append(index.DocList, docID),
			Count:   index.Count + 1,
			Token:   token,
		})
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func BuildRevIndex(docs []string) error {
	for _, doc := range docs {
		err := InsertRevIndex(doc)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

func Calc(tokens []string) ([]*Index, error) {
	var indexes []*Index
	for _, token := range tokens {
		var index Index
		err := storage.Get(token, &index)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		indexes = append(indexes, &index)
	}
	return indexes, nil
}
