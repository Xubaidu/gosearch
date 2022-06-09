package index

import (
	"fmt"
	"testing"
)

func TestCalc(t *testing.T) {
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
	InitDB()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			docs := []string{
				"我来到清华大学",
				"你来到清华大学",
			}
			err := BuildRevIndex(docs)
			if err != nil {
				t.Error(err)
			}
			indexes, err := Calc(tt.args.doc)
			if err != nil {
				t.Error(err)
			}
			for _, index := range indexes {
				fmt.Println(*index)
			}
		})
	}
}
