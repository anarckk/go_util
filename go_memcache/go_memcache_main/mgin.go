package main

import (
	"log"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	mc := memcache.New("go_memcache:11211")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})
	it, err := mc.Get("foo")
	if err != nil {
		panic(err)
	}
	log.Println(string(it.Value))
}
