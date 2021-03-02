package types

type Node string

func (n Node) String() string {
	/*
		hashed := hasher{}.Sum64([]byte(n))
		return fmt.Sprintf("%x", hashed)
	*/
	return string(n)
}
