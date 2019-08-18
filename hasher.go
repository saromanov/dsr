package dsr

import (
	"hash/fnv"
)

type Hasher interface {
	Sum([]byte) uint64
}

type hasherfnv struct {
}

func (h *hasherfnv) Sum(b []byte) uint64 {
	data := fnv.New64a()
	data.Write(b)
	return data.Sum64()
}
func NewHasher() Hasher {
	return &hasherfnv{}
}
