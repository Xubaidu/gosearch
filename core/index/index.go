package index

import (
	"go-search/core/model"
	"go-search/core/storage"
	"go-search/utils"
	"sync"
)

type TokenIndex = model.TokenIndex

// RevTokenDB maps token to TokenIndex. It is a reverse TokenIndex
var RevTokenDB *storage.LeveldbStorage

var TokenNum int64

func init() {
	RevTokenDB = storage.RevTokenDB
	TokenNum = TotalToken()
}

func InsertRevTokenIndex(doc string, docID int64) error {
	tokens := utils.Tokenizer(doc)
	defer func() {
		mut := sync.Mutex{}
		mut.Lock()
		defer mut.Unlock()
		TokenNum += int64(len(tokens))
	}()
	for _, token := range tokens {
		var i TokenIndex
		if has := RevTokenDB.Has(token); has {
			err := RevTokenDB.Get(token, &i)
			if err != nil {
				return err
			}
			// TODO: ensure thread safety
			i.DocList = append(i.DocList, docID)
			i.Count++
		} else {
			i.DocList = []int64{docID}
			i.Count = 1
		}
		if err := RevTokenDB.Set(token, i); err != nil {
			return err
		}
	}
	return nil
}

func BuildRevTokenIndex(docs []string) error {
	for _, doc := range docs {
		if err := InsertRevDocTokenIndex(doc); err != nil {
			return err
		}
		if err := InsertForwardDocTokenIndex(doc); err != nil {
			return err
		}
		if err := InsertRevTokenIndex(doc, DocNum); err != nil {
			return err
		}
	}
	return nil
}

func TotalToken() int64 {
	RevTokenDB.Open()
	var count int64
	iter := RevTokenDB.DB.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		var i model.TokenIndex
		if err := utils.Decode(iter.Value(), &i); err == nil {
			count += int64(len(i.DocList))
		}
	}
	return count
}
