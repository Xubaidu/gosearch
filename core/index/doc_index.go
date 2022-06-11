package index

import (
	"fmt"
	"go-search/core/storage"
	"strconv"
	"sync"
)

// ForwardDocDB maps docID to doc. It is a forward TokenIndex.
var ForwardDocDB *storage.LeveldbStorage

// RevDocDB maps doc to docID. It is a reverse TokenIndex.
var RevDocDB *storage.LeveldbStorage

// DocNum
// TODO: currently, we use DocNum to represent docID. But, is it a good idea?
var DocNum int64

func init() {
	ForwardDocDB = storage.ForwardDocDB
	RevDocDB = storage.RevDocDB
	DocNum = ForwardDocDB.Total()
}

func InsertForwardDocTokenIndex(doc string) error {
	docID := strconv.FormatInt(DocNum, 10)
	if has := ForwardDocDB.Has(docID); has {
		return fmt.Errorf("insert doc failed, doc already exists")
	}
	return ForwardDocDB.Set(docID, doc)
}

func InsertRevDocTokenIndex(doc string) error {
	if has := RevDocDB.Has(doc); has {
		return fmt.Errorf("insert doc failed, doc already exists")
	}
	defer func() {
		mut := sync.Mutex{}
		mut.Lock()
		defer mut.Unlock()
		DocNum++
	}()
	return RevDocDB.Set(doc, DocNum)
}
