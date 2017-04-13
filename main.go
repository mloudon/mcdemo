package main

import (
	"fmt"

	"github.com/bradfitz/gomemcache/memcache"
)

func main() {
	mc := memcache.New("172.17.0.6:11211")
	mc.Set(&memcache.Item{Key: "foo", Value: []byte("my value")})

	it, err := mc.Get("foo")
	if err != nil {
		panic(err)
	}
	fmt.Printf("got %+s\n", string(it.Value))
}
