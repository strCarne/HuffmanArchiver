package huff

import "HuffmanArchiver/priorityq"

// New archived file extension
const extensionName = ".huff"

// Huffman struct that compresses and decompresses files
type Huffman struct {
	fileName     string
	newExtension string
}

// buildHuffmanTree function builds one single
// tree with all the symbols.
// This is the main part of Huffman's algorithm
// to build new symbol's codes for compressing
// and decompressing.
func (h *Huffman) buildHuffmanTree(prQ *priorityq.PriorityQ[*Node]) *Node {
	for prQ.Len() != 1 {
		n1 := prQ.Pop().(*Node)
		n2 := prQ.Pop().(*Node)
		prQ.Push(mergeNodes(n1, n2))
	}
	return prQ.Pop().(*Node)
}

// makeCodes function builds new codes
// based on Huffman's binary code tree
func (h *Huffman) makeCodes(root *Node) []string {
	codes := make([]string, 256)
	ptr := root
	dfs(ptr, codes, "", "")
	return codes
}

// dfs - Depth-First Search for visiting
// all the nodes of tree and build codes
// based on what edges it went through to
// the leaf
func dfs(node *Node, codes []string, tmp, plus string) {
	if node == nil {
		return
	}
	tmp += plus
	if node.IsLeaf() {
		codes[node.symbol] = tmp
	}
	dfs(node.bit0, codes, tmp, "0")
	dfs(node.bit1, codes, tmp, "1")
}
