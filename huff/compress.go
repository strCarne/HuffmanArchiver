package huff

import (
	"HuffmanArchiver/priorityq"
	"HuffmanArchiver/sll"
	"log"
	"os"
	"path/filepath"
)

// Wrapping of the compression function
func (h *Huffman) CompressFile(fileName string) {

	data, err := os.ReadFile(fileName)
	h.fileName = fileName

	if err != nil {
		log.Fatal(err)
	}

	archData := h.compressBytes(data)
	os.Remove(fileName)
	os.WriteFile(fileName[:len(fileName)-len(filepath.Ext(fileName))]+extensionName, archData, 0777)
}

// The main file byte's compression function containing
// successive calls to the Huffman algorithm compression steps
func (h *Huffman) compressBytes(data []byte) []byte {
	freqs := h.calcFreqs(data)

	prQ := priorityq.NewPrQ(buildNodes(freqs))

	root := h.buildHuffmanTree(prQ)

	codes := h.makeCodes(root)

	archedData := h.compress(data, codes)

	header := h.createHeader(len(data), freqs)

	res := append(header, archedData...)

	return res
}

// calcFreqs function calculates the frequency of
// each symbol's (byte's) appearing in the base alpabet
func (h *Huffman) calcFreqs(data []byte) []int {
	freqs := make([]int, 256)
	for _, Byte := range data {
		freqs[Byte]++
	}
	normalize(freqs)
	return freqs
}

func normalize(freqs []int) {
	max := findMax(freqs)
	if max <= 255 {
		return
	}
	for i := range freqs {
		if freqs[i] > 0 {
			freqs[i] = 1 + freqs[i]*255/(max+1)
		}
	}
}

func findMax(arr []int) int {
	max := arr[0]
	for i := 1; i < len(arr); i++ {
		if arr[i] > max {
			max = arr[i]
		}
	}
	return max
}

func (h *Huffman) compress(data []byte, codes []string) []byte {
	archedData := &sll.List{}
	var sum byte = 0
	var bit byte = 1

	for _, symbol := range data {

		for _, char := range codes[symbol] {
			if char == '1' {
				sum |= bit
			}
			if bit < 128 {
				bit <<= 1
			} else {
				archedData.Append(sum)
				sum = 0
				bit = 1
			}
		}

	}

	if bit > 1 {
		archedData.Append(sum)
	}
	return archedData.ToArr()
}

func (h *Huffman) createHeader(numOfSymb int, freqs []int) []byte {
	header := &sll.List{}

	oldExt := filepath.Ext(h.fileName)
	header.Append(byte(len(oldExt)))
	for i := range oldExt {
		header.Append(oldExt[i])
	}

	//  Writing 4 bytes of length into the list using bit shifts
	header.Append(byte(numOfSymb & 255))
	header.Append(byte((numOfSymb >> 8) & 255))
	header.Append(byte((numOfSymb >> 16) & 255))
	header.Append(byte((numOfSymb >> 24) & 255))

	for _, val := range freqs {
		header.Append(byte(val))
	}

	return header.ToArr()
}
