package mathx

import (
	"math/rand"
	"sync"
)

type lockedRandSource struct {
	source rand.Source
	lock   sync.Mutex
}

func NewLockedRandSource(seed int64) *lockedRandSource {
	return &lockedRandSource{
		source: rand.NewSource(seed),
	}
}

func (ls *lockedRandSource) Int63() int64 {
	ls.lock.Lock()
	defer ls.lock.Unlock()

	return ls.source.Int63()
}

func (ls *lockedRandSource) Seed(seed int64) {
	ls.lock.Lock()
	defer ls.lock.Unlock()

	ls.source.Seed(seed)
}
