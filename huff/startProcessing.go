package huff

import (
	"strings"
	"sync"
)

func StartProcessing(wg *sync.WaitGroup, fileName string) {
	defer wg.Done()
	if len(fileName) == 0 {
		return
	}

	huffman := Huffman{}

	if strings.HasSuffix(fileName, extensionName) {
		huffman.DecompressFile(fileName)
	} else {
		huffman.CompressFile(fileName)
	}
}
