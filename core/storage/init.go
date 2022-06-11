package storage

var ForwardDocDB *LeveldbStorage
var RevDocDB *LeveldbStorage
var RevTokenDB *LeveldbStorage

func init() {
	ForwardDocDB = NewLeveldbStorage("forward_doc_db")
	RevDocDB = NewLeveldbStorage("rev_doc_db")
	RevTokenDB = NewLeveldbStorage("rev_token_db")
}
