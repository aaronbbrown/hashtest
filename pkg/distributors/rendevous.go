package distributors

import (
	"github.com/aaronbbrown/hashtest/pkg/hash"
	grend "github.com/dgryski/go-rendezvous"
)

type Rendezvous struct {
	nodes []string
	keys  []string
}

func NewRendezvous(nodes, keys []string) *Rendezvous {
	return &Rendezvous{
		nodes: nodes,
		keys:  keys,
	}
}

func (d *Rendezvous) Distribute(hasher hash.Hasher) (map[string]string, error) {
	result := make(map[string]string)
	r := grend.New(d.nodes, func(s string) uint64 {
		return hasher.Sum64([]byte(s))
	})

	for _, node := range d.nodes {
		r.Add(node)
	}

	for _, key := range d.keys {
		node := r.Lookup(key)
		result[key] = node
	}

	return result, nil
}
