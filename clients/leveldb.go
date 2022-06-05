package clients

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var DB *leveldb.DB

func InitLevelDB() error {
	var err error
	DB, err = leveldb.OpenFile("leveldb", nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
