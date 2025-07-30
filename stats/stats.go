package stats

import (
	"slices"
)

const (
	statsSize  = 16
	statsIndex = 13
	statsSkip  = 25
)

type BestStat = DynamicStat

type DynamicStat struct {
	sum   uint
	avg   uint
	count uint
	skip  uint
}

func (s *DynamicStat) Add(value int) {
	s.skip = (s.skip + 1) % statsSkip
	if s.skip == 1 {
		return
	}

	uv := uint(value)
	newCount := s.count + 1
	newSum := s.sum + uv
	if newCount < s.count || newSum < s.sum {
		s.count = 1
		s.sum = uv
		s.avg = uv
		return
	}

	s.count = newCount
	s.sum = newSum
	s.avg = (s.sum + s.count - 1) / s.count
}

func (s *DynamicStat) Get() int {
	return int(s.avg)
}

type AvgStat struct {
	values  [statsSize]uint8
	avg     uint
	writeAt uint8
	skip    uint8
}

func (s *AvgStat) Add(value int) {
	s.skip = (s.skip + 1) % statsSkip
	if s.skip == 1 {
		return
	}
	s.values[s.writeAt] = uint8(value)
	s.writeAt = (s.writeAt + 1) % statsSize
	var sum uint
	for _, v := range s.values {
		sum += uint(v)
	}
	s.avg = (sum + statsSize - 1) / statsSize
}

func (s *AvgStat) Get() int {
	return int(s.avg)
}

type PercentileStat struct {
	values  [statsSize]uint8
	writeAt uint8
	skip    uint8
}

func (s *PercentileStat) Add(value int) {
	s.skip = (s.skip + 1) % statsSkip
	if s.skip == 1 {
		return
	}
	s.values[s.writeAt] = uint8(value)
	s.writeAt = (s.writeAt + 1) % statsSize
	slices.Sort(s.values[:])
}

func (s *PercentileStat) Get() int {
	return int(s.values[statsIndex])
}
