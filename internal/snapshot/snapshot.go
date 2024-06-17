package snapshot

import (
	"time"

	"github.com/delavalom/dar/internal/config"
	"github.com/delavalom/dar/internal/hashMap"
)

type Commit struct {
	Key      string
	Tree     *hashMap.Tree
	Metadata struct {
		Author  string
		Date    string
		Message string
	}
	Parent *Commit
}

func New(key string, tree *hashMap.Tree, message string) *Commit {
	configuration := config.Config{
		Author: "Luis Angel Arvelo",
		Email:  "hi@delavalom",
	}
	return &Commit{
		Key:  key,
		Tree: tree,
		Metadata: struct {
			Author  string
			Date    string
			Message string
		}{
			Author:  configuration.Author,
			Date:    time.Now().String(),
			Message: message,
		},
	}
}
