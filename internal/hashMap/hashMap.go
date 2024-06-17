package hashMap

import (
	"crypto/sha1"
)

type Tree struct {
	HasSubTree  bool             `json:"hasSubTree"`
	SubTree     map[string]*Tree `json:"subTree"`     // nil if HasSubTree is false
	FileName    string           `json:"fileName"`    // empty if HasSubTree is true
	FileMode    string           `json:"fileMode"`    // empty if HasSubTree is true
	FileContent string           `json:"fileContent"` // empty if HasSubTree is true
}

type Files map[string]*File

type File struct {
	IsDir   bool
	Name    string
	Content string           // empty if IsDir is true
	Trees   map[string]*File // empty if IsDir is false
}

func New() map[string]*Tree {
	return make(map[string]*Tree)
}

func Store(content []byte) {
	hash := sha1.New()
	hash.Write(content)

}
