package cache

import (
	"github.com/bradfitz/gomemcache/memcache"
	l5g "github.com/neocortical/log5go"
)

var log = l5g.Logger(l5g.LogAll)

type Cacher interface {
	PutFile(fullpath string) error
	GetFile(name string) ([]byte, error)
}

type cacher struct {
	mc *memcache.Client
}

func New(mc *memcache.Client) Cacher {
	return &cacher{
		mc: mc,
	}
}

func (c *cacher) PutFile(fullpath string) (err error) {
	return
}

func (c *cacher) GetFile(name string) (data []byte, err error) {
	return
}
