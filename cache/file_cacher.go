package cache

import (
	"os"

	"github.com/bradfitz/gomemcache/memcache"
)

type FileCacher interface {
	Store(fname string) error
	Retrieve(fname, outname string) error
}

type fileCacher struct {
	cacher Cacher
}

func NewFileCacher(mc *memcache.Client) FileCacher {
	return &fileCacher{cacher: NewCacher(mc)}
}

func (fc *fileCacher) Store(fname string) (err error) {
	fi, err := os.Open(fname)
	if err != nil {
		return
	}
	defer fi.Close()

	return fc.cacher.Store(fname, fi)
}

func (fc *fileCacher) Retrieve(fname, outname string) (err error) {
	fi, err := os.Create(outname)
	if err != nil {
		return
	}
	defer fi.Close()

	return fc.cacher.Retrieve(fname, fi)
}
