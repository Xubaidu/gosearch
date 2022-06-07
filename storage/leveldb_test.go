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
	var s = &LeveldbStorage{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_ = s.InitLevelDB()
			_ = s.db.Put([]byte("key"), []byte("value"), nil)
			data, _ := s.db.Get([]byte("key"), nil)
			fmt.Println(string(data))
		})
	}
}
