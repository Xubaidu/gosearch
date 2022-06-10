package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
	"go-search/core/model"
	"go-search/utils"
	"log"
)

type LeveldbStorage struct {
	db     *leveldb.DB
	path   string
	closed bool
}

func OpenDB(path string) (*leveldb.DB, error) {
	// FIXME: add bloom filter to reduce the time of disk IO
	db, err := leveldb.OpenFile(path, nil)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return db, nil
}

func (s *LeveldbStorage) Open() {
	// FIXME: no exception. is it ok?
	if !s.closed {
		log.Println("db is already opened")
		return
	}
	s.db, _ = OpenDB(s.path)
	s.closed = false
}

func NewLevelDB(path string) *LeveldbStorage {
	s := &LeveldbStorage{
		path:   path,
		closed: true,
	}
	return s
}

func (s *LeveldbStorage) Close() {
	if s.closed {
		log.Println("db is already closed")
		return
	}
	if err := s.db.Close(); err != nil {
		log.Println(err)
		return
	}
	s.closed = true
}

func (s *LeveldbStorage) Total() int64 {
	s.Open()
	var count int64
	iter := s.db.NewIterator(nil, nil)
	for iter.Next() {
		count++
	}
	iter.Release()
	return count
}

func (s *LeveldbStorage) TotalToken() int64 {
	s.Open()
	var count int64
	iter := s.db.NewIterator(nil, nil)
	for iter.Next() {
		var i model.Index
		if err := utils.Decode(iter.Value(), i); err == nil {
			count += int64(len(i.DocList))
		}
	}
	iter.Release()
	return count
}

func (s *LeveldbStorage) Set(key string, value interface{}) error {
	s.Open()
	v, err := utils.Encode(value)
	if err != nil {
		log.Println(err)
		return err
	}
	err = s.db.Put([]byte(key), v, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func (s *LeveldbStorage) Get(key string, value interface{}) error {
	s.Open()
	buffer, err := s.db.Get([]byte(key), nil)
	if err != nil {
		log.Println(err)
		return err
	}
	err = utils.Decode(buffer, value)
	return nil
}

func (s *LeveldbStorage) Delete(key string) error {
	s.Open()
	return s.db.Delete([]byte(key), nil)
}

func (s *LeveldbStorage) Has(key string) (bool, error) {
	s.Open()
	has, err := s.db.Has([]byte(key), nil)
	if err != nil {
		log.Println(err)
		// FIXME: error handling is not good. should we panic?
		return false, err
	}
	return has, nil
}
