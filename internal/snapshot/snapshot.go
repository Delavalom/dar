package snapshot

import (
	"time"

	"github.com/delavalom/dar/internal/config"
	"github.com/delavalom/dar/internal/hashMap"
)

type metadata struct {
	Author  string `json:"author"`
	Date    string `json:"date"`
	Message string `json:"message"`
}

type Commit struct {
	Key      string          `json:"key"`
	Tree     hashMap.HashMap `json:"tree"`
	Metadata metadata        `json:"metadata"`
	Parent   *Commit         `json:"parent"`
}

func New(key string, tree hashMap.HashMap, message string) *Commit {
	configuration := config.Config{
		Author: "Luis Angel Arvelo",
		Email:  "hi@delavalom",
	}
	return &Commit{
		Key:  key,
		Tree: tree,
		Metadata: metadata{
			Author:  configuration.Author,
			Date:    time.Now().String(),
			Message: message,
		},
	}
}
