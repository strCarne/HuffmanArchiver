package huff

import (
	"HuffmanArchiver/priorityq"
	"HuffmanArchiver/sll"
	"log"
	"os"
	"strings"
)

func (h *Huffman) DecompressFile(archFileName string) {
	archData, err := os.ReadFile(archFileName)

	if err != nil {
		log.Fatal(err)
	}

	data := h.decompressBytes(archData)

	os.Remove(archFileName)
	os.WriteFile(archFileName[:len(archFileName)-5]+h.newExtension, data, 0777)
}

func (h *Huffman) decompressBytes(archData []byte) []byte {
	numOfSymb, startIndex, freqs := h.parseHeader(archData)

	prQ := priorityq.NewPrQ(buildNodes(freqs))

	root := h.buildHuffmanTree(prQ)

	data := h.decompress(archData, startIndex, numOfSymb, root)

	return data
}

func (h *Huffman) parseHeader(archData []byte) (numOfSymb int, startIndex int, freqs []int) {
	oldExtLen := int(archData[0])
	sb := strings.Builder{}
	for i := 1; i <= oldExtLen; i++ {
		sb.WriteByte(archData[i])
	}
	h.newExtension = sb.String()

	startIndex = oldExtLen + 1
	numOfSymb = int(archData[startIndex]) | (int(archData[startIndex+1]) << 8) | (int(archData[startIndex+2]) << 16) | (int(archData[startIndex+3]) << 24)
	freqs = make([]int, 256)
	for i := range freqs {
		freqs[i] = int(archData[startIndex+i+4])
	}
	startIndex += 4 + 256
	return numOfSymb, startIndex, freqs
}

func (h *Huffman) decompress(archData []byte, startIndex int, numOfSymb int, root *Node) []byte {
	size := 0
	ptr := root
	data := &sll.List{}
	for i := startIndex; i < len(archData); i++ {
		for bit := 1; bit <= 128; bit <<= 1 {
			zero := (archData[i] & byte(bit)) == 0
			if zero {
				ptr = ptr.bit0
			} else {
				ptr = ptr.bit1
			}
			if ptr.bit0 != nil {
				continue
			}
			if size < numOfSymb {
				data.Append(ptr.symbol)
			}
			ptr = root
			size++
		}
	}
	return data.ToArr()
}
