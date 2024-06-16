package storage

import (
	"compress/zlib"
	"fmt"
	"os"

	"github.com/delavalom/dar/internal/hash"
)

type Hash string
type Blob string

type Storage struct {
	Commit map[Hash]Blob
	Tree   map[Hash]Blob
	Blob   Blob
}

func New() *Storage {
	return &Storage{
		Commit: make(map[Hash]Blob),
		Tree:   make(map[Hash]Blob),
	}
}
func (s *Storage) AddBlob(hash Hash, blob Blob) {
	s.Commit[hash] = blob
}

func (s *Storage) AddTree(hash Hash, blob Blob) {
	s.Commit[hash] = blob
}

var IgnoreFiles = map[string]bool{
	".git": true,
	".dar": true,
}

func ReadFiles(files []os.DirEntry, prefix string) {
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
			// fmt.Println(filePath)
			fmt.Println(file.Type().Perm().String())
			length := len(content)
			content = append([]byte("\n"), content...)
			content = append([]byte{byte(length)}, content...)
			key := hash.New(content)
			file, err := os.Create(fmt.Sprintf(".dar/objects/%s", key))
			if err != nil {
				panic(err)
			}
			_, err = zlib.NewWriter(file).Write(content)
			if err != nil {
				panic(err)
			}

		} else {
			if IgnoreFiles[filePath] {
				continue
			}
			files, err := os.ReadDir(filePath)
			if err != nil {
				panic("Couldn't open the Directory: " + err.Error())
			}
			ReadFiles(files, filePath)
		}
	}
}

type Files map[string]*File

type File struct {
	IsDir   bool
	Name    string
	Content string           // empty if IsDir is true
	Trees   map[string]*File // empty if IsDir is false
}
