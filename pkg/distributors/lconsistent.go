package distributors

import (
	"github.com/aaronbbrown/hashtest/pkg/hash"
	"github.com/lafikl/consistent"
)

type LafiklConsistent struct {
	nodes []string
	keys  []string
}

func NewLafiklConsistent(nodes, keys []string) *LafiklConsistent {
	return &LafiklConsistent{
		nodes: nodes,
		keys:  keys,
	}
}

func (d *LafiklConsistent) Distribute(_ hash.Hasher) (map[string]string, error) {
	result := make(map[string]string)
	distributor := consistent.New()

	for _, node := range d.nodes {
		distributor.Add(node)
	}

	for _, key := range d.keys {
		node, err := distributor.Get(key)
		if err != nil {
			return result, err
		}

		result[key] = node
	}

	return result, nil
}
