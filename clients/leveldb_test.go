package clients

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
			err := InitLevelDB()
			if err != nil {
				fmt.Println(err)
			}
			err = DB.Put([]byte("key"), []byte("value"), nil)
			if err != nil {
				fmt.Println(err)
			}
			data, err := DB.Get([]byte("key"), nil)
			if err != nil {
				fmt.Println(err)
			}
			fmt.Println(string(data))
		})
	}
}
