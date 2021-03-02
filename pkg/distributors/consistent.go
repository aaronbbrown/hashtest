package distributors

import (
	"github.com/aaronbbrown/hashtest/pkg/hash"
	"github.com/aaronbbrown/hashtest/pkg/types"
	"github.com/buraksezer/consistent"
)

type Consistent struct {
	nodes []string
	keys  []string
}

func NewConsistent(nodes, keys []string) *Consistent {
	return &Consistent{
		nodes: nodes,
		keys:  keys,
	}
}

func (c *Consistent) Distribute(hasher hash.Hasher) (map[string]string, error) {
	result := make(map[string]string)

	cfg := consistent.Config{
		PartitionCount:    541,
		ReplicationFactor: 2,
		Load:              1.25,
		Hasher:            hasher,
	}

	distributor := consistent.New(nil, cfg)

	for _, node := range c.nodes {
		distributor.Add(types.Node(node))
	}

	for _, key := range c.keys {
		result[key] = distributor.LocateKey([]byte(key)).String()
	}

	return result, nil
}
