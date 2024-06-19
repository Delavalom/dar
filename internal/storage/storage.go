package storage

import (
	"fmt"
	"os"

	"github.com/delavalom/dar/internal/compression"
	"github.com/delavalom/dar/internal/hash"
	"github.com/delavalom/dar/internal/hashMap"
)

var IGNORE_FILES = map[string]bool{
	".git": true,
	".dar": true,
	"dar":  true,
}

type Storage interface {
	WriteTree(files []os.DirEntry, prefix string, tree hashMap.HashMap)
	WriteFiles(tree hashMap.HashMap)
}

// WriteTree reads the files and writes them to the tree hash map structure
// It uses a recursive function to traverse the file system with Depth First Search
func WriteTree(files []os.DirEntry, prefix string, tree hashMap.HashMap) {
	compressor := compression.New()

	for _, file := range files {
		filePath := file.Name()
		if prefix != "" {
			filePath = fmt.Sprintf("%s/%s", prefix, filePath)
		}
		if !file.IsDir() {
			content, err := os.ReadFile(filePath)
			if err != nil {
				panic(err)
			}
			// creating a new hash of the content
			key := hash.New(content)

			// adding the file to the tree
			tree[filePath] = &hashMap.Tree{
				FileContent: key,
				HasSubTree:  false,
				FileName:    file.Name(),
				FileMode:    file.Type().Perm().String(),
			}

			// compressing the content and writing it to the file
			compressor.WriteFile(content, key)
		} else {
			if IGNORE_FILES[filePath] {
				continue
			}
			files, err := os.ReadDir(filePath)

			tree[filePath] = &hashMap.Tree{
				HasSubTree: true,
				SubTree:    make(map[string]*hashMap.Tree),
			}

			if err != nil {
				panic("Couldn't open the Directory: " + err.Error())
			}
			WriteTree(files, filePath, tree[filePath].SubTree)
		}
	}
}

// WriteFiles reads the tree hash map structure and writes the files to the file system
// It uses a queue to traverse the tree with Breadth First Search
func WriteFiles(tree hashMap.HashMap) {
	compressor := compression.New()

	type queueItem struct {
		path string
		tree hashMap.HashMap
	}
	// Create a queue
	queue := make([]*queueItem, 0)
	visited := make(map[string]bool)

	queue = append(queue, &queueItem{
		path: ".",
		tree: tree,
	})

	for len(queue) > 0 {
		// Dequeue
		current := queue[0]
		queue = queue[1:]
		visited[current.path] = true

		// Iterate over the current tree
		for key, file := range current.tree {
			if file.HasSubTree {
				if !visited[key] {
					// Create the directory
					if err := os.MkdirAll(key, os.ModePerm); err != nil {
						panic("Creating directory of the file system" + err.Error())
					}
					visited[key] = true
				}
				// Enqueue the subtree
				queue = append(queue, &queueItem{path: key, tree: file.SubTree})
			} else {
				if !visited[key] {
					// Create and write the file
					contentPath := fmt.Sprintf(".dar/objects/%s", file.FileContent)
					compressor.ReadAndWriteDecompressedFile(key, contentPath)

					visited[key] = true
				}
			}
		}
	}
}
