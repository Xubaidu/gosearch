package index

import (
	"errors"
	"go-search/core/model"
	"go-search/core/storage"
	"go-search/core/word"
	"log"
	"strconv"
)

type Index = model.Index

var ForwardDocDB *storage.LeveldbStorage
var RevIndexDB *storage.LeveldbStorage
var RevDocDB *storage.LeveldbStorage

func init() {
	ForwardDocDB = storage.ForwardDocDB
	RevIndexDB = storage.RevIndexDB
	RevDocDB = storage.RevDocDB
}

func InsertRevIndex(doc string) error {
	if err := InsertRevDoc(doc); err != nil {
		return err
	}
	docID := GetDocNum()
	if err := InsertForwardDoc(docID, doc); err != nil {
		return err
	}
	tokens := word.Tokenizer(doc)
	for _, token := range tokens {
		var index Index
		if has, _ := RevIndexDB.Has(token); has {
			err := RevIndexDB.Get(token, &index)
			if err != nil {
				log.Println(err)
				return err
			}
			index.DocList = append(index.DocList, docID)
			index.Count++
		} else {
			index.DocList = []int64{docID}
			index.Count = 1
		}
		err := RevIndexDB.Set(token, index)
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

func InsertRevDoc(doc string) error {
	if has, _ := RevDocDB.Has(doc); has {
		err := errors.New("insert doc failed, doc already exists")
		log.Println(err)
		return err
	}
	return RevDocDB.Set(doc, GetDocNum())
}

func InsertForwardDoc(docID int64, doc string) error {
	key := strconv.FormatInt(docID, 10)
	if has, _ := ForwardDocDB.Has(key); has {
		err := errors.New("insert doc failed, doc already exists")
		log.Println(err)
		return err
	}
	return ForwardDocDB.Set(key, doc)
}

func GetDocNum() int64 {
	return RevDocDB.Total()
}
