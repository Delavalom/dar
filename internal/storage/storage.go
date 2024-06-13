package storage

type Hash string
type Blob string

type Storage struct {
	Commit map[Hash]Blob
	Tree   map[Hash]Blob
	Blob   Blob
}
