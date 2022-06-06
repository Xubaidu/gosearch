package index

import (
	"fmt"
	"go-search/storage"
	"log"
	"testing"
)

func TestCalc(t *testing.T) {
	type args struct {
		tokens []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				tokens: []string{"我", "来", "到", "清华", "大学"},
			},
		},
	}
	err := storage.InitLevelDB()
	if err != nil {
		log.Println(err)
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			docs := []string{
				"我来到清华大学",
				"你来到清华大学",
			}
			for _, doc := range docs {
				err := InsertRevIndex(doc)
				if err != nil {
					t.Error(err)
				}
			}
			indexes, err := Calc(tt.args.tokens)
			if err != nil {
				t.Error(err)
			}
			for _, index := range indexes {
				fmt.Println(*index)
			}
		})
	}
}
