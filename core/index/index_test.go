package index

import (
	"fmt"
	"log"
	"testing"
)

func TestBuildRevTokenIndex(t *testing.T) {
	tests := []struct {
		name string
		docs []string
	}{
		{
			name: "test 1",
			docs: []string{
				"我来到清华大学",
				"你来到清华大学",
				"华东师范大学建校七十周年",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := BuildRevTokenIndex(tt.docs); err != nil {
				log.Println(err)
			}
			fmt.Printf("TokenNum = %d, DocNum = %d\n", TokenNum, DocNum)
		})
	}
}
