package cache

import (
	"errors"
	"fmt"
	"io"

	"github.com/bradfitz/gomemcache/memcache"
	l5g "github.com/neocortical/log5go"
)

var log = l5g.Logger(l5g.LogAll)

const (
	partKeyTmpl = "%s.%d.part"
	endKeyTmpl  = "%s.%d.end"
)

type Cacher interface {
	Store(key string, r io.Reader) error
	Retrieve(key string, w io.Writer) error
}

type cacher struct {
	mc *memcache.Client
}

func NewCacher(mc *memcache.Client) Cacher {
	return &cacher{
		mc: mc,
	}
}

func (c *cacher) Store(key string, r io.Reader) (err error) {
	chunks, err := readChunks(r)
	if err != nil {
		return
	}

	for i, chunk := range chunks {
		k := fmt.Sprintf(partKeyTmpl, key, i)
		log.Debug("storing part key %s", k)
		err = c.mc.Set(&memcache.Item{Key: k, Value: chunk})
		if err != nil {
			return
		}
	}
	k := fmt.Sprintf(endKeyTmpl, key, len(chunks))
	log.Debug("storing end key %s", k)
	c.mc.Set(&memcache.Item{Key: k}) //store EOF record
	return
}

func (c *cacher) Retrieve(key string, w io.Writer) error {
	chunks := [][]byte{}
	i := 0
	for {
		k := fmt.Sprintf(partKeyTmpl, key, i)
		log.Debug("fetching part key %s", k)
		it, err := c.mc.Get(k)
		if err == memcache.ErrCacheMiss {
			k = fmt.Sprintf(endKeyTmpl, key, i)
			log.Debug("fetching end key %s", k)
			_, e := c.mc.Get(k) //check for EOF record
			if e == nil {
				break
			}

			if e == memcache.ErrCacheMiss {
				err = errors.New("EOF not found")
			}
		}

		if err != nil {
			return err
		}

		chunks = append(chunks, it.Value)
		i++
	}

	return writeChunks(w, chunks)
}

func readChunks(r io.Reader) (chunks [][]byte, err error) {
	chunk := make([]byte, 1024)
	for {
		n, e := r.Read(chunk)
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

func writeChunks(w io.Writer, chunks [][]byte) error {
	for _, chunk := range chunks {
		_, err := w.Write(chunk)
		if err != nil {
			return err
		}
	}
	return nil
}
