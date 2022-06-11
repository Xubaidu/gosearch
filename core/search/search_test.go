package search

import (
	"fmt"
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
				doc: "中国最好的大学",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := Search(tt.args.doc)
			fmt.Println(ans)
		})
	}
}
