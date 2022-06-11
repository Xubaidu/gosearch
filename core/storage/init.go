package storage

var ForwardDocDB *LeveldbStorage
var RevDocDB *LeveldbStorage
var RevIndexDB *LeveldbStorage

func init() {
	ForwardDocDB = NewLevelDB("go-search/core/storage/forward_doc_db")
	RevDocDB = NewLevelDB("go-search/core/storage/rev_doc_db")
	RevIndexDB = NewLevelDB("go-search/core/storage/rev_index_db")
}
