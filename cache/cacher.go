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
