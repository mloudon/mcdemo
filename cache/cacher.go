package cache

import (
	"fmt"
	"io"
	"os"

	"github.com/bradfitz/gomemcache/memcache"
	l5g "github.com/neocortical/log5go"
)

var log = l5g.Logger(l5g.LogAll)

type Cacher interface {
	PutFile(name string) error
	GetFile(name, outname string) error
}

type cacher struct {
	mc *memcache.Client
}

func New(mc *memcache.Client) Cacher {
	return &cacher{
		mc: mc,
	}
}

func (c *cacher) PutFile(name string) (err error) {
	chunks, err := readChunks(name)
	if err != nil {
		return
	}

	for i, chunk := range chunks {
		key := fmt.Sprintf("%s-%d", name, i)
		c.mc.Set(&memcache.Item{Key: key, Value: chunk})
	}
	return
}

func (c *cacher) GetFile(name, outname string) (err error) {
	i := 0
	for {
		key := fmt.Sprintf("%s-%d", name, i)
		it, err := mc.Get(key)
		if err != nil {
			return
		}
		i++
	}
}

func readChunks(name string) (chunks [][]byte, err error) {
	fi, err := os.Open(name)
	if err != nil {
		return
	}

	defer fi.Close()

	chunk := make([]byte, 1024)
	for {
		n, e := fi.Read(chunk)
		if e != nil && e != io.EOF {
			err = e
			return
		}
		if n == 0 {
			break
		}

		chunks = append(chunks, chunk)
	}

	return
}

func writeChunks(name string, chunks [][]byte) (err error) {
	fi, err := os.Open(name)
	if err != nil {
		return
	}

	defer fi.Close()

	for _, chunk := range chunks {
		//write chunk
	}

	return
}
