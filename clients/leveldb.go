package clients

import (
	"bytes"
	"encoding/gob"
	"github.com/syndtr/goleveldb/leveldb"
	"io/ioutil"
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

//序列化到文件
func Encoder(data interface{}) []byte {
	if data == nil {
		return nil
	}
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	return buffer.Bytes()
}

//解码序列化文件
func Decoder(data []byte, v interface{}) {
	if data == nil {
		return
	}
	buffer := bytes.NewBuffer(data)
	decoder := gob.NewDecoder(buffer)
	err := decoder.Decode(v)
	if err != nil {
		panic(err)
	}
}

func Set(key []byte, value interface{}) {
	var v = Encoder(value)
	err := DB.Put(key, v, nil)
	if err != nil {
		panic(err)
	}
}

func Get(key []byte) interface{} {
	buffer, err := DB.Get(key, nil)
	if err != nil {
		return err
	}
	var v interface{}
	Decoder(buffer, v)
	return v
}

func Delete(key []byte) error {
	return DB.Delete(key, nil)
}

//写入二进制数据到磁盘
func storeToGob(data interface{}, filename string) {
	buffer := new(bytes.Buffer)
	encoder := gob.NewEncoder(buffer)
	err := encoder.Encode(data)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, buffer.Bytes(), 0600)
	if err != nil {
		panic(err)
	}
}

//磁盘加载二进制数据
func load(data interface{}, filename string) {
	raw, err := ioutil.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	buffer := bytes.NewBuffer(raw)
	dec := gob.NewDecoder(buffer)
	err = dec.Decode(data)
	if err != nil {
		panic(err)
	}
}
