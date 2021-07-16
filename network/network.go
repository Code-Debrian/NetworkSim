package network

type Network struct {
	Links     [][]Link
	Nodes     []Node
	NodeCount int
}

func (n *Network) Init(nodeCount int) {
	n.NodeCount = nodeCount

	n.Links = make([][]Link, n.NodeCount)
	for i := range n.Links {
		n.Links[i] = make([]Link, 0)
	}
}

// TODO: func for transfering bytes from node to node
// Need to do some research first and write down the whole process
