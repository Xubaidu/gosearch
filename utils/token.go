package utils

import (
	"github.com/yanyiwu/gojieba"
)

type Jieba = gojieba.Jieba

func Tokenizer(s string) []string {
	Jb := gojieba.NewJieba()
	defer Jb.Free()
	words := Jb.CutForSearch(s, true)
	return words
}
