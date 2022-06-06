package storage

import (
	"bytes"
	"encoding/gob"
	"errors"
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

// Encoder 接受任意类型的输入并序列化为 []byte 输出
func Encoder(data interface{}) ([]byte, error) {
	if data == nil {
		err := errors.New("data is nil")
		log.Println(err)
		return nil, err
	}
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return buffer.Bytes(), nil
}

// Decoder 接受 []byte 并反序列化到 v
func Decoder(data []byte, v interface{}) error {
	if data == nil {
		err := errors.New("data is nil")
		log.Println(err)
		return err
	}
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(v)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Set(key string, value interface{}) error {
	k := []byte(key)
	v, err := Encoder(value)
	if err != nil {
		log.Println(err)
		return err
	}
	err = DB.Put(k, v, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func Get(key string, value interface{}) error {
	k := []byte(key)
	buffer, err := DB.Get(k, nil)
	if err != nil {
		log.Println(err)
		return err
	}
	err = Decoder(buffer, value)
	return nil
}

func Delete(key string) error {
	k := []byte(key)
	return DB.Delete(k, nil)
}
