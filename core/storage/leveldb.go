package storage

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
	"go-search/utils"
	"log"
)

type LeveldbStorage struct {
	DB     *leveldb.DB
	path   string
	closed bool
}

func NewLeveldbStorage(path string) *LeveldbStorage {
	s := &LeveldbStorage{
		path:   path,
		closed: true,
	}
	return s
}

func OpenDB(path string) (*leveldb.DB, error) {
	// 使用布隆过滤器，查询的时候可以先判断 key 是否存在，从而减少磁盘 IO 的次数
	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}
	db, err := leveldb.OpenFile(path, o)
	return db, err
}

func (s *LeveldbStorage) Open() {
	if !s.closed {
		return
	}
	s.DB, _ = OpenDB(s.path)
	s.closed = false
	return
}

func (s *LeveldbStorage) Close() {
	if s.closed {
		return
	}
	_ = s.DB.Close()
	s.closed = true
}

func (s *LeveldbStorage) Total() int64 {
	s.Open()
	var count int64
	iter := s.DB.NewIterator(nil, nil)
	defer iter.Release()
	for iter.Next() {
		count++
	}
	return count
}

func (s *LeveldbStorage) Set(key string, value interface{}) error {
	s.Open()
	v, err := utils.Encode(value)
	if err != nil {
		log.Println(err)
		return err
	}
	err = s.DB.Put([]byte(key), v, nil)
	if err != nil {
		return err
	}
	return nil
}

func (s *LeveldbStorage) Get(key string, value interface{}) error {
	s.Open()
	buffer, err := s.DB.Get([]byte(key), nil)
	if err != nil {
		return err
	}
	err = utils.Decode(buffer, value)
	return err
}

func (s *LeveldbStorage) Delete(key string) error {
	s.Open()
	return s.DB.Delete([]byte(key), nil)
}

func (s *LeveldbStorage) Has(key string) bool {
	s.Open()
	has, _ := s.DB.Has([]byte(key), nil)
	return has
}
