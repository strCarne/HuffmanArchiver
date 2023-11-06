package main

import (
	"HuffmanArchiver/huff"
	"log"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	workDir := MustGetWD()
	fileNames := parseCL()
	wg := sync.WaitGroup{}
	for _, fileName := range fileNames {
		wg.Add(1)
		go huff.StartProcessing(&wg, filepath.Join(workDir, fileName))
	}
	wg.Wait()
}

func MustGetWD() string {
	workDir, err := os.Getwd()
	if err != nil {
		log.Fatal(err)
	}
	return workDir
}

func parseCL() []string {
	return os.Args[1:]
}
