package search

import (
	"fmt"
	"go-search/core/index"
	"go-search/core/storage"
	"testing"
)

func TestSearch(t *testing.T) {
	type args struct {
		doc string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "test1",
			args: args{
				doc: "我来到清华大学",
			},
		},
	}
	storage.InitDB()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			docs := []string{
				"我来到清华大学",
				"你来到清华大学",
			}
			err := index.BuildRevIndex(docs)
			if err != nil {
				t.Error(err)
			}
			ans := Search(tt.args.doc)
			if err != nil {
				t.Error(err)
			}
			fmt.Println(ans)
		})
	}
}
