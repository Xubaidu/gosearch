package storage

import (
	"fmt"
	"testing"
)

func TestInitLevelDB(t *testing.T) {
	tests := []struct {
		name string
	}{
		{
			name: "test 1",
		},
	}
	s := NewLevelDB("test")
	s.Open()
	defer s.Close()
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = s.Set("key", "value")
			var data string
			_ = s.Get("key", &data)
			fmt.Println(data, data == "value")
			fmt.Println(s.Total())
		})
	}
}
