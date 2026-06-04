package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"path/filepath"
)

func main() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()

	var nfiles, nbytes int64
	for fileSize := range fileSizes {
		nfiles++
		nbytes += fileSize
	}
	fmt.Printf("%d file %.1f GB\n", nfiles, float64(nbytes)/1e9)
}

func dirents(dir string) []os.DirEntry {
	entries, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	return entries
}

func walkDir(dir string, fileSizes chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subDir := filepath.Join(dir, entry.Name())
			walkDir(subDir, fileSizes)
		} else {
			info, _ := entry.Info()
			fileSizes <- info.Size()
		}
	}
}
