package index

import (
	"errors"
	"go-search/searcher/word"
	"go-search/storage"
	"log"
)

var DocDB *storage.LeveldbStorage
var RevIndexDB *storage.LeveldbStorage

type Index struct {
	DocList []int64
	Count   int64
	Token   string
}

func InitDB() {
	DocDB = storage.NewLevelDB("doc_db")
	RevIndexDB = storage.NewLevelDB("rev_index_db")
}

func InsertRevIndex(doc string) error {
	if err := InsertDoc(doc); err != nil {
		return err
	}
	docID := GetDocNum()
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
			index.Token = token
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

func Calc(doc string) ([]*Index, error) {
	var indexes []*Index
	tokens := word.Tokenizer(doc)
	for _, token := range tokens {
		var index Index
		if err := RevIndexDB.Get(token, &index); err == nil {
			indexes = append(indexes, &index)
		}
	}
	return indexes, nil
}

func InsertDoc(doc string) error {
	if has, _ := DocDB.Has(doc); has {
		err := errors.New("insert doc failed, doc already exists")
		log.Println(err)
		return err
	}
	return DocDB.Set(doc, GetDocNum())
}

func GetDocNum() int64 {
	return DocDB.Total()
}
