package word

import (
	"github.com/yanyiwu/gojieba"
)

type Jieba = gojieba.Jieba

func tokenizer(s string) []string {
	Jb := gojieba.NewJieba()
	defer Jb.Free()
	words := Jb.CutForSearch(s, true)
	return words
}
