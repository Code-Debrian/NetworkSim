package network

type Network struct {
	Links [][]Link
	Nodes []Node
}

func (n *Network) InitLinks(size int) {
	n.Links = make([][]Link, size)
	for i := range n.Links {
		n.Links[i] = make([]Link, size)
	}
}
