package distributors

import "github.com/aaronbbrown/hashtest/pkg/hash"

type KeyDistributor interface {
	Distribute(hash.Hasher) (map[string]string, error)
}
