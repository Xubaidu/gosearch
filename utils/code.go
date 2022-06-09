package utils

import (
	"bytes"
	"encoding/gob"
	"errors"
	"log"
)

// Encode 接受任意类型的输入并序列化为 []byte 输出
func Encode(data interface{}) ([]byte, error) {
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

// Decode 接受 []byte 并反序列化到 v
func Decode(data []byte, v interface{}) error {
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
