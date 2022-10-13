package bigcache

import (
	"github.com/allegro/bigcache"
	log "github.com/sirupsen/logrus"
	"time"
)

var bcConfig = bigcache.Config{
	Shards:             1024,
	LifeWindow:         999999 * time.Minute,
	MaxEntriesInWindow: 1000 * 10 * 60,
	MaxEntrySize:       500,
	Verbose:            true,
	HardMaxCacheSize:   8192,
	OnRemove:           nil,
	OnRemoveWithReason: nil,
}

type BigcacheHandler interface {
	New() *bigcache.BigCache
}
type BigcacheSvc struct {
}

func (b BigcacheSvc) New() *bigcache.BigCache {
	cache, err := bigcache.NewBigCache(bcConfig)

	if err != nil {
		log.Fatal(err)
	}

	return cache
}

func NewBigcacheHandler() BigcacheHandler {
	return &BigcacheSvc{}
}
