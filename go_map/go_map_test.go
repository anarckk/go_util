package go_map

import (
	"log"
	"testing"
)

func TestComposeStrArray2(t *testing.T) {
	var arr []map[string]string = []map[string]string{
		{"name": "张三"},
		{"name": "李四"},
	}
	log.Println(ComposeArrayT[map[string]string](arr, func(m map[string]string) string {
		return m["name"]
	}))
}
