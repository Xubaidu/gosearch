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
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = InitLevelDB()
			_ = DB.Put([]byte("key"), []byte("value"), nil)
			data, _ := DB.Get([]byte("key"), nil)
			fmt.Println(string(data))
		})
	}
}
