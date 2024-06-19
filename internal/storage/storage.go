package storage

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"os"

	"github.com/delavalom/dar/internal/hash"
	"github.com/delavalom/dar/internal/hashMap"
)

type Storage struct {
}

func New() *Storage {
	return &Storage{}
}

var IgnoreFiles = map[string]bool{
	".git": true,
	".dar": true,
	"dar":  true,
}

func ReadFiles(files []os.DirEntry, prefix string, tree hashMap.HashMap) {
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

			// creating a new file in the .dar/objects directory with the hash as the name
			file, err := os.Create(fmt.Sprintf(".dar/objects/%s", key))
			if err != nil {
				panic(err)
			}

			// compressing the content and writing it to the file
			var buf bytes.Buffer

			w := zlib.NewWriter(&buf)
			if _, err := w.Write(content); err != nil {
				panic(err)
			}

			if err = w.Close(); err != nil {
				panic(err)
			}

			if _, err = io.Copy(file, &buf); err != nil {
				panic(err)
			}
		} else {
			if IgnoreFiles[filePath] {
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
			ReadFiles(files, filePath, tree[filePath].SubTree)
		}
	}
}

func ReadTreeAndWriteFiles(tree hashMap.HashMap) {
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
					fmt.Println(contentPath)
					compressedFile, err := os.Open(contentPath)
					if err != nil {
						panic(err)
					}
					defer compressedFile.Close()

					reader, err := zlib.NewReader(compressedFile)
					if err != nil {
						panic("Reading contentFile content: " + err.Error())
					}
					defer reader.Close()

					file, err := os.Create(key)
					if err != nil {
						panic("Creating file path: " + err.Error())
					}
					defer file.Close()

					if _, err = io.Copy(file, reader); err != nil {
						panic("copy from compressed file to repository file system: " + err.Error())
					}

					visited[key] = true
				}
			}
		}
	}
}
