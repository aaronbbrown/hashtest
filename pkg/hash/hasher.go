package hash

import (
	"hash/maphash"

	"github.com/cespare/xxhash"
)

// consistent package doesn't provide a default hashing function.
// You should provide a proper one to distribute keys/members uniformly.
type Hasher interface {
	Sum64(data []byte) uint64
}

type MapHash struct{}

func (h MapHash) Sum64(data []byte) uint64 {
	var hash maphash.Hash
	hash.Write(data)

	return hash.Sum64()
}

type XXHash struct{}

func (h XXHash) Sum64(data []byte) uint64 {
	return xxhash.Sum64(data)
}
