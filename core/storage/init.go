package storage

var ForwardDocDB *LeveldbStorage
var RevDocDB *LeveldbStorage
var RevTokenDB *LeveldbStorage

func init() {
	// TODO: use environment variable to replace the following hard-coded path
	ForwardDocDB = NewLeveldbStorage("../storage/forward_doc_db")
	RevDocDB = NewLeveldbStorage("../storage/rev_doc_db")
	RevTokenDB = NewLeveldbStorage("../storage/rev_token_db")
}
