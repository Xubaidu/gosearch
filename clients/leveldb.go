package clients

import (
	"github.com/syndtr/goleveldb/leveldb"
	"log"
)

var DB *leveldb.DB

func InitLevelDB() error {
	DB, err := leveldb.OpenFile("clients/leveldb", nil)
	defer DB.Close()
	if err != nil {
		log.Printf("open db failed, err:%v\n", err)
		return err
	}
	return nil
}
