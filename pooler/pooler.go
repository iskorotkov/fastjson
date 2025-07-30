package pooler

import (
	"reflect"
	"sync"
)

func NewPooler() *Pooler {
	return &Pooler{
		pools: make(map[SliceTier]*sync.Pool),
	}
}

type Pooler struct {
	pools map[SliceTier]*sync.Pool
	mu    sync.RWMutex
}

func (p *Pooler) Append(slice *reflect.Value, item *reflect.Value) *reflect.Value {
	length := slice.Len()
	if length < slice.Cap() {
		slice.SetLen(length + 1)
		slice.Index(length).Set(*item)
		return slice
	}

	tier := SliceTier{
		Type: slice.Type(),
		Size: max(lowerPowerOf2(2*length), 3),
	}

	newSlice := p.Get(tier)
	newSlice.SetLen(length + 1)

	reflect.Copy(*newSlice, *slice)
	newSlice.Index(length).Set(*item)

	p.Put(slice)

	return newSlice
}

func (p *Pooler) Get(tier SliceTier) *reflect.Value {
	pool := p.pool(tier)
	return pool.Get().(*reflect.Value)
}

func (p *Pooler) Put(slice *reflect.Value) {
	slice.SetLen(0)
	tier := SliceTier{
		Type: slice.Type(),
		Size: lowerPowerOf2(slice.Cap()),
	}
	pool := p.pool(tier)
	pool.Put(slice)
}

func (p *Pooler) pool(tier SliceTier) *sync.Pool {
	p.mu.RLock()
	pool, ok := p.pools[tier]
	p.mu.RUnlock()

	if ok {
		return pool
	}

	pool = &sync.Pool{
		New: func() any {
			v := reflect.New(tier.Type).Elem()
			v.Grow(tier.Size)
			return &v
		},
	}

	p.mu.Lock()
	p.pools[tier] = pool
	p.mu.Unlock()

	return pool
}

type SliceTier struct {
	Type reflect.Type
	Size int
}

func lowerPowerOf2(x int) int {
	if x <= 0 {
		return 0
	}
	y := 1
	for y < x {
		y <<= 1
	}
	return y - 1
}
