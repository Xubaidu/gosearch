package storage

var ForwardDocDB *LeveldbStorage
var RevDocDB *LeveldbStorage
var RevIndexDB *LeveldbStorage

func InitDB() {
	ForwardDocDB = NewLevelDB("forward_doc_db")
	RevDocDB = NewLevelDB("rev_doc_db")
	RevIndexDB = NewLevelDB("rev_index_db")
}
