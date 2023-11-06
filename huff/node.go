package huff

type Node struct {
	bit0   *Node
	bit1   *Node
	symbol byte
	freq   int
}

func NewNode(symbol byte, frequency int) *Node {
	return &Node{
		symbol: symbol,
		freq:   frequency,
	}
}

func mergeNodes(n1, n2 *Node) *Node {
	return &Node{
		bit0: n1,
		bit1: n2,
		freq: n1.freq + n2.freq,
	}
}

func (n *Node) Priority() int {
	return n.freq
}

func buildNodes(freqs []int) []*Node {
	nodes := make([]*Node, 0)
	for symbol, freq := range freqs {
		if freq != 0 {
			nodes = append(nodes, NewNode(byte(symbol), freq))
		}
	}
	return nodes
}

func (n *Node) IsLeaf() bool {
	return n.bit0 == nil || n.bit1 == nil
}
