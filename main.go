package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/aaronbbrown/hashtest/pkg/distributors"
	"github.com/aaronbbrown/hashtest/pkg/hash"
)

func main() {
	var alg string
	var memberCount int
	var keyCount int
	var hashAlg string
	var keys []string
	var nodes []string
	var distributor distributors.KeyDistributor
	var hasher hash.Hasher

	flag.IntVar(&memberCount, "member-count", 8, "Number of members")
	flag.IntVar(&keyCount, "key-count", 50, "Number of keys")
	flag.StringVar(&alg, "algorithm", "consistent", "algorithm to use (supported: consistent, lafikl-consistent, rendezvous)")
	flag.StringVar(&hashAlg, "hash-algorithm", "maphash", "hashing algorithm (supported: maphash, xxhash)")
	flag.Parse()

	for i := 0; i < memberCount; i++ {
		nodes = append(nodes, fmt.Sprintf("envoy_%d", i))
	}

	for i := 0; i < keyCount; i++ {
		keys = append(keys, fmt.Sprintf("192.168.10.%d:80", i))
	}

	switch alg {
	case "consistent":
		distributor = distributors.NewConsistent(nodes, keys)
	case "lafikl-consistent":
		distributor = distributors.NewLafiklConsistent(nodes, keys)
	case "rendezvous":
		distributor = distributors.NewRendezvous(nodes, keys)
	default:
		log.Fatal("Unknown algorithm", alg)
	}

	switch hashAlg {
	case "maphash":
		hasher = hash.MapHash{}
	case "xxhash":
		hasher = hash.XXHash{}
	default:
		log.Fatal("Unknown hash algorithm", hashAlg)
	}

	distribution, err := distributor.Distribute(hasher)
	if err != nil {
		log.Fatal(err)
	}

	for key, node := range distribution {
		fmt.Println(node, ":", key)
	}
}
