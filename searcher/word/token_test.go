package word

import (
	"fmt"
	"testing"
)

func Test_tokenizer(t *testing.T) {
	tests := []struct {
		name string
		text string
	}{
		{
			name: "test 1",
			text: "我来到清华大学",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ans := tokenizer(tt.text)
			fmt.Println(ans)
		})
	}
}
